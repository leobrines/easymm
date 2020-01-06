package steam

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var (
	steamLogin = "https://steamcommunity.com/openid/login"

	openIDMode       = "checkid_setup"
	openIDNs         = "http://specs.openid.net/auth/2.0"
	openIDIdentifier = "http://specs.openid.net/auth/2.0/identifier_select"

	steamOpenIDUrlRegexp   = regexp.MustCompile("^(http|https)://steamcommunity.com/openid/id/[0-9]{15,25}$")
	digitsExtractionRegexp = regexp.MustCompile("\\D+")
)

type OpenId struct {
	root      string
	returnUrl string
	data      url.Values
}

func NewOpenId(r *http.Request) *OpenId {
	id := new(OpenId)

	var proto string
	if r.TLS != nil {
		proto = "https://"
	} else {
		proto = "http://"
	}

	id.root = proto + r.Host

	uri := r.RequestURI
	if i := strings.Index(uri, "openid."); i != -1 {
		uri = uri[0 : i-1]
	}
	id.returnUrl = id.root + uri

	switch r.Method {
	case "POST":
		r.ParseForm()
		id.data = r.Form
	case "GET":
		id.data = r.URL.Query()
	}

	return id
}

func (id OpenId) AuthUrl() string {
	data := make(url.Values)
	data.Set("openid.claimed_id", openIDIdentifier)
	data.Set("openid.identity", openIDIdentifier)
	data.Set("openid.mode", openIDMode)
	data.Set("openid.ns", openIDNs)
	data.Set("openid.realm", id.root)
	data.Set("openid.return_to", id.returnUrl)

	url := steamLogin + "?" + data.Encode()
	return url
}

func (id *OpenId) ValidateAuth() (string, error) {
	if id.Mode() != "id_res" {
		return "", fmt.Errorf("Mode must equal to \"id_res\"")
	}

	if id.data.Get("openid.return_to") != id.returnUrl {
		return "", fmt.Errorf("The \"return_to url\" must match the url of current request.")
	}

	params := make(url.Values)
	params.Set("openid.assoc_handle", id.data.Get("openid.assoc_handle"))
	params.Set("openid.signed", id.data.Get("openid.signed"))
	params.Set("openid.sig", id.data.Get("openid.sig"))
	params.Set("openid.ns", id.data.Get("openid.ns"))

	split := strings.Split(id.data.Get("openid.signed"), ",")
	for _, item := range split {
		params.Set("openid."+item, id.data.Get("openid."+item))
	}
	params.Set("openid.mode", "check_authentication")

	resp, err := http.PostForm(steamLogin, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	response := strings.Split(string(content), "\n")
	if response[0] != "ns:"+openIDNs {
		return "", fmt.Errorf("Wrong ns in the response")
	}
	if strings.HasSuffix(response[1], "false") {
		return "", fmt.Errorf("Unable validate openId")
	}

	openIdUrl := id.data.Get("openid.claimed_id")
	if !steamOpenIDUrlRegexp.MatchString(openIdUrl) {
		return "", errors.New("Invalid steam id pattern.")
	}

	return digitsExtractionRegexp.ReplaceAllString(openIdUrl, ""), nil
}

func (id OpenId) Mode() string {
	return id.data.Get("openid.mode")
}

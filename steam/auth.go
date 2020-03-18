package steam

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func AuthURL(r *http.Request) string {
	var proto, root, return_url string

	if r.TLS != nil {
		proto = "https://"
	} else {
		proto = "http://"
	}

	root = proto + r.Host
	return_url = root + steamLoginCallbackEnpoint

	if i := strings.Index(r.RequestURI, "openid."); i != -1 {
		return_url += r.RequestURI[0 : i-1]
	}

	query := make(url.Values)
	query.Set("openid.claimed_id", openIDIdentifier)
	query.Set("openid.identity", openIDIdentifier)
	query.Set("openid.mode", openIDMode)
	query.Set("openid.ns", openIDNs)
	query.Set("openid.realm", root)
	query.Set("openid.return_to", return_url)
	return steamLogin + "?" + query.Encode()
}

func ValidateAuth(r *http.Request) (string, error) {
	query := r.URL.Query()
	fmt.Printf("original query validate: %v\n", query)

	// Validate request query values
	if query.Get("openid.mode") != "id_res" {
		return "", fmt.Errorf("mode must equal to \"id_res\"")
	}

	steamOpenidURL := query.Get("openid.claimed_id")
	if !steamOpenIDUrlRegexp.MatchString(steamOpenidURL) {
		return "", errors.New("Invalid steam id pattern.")
	}

	steamid := digitsExtractionRegexp.ReplaceAllString(steamOpenidURL, "")

	// Validate query with steam API
	postQuery := make(url.Values)

	keysSigned := strings.Split(query.Get("openid.signed"), ",")
	keysSigned = append(keysSigned, "sig")
	keysSigned = append(keysSigned, "ns")

	for _, item := range keysSigned {
		k := "openid." + item
		postQuery.Set(k, query.Get(k))
	}

	postQuery.Set("openid.mode", "check_authentication")

	resp, err := http.PostForm(steamLogin, postQuery)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Validate steam API response
	response := strings.Split(string(content), "\n")
	if response[0] != "ns:"+openIDNs {
		return "", fmt.Errorf("wrong ns in the response")
	}
	if strings.HasSuffix(response[1], "false") {
		return "", fmt.Errorf("unable validate openId")
	}

	return steamid, nil
}

package steam

import (
	"os"
	"regexp"
)

var (
	steamLogin                = "https://steamcommunity.com/openid/login"
	steamLoginCallbackEnpoint = "/steam/login/callback"

	openIDMode       = "checkid_setup"
	openIDNs         = "http://specs.openid.net/auth/2.0"
	openIDIdentifier = "http://specs.openid.net/auth/2.0/identifier_select"

	steamOpenIDUrlRegexp   = regexp.MustCompile("^(http|https)://steamcommunity.com/openid/id/[0-9]{15,25}$")
	digitsExtractionRegexp = regexp.MustCompile("\\D+")

	steamAPIGetPlayerSummariesURL = "http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s"

	apiKey = os.Getenv("STEAM_API_KEY")
)

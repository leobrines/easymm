package steam

type PlayerSummaries struct {
	SteamId                  string `json:"steamid"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	ProfileState             int    `json:"profilestate"`
	PersonaName              string `json:"personaname"`
	LastLogOff               int    `json:"lastlogoff"`
	ProfileUrl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	PersonaState             int    `json:"personastate"`

	CommentPermission int    `json:"commentpermission"`
	RealName          string `json:"realname"`
	PrimaryClanId     string `json:"primaryclanid"`
	TimeCreated       int    `json:"timecreated"`
	LocCountryCode    string `json:"loccountrycode"`
	LocStateCode      string `json:"locstatecode"`
	LocCityId         int    `json:"loccityid"`
	GameId            int    `json:"gameid"`
	GameExtraInfo     string `json:"gameextrainfo"`
	GameServerIp      string `json:"gameserverip"`
}

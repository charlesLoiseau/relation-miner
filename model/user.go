package model

type User struct {
	Id                             int         `json:"id"`
	IdStr                          string      `json:"id_str"`
	Name                           string      `json:"name"`
	ScreenName                     string      `json:"screen_name"`
	Location                       string      `json:"location"`
	ProfileLocation                interface{} `json:"profile_location"`
	Description                    string      `json:"description"`
	Url                            string      `json:"url"`
	Entities                       interface{} `json:"entities"`
	Protected                      bool        `json:"protected"`
	FollowersCount                 int         `json:"followers_count"`
	FriendsCount                   int         `json:"friends_count"`
	ListedCount                    int         `json:"listed_count"`
	CreatedQt                      string      `json:"created_at"`
	FavouritesCount                int         `json:"favourites_count"`
	UtcOffset                      string      `json:"utc_offset"`
	TimeZone                       string      `json:"time_zone"`
	GeoEnabled                     bool        `json:"geo_enabled"`
	Verified                       bool        `json:"verified"`
	StatusesCount                  int         `json:"statuses_count"`
	Lang                           string      `json:"lang"`
	ContributorsEnabled            bool        `json:"contributors_enabled"`
	IsTranslator                   bool        `json:"is_translator"`
	IsTranslationEnabled           bool        `json:"is_translation_enabled"`
	ProfileBackgroundColor         string      `json:"profile_background_color"`
	ProfileBackgroundImageUrl      string      `json:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps string      `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool        `json:"profile_background_tile"`
	ProfileImageUrl                string      `json:"profile_image_url"`
	ProfileImageUrlHttps           string      `json:"profile_image_url_https"`
	ProfileBannerUrl               string      `json:"profile_banner_url"`
	ProfileLinkColor               string      `json:"profile_link_color"`
	ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string      `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
	HasExtendedProfile             bool        `json:"has_extended_profile"`
	DefaultProfile                 bool        `json:"default_profile"`
	DefaultProfileImage            bool        `json:"default_profile_image"`
	Following                      string      `json:"following"`
	FollowRequestSent              string      `json:"follow_request_sent"`
	Notifications                  string      `json:"notifications"`
	TranslatorType                 string      `json:"translator_type"`
}

type Users []*User
type UsersList struct {
	Users
}

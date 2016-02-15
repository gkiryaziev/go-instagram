package libs

// Get username info.
type UserNameInfo struct {
	Status string `json:"status"`
	User   struct {
		Username                     string `json:"username"`
		IsProfileActionNeeded        bool   `json:"is_profile_action_needed"`
		HasAnonymousProfilePicture   bool   `json:"has_anonymous_profile_picture"`
		MediaCount                   int    `json:"media_count"`
		FollowingCount               int    `json:"following_count"`
		IsNeedy                      bool   `json:"is_needy"`
		AutoExpandChaining           bool   `json:"auto_expand_chaining"`
		HasChaining                  bool   `json:"has_chaining"`
		UsertagReviewEnabled         bool   `json:"usertag_review_enabled"`
		GeoMediaCount                int    `json:"geo_media_count"`
		IncludeDirectBlacklistStatus bool   `json:"include_direct_blacklist_status"`
		ProfilePicURL                string `json:"profile_pic_url"`
		UsertagsCount                int    `json:"usertags_count"`
		Biography                    string `json:"biography"`
		FullName                     string `json:"full_name"`
		FollowerCount                int    `json:"follower_count"`
		Pk                           int64  `json:"pk"`
		IsVerified                   bool   `json:"is_verified"`
		IsPrivate                    bool   `json:"is_private"`
		ExternalURL                  string `json:"external_url"`
	} `json:"user"`
	Message string `json:"message"` // from Error
}

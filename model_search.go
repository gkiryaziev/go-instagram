package instagram

// SearchTags struct
type SearchTags struct {
	HasMore interface{} `json:"has_more"`
	Status  string      `json:"status"`
	Results []struct {
		MediaCount float64 `json:"media_count"`
		Name       string  `json:"name"`
		ID         int64   `json:"id"`
	} `json:"results"`
	Message string `json:"message"` // from Error
}

// SearchUsers struct
type SearchUsers struct {
	HasMore    bool   `json:"has_more"`
	Status     string `json:"status"`
	NumResults int    `json:"num_results"`
	Users      []struct {
		Username                   string `json:"username"`
		HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
		Byline                     string `json:"byline"`
		FriendshipStatus           struct {
			Following       bool `json:"following"`
			IncomingRequest bool `json:"incoming_request"`
			OutgoingRequest bool `json:"outgoing_request"`
			IsPrivate       bool `json:"is_private"`
		} `json:"friendship_status"`
		MutualFollowersCount float64 `json:"mutual_followers_count"`
		ProfilePicURL        string  `json:"profile_pic_url"`
		FullName             string  `json:"full_name"`
		FollowerCount        int     `json:"follower_count"`
		Pk                   int64   `json:"pk"`
		IsVerified           bool    `json:"is_verified"`
		IsPrivate            bool    `json:"is_private"`
	} `json:"users"`
	Message string `json:"message"` // from Error
}

package libs

type Login struct {
	DeviceId          string `json:"device_id"`
	Guid              string `json:"guid"`
	UserName          string `json:"username"`
	Password          string `json:"password"`
	Csrftoken         string `json:"csrftoken"`
	LoginAttemptCount string `json:"login_attempt_count"`
}

type LoginResponse struct {
	LoggedInUser struct {
		Username                   string `json:"username"`
		HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
		ProfilePicURL              string `json:"profile_pic_url"`
		FullName                   string `json:"full_name"`
		Pk                         int64  `json:"pk"`
		IsPrivate                  bool   `json:"is_private"`
	} `json:"logged_in_user"`
	Status  string `json:"status"`
	Message string `json:"message"` // from Error
}

type MediaLikers struct {
	Status    string `json:"status"`
	UserCount int    `json:"user_count"`
	Users     []struct {
		Username      string `json:"username"`
		Pk            int64  `json:"pk"`
		ProfilePicURL string `json:"profile_pic_url"`
		IsPrivate     bool   `json:"is_private"`
		FullName      string `json:"full_name"`
	} `json:"users"`
	Message string `json:"message"` // from Error
}

type RecentActivity struct {
	Status     string `json:"status"`
	NewStories []struct {
		Pk     string `json:"pk"`
		Counts struct {
		} `json:"counts"`
		Args struct {
			Media []struct {
				Image string `json:"image"`
				ID    string `json:"id"`
			} `json:"media"`
			Links []struct {
				Start int    `json:"start"`
				End   int    `json:"end"`
				ID    string `json:"id"`
				Type  string `json:"type"`
			} `json:"links"`
			Text         string  `json:"text"`
			ProfileID    int64   `json:"profile_id"`
			ProfileImage string  `json:"profile_image"`
			Timestamp    float64 `json:"timestamp"`
		} `json:"args"`
		Type int `json:"type"`
	} `json:"new_stories"`
	OldStories []struct {
		Pk     string `json:"pk"`
		Counts struct {
		} `json:"counts"`
		Args struct {
			Media []struct {
				Image string `json:"image"`
				ID    string `json:"id"`
			} `json:"media"`
			Links []struct {
				Start int    `json:"start"`
				End   int    `json:"end"`
				ID    string `json:"id"`
				Type  string `json:"type"`
			} `json:"links"`
			Text         string  `json:"text"`
			ProfileID    int64   `json:"profile_id"`
			ProfileImage string  `json:"profile_image"`
			Timestamp    float64 `json:"timestamp"`
		} `json:"args"`
		Type int `json:"type"`
	} `json:"old_stories"`
	ContinuationToken    int        `json:"continuation_token"`
	FriendRequestStories []struct{} `json:"friend_request_stories"`
	Counts               struct {
		Requests    int `json:"requests"`
		PhotosOfYou int `json:"photos_of_you"`
	} `json:"counts"`
	Subscription interface{} `json:"subscription"`
	Message      string      `json:"message"` // from Error
}

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

type MediaComments struct {
	Status          string `json:"status"`
	CaptionIsEdited bool   `json:"caption_is_edited"`
	HasMoreComments bool   `json:"has_more_comments"`
	Comments        []struct {
		Status    string  `json:"status"`
		MediaID   int64   `json:"media_id"`
		Text      string  `json:"text"`
		CreatedAt float64 `json:"created_at"`
		User      struct {
			Username                   string `json:"username"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			ProfilePicURL              string `json:"profile_pic_url"`
			FullName                   string `json:"full_name"`
			Pk                         int64  `json:"pk"`
			IsPrivate                  bool   `json:"is_private"`
		} `json:"user"`
		ContentType  string `json:"content_type"`
		CreatedAtUtc int    `json:"created_at_utc"`
		Pk           int64  `json:"pk"`
		Type         int    `json:"type"`
	} `json:"comments"`
	Caption struct {
		Status       string `json:"status"`
		UserID       int64  `json:"user_id"`
		CreatedAtUtc int    `json:"created_at_utc"`
		CreatedAt    int    `json:"created_at"`
		BitFlags     int    `json:"bit_flags"`
		User         struct {
			Username      string `json:"username"`
			Pk            int64  `json:"pk"`
			ProfilePicURL string `json:"profile_pic_url"`
			IsPrivate     bool   `json:"is_private"`
			FullName      string `json:"full_name"`
		} `json:"user"`
		ContentType string `json:"content_type"`
		Text        string `json:"text"`
		Pk          int64  `json:"pk"`
		Type        int    `json:"type"`
	} `json:"caption"`
	CommentCount int    `json:"comment_count"`
	Message      string `json:"message"` // from Error
}

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

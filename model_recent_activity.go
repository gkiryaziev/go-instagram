package instagram

// RecentActivity struct
type RecentActivity struct {
	Status               string                  `json:"status"`
	NewStories           []RecentActivityStories `json:"new_stories"`
	OldStories           []RecentActivityStories `json:"old_stories"`
	ContinuationToken    int                     `json:"continuation_token"`
	FriendRequestStories []struct{}              `json:"friend_request_stories"`
	Counts               struct {
		Relationships int `json:"relationships"`
		PhotosOfYou   int `json:"photos_of_you"`
		Usertags      int `json:"usertags"`
		Comments      int `json:"comments"`
		Likes         int `json:"likes"`
		Requests      int `json:"requests"`
	} `json:"counts"`
	Subscription interface{} `json:"subscription"`
	Message      string      `json:"message"` // from Error
}

// RecentActivityStories struct
type RecentActivityStories struct {
	Pk     string `json:"pk"`
	Counts struct {
		Relationships int `json:"relationships"`
		Usertags      int `json:"usertags"`
		Likes         int `json:"likes"`
		Comments      int `json:"comments"`
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
		InlineFollow struct {
			Following bool `json:"following"`
			UserInfo  struct {
				Username      string `json:"username"`
				ProfilePicURL string `json:"profile_pic_url"`
				ID            int64  `json:"id"`
				IsPrivate     bool   `json:"is_private"`
			} `json:"user_info"`
			OutgoingRequest bool `json:"outgoing_request"`
		} `json:"inline_follow"`
	} `json:"args"`
	Type int `json:"type"`
}

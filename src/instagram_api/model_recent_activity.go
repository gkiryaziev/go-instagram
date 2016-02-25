package instagram_api

// Get recent activity.
type RecentActivity struct {
	Status               string                  `json:"status"`
	NewStories           []RecentActivityStories `json:"new_stories"`
	OldStories           []RecentActivityStories `json:"old_stories"`
	ContinuationToken    int                     `json:"continuation_token"`
	FriendRequestStories []struct{}              `json:"friend_request_stories"`
	Counts               struct {
		Requests    int `json:"requests"`
		PhotosOfYou int `json:"photos_of_you"`
	} `json:"counts"`
	Subscription interface{} `json:"subscription"`
	Message      string      `json:"message"` // from Error
}

type RecentActivityStories struct {
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
}

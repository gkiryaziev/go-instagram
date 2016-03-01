package instagram_api

// Get media comments.
type Media struct {
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

// Get media likers.
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

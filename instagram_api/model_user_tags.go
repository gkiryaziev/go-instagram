package instagram_api

// Get user tags
type UserTags struct {
	Status              string `json:"status"`
	NumResults          int    `json:"num_results"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Items               []struct {
		Code     string `json:"code"`
		Usertags struct {
			In []struct {
				Position []float64 `json:"position"`
				User     struct {
					Username      string `json:"username"`
					Pk            int    `json:"pk"`
					ProfilePicURL string `json:"profile_pic_url"`
					IsPrivate     bool   `json:"is_private"`
					FullName      string `json:"full_name"`
				} `json:"user"`
			} `json:"in"`
		} `json:"usertags"`
		MaxNumVisiblePreviewComments int `json:"max_num_visible_preview_comments"`
		LikeCount                    int `json:"like_count"`
		ImageVersions2               struct {
			Candidates []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"candidates"`
		} `json:"image_versions2"`
		ID             string `json:"id"`
		ClientCacheKey string `json:"client_cache_key"`
		Comments       []struct {
			Status       string `json:"status"`
			UserID       int    `json:"user_id"`
			CreatedAtUtc int    `json:"created_at_utc"`
			CreatedAt    int    `json:"created_at"`
			BitFlags     int    `json:"bit_flags"`
			User         struct {
				Username      string `json:"username"`
				Pk            int    `json:"pk"`
				ProfilePicURL string `json:"profile_pic_url"`
				IsPrivate     bool   `json:"is_private"`
				FullName      string `json:"full_name"`
			} `json:"user"`
			ContentType string `json:"content_type"`
			Text        string `json:"text"`
			MediaID     int64  `json:"media_id"`
			Pk          int64  `json:"pk"`
			Type        int    `json:"type"`
		} `json:"comments"`
		DeviceTimestamp      float64 `json:"device_timestamp"`
		CommentCount         int     `json:"comment_count"`
		MediaType            int     `json:"media_type"`
		OrganicTrackingToken string  `json:"organic_tracking_token"`
		CaptionIsEdited      bool    `json:"caption_is_edited"`
		OriginalHeight       int     `json:"original_height"`
		User                 struct {
			Username                   string `json:"username"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			IsUnpublished              bool   `json:"is_unpublished"`
			FriendshipStatus           struct {
				Following       bool `json:"following"`
				OutgoingRequest bool `json:"outgoing_request"`
			} `json:"friendship_status"`
			ProfilePicURL string `json:"profile_pic_url"`
			IsFavorite    bool   `json:"is_favorite"`
			FullName      string `json:"full_name"`
			Pk            int64  `json:"pk"`
			IsPrivate     bool   `json:"is_private"`
		} `json:"user"`
		Pk              int64       `json:"pk"`
		HasLiked        bool        `json:"has_liked"`
		HasMoreComments bool        `json:"has_more_comments"`
		PhotoOfYou      bool        `json:"photo_of_you"`
		Caption         interface{} `json:"caption"`
		TakenAt         float64     `json:"taken_at"`
		OriginalWidth   int         `json:"original_width"`
		FilterType      int         `json:"filter_type,omitempty"`
		Lng             float64     `json:"lng,omitempty"`
		Lat             float64     `json:"lat,omitempty"`
		Location        struct {
			ExternalSource   string      `json:"external_source"`
			City             string      `json:"city"`
			Name             string      `json:"name"`
			FacebookPlacesID interface{} `json:"facebook_places_id"`
			ExternalID       interface{} `json:"external_id"`
			State            string      `json:"state"`
			Address          string      `json:"address"`
			Lat              float64     `json:"lat"`
			Pk               int         `json:"pk"`
			Lng              float64     `json:"lng"`
			FoursquareV2ID   string      `json:"foursquare_v2_id"`
		} `json:"location,omitempty"`
	} `json:"items"`
	MoreAvailable  bool          `json:"more_available"`
	TotalCount     int           `json:"total_count"`
	RequiresReview bool          `json:"requires_review"`
	NewPhotos      []interface{} `json:"new_photos"`
	Message        string        `json:"message"` // from Error
}

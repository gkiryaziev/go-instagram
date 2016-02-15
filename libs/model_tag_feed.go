package libs

// Get tagged media.
type TagFeed struct {
	RankedItems []struct {
		Code     string `json:"code"`
		Usertags struct {
			In []struct {
				Position []float64 `json:"position"`
				User     struct {
					Username      string `json:"username"`
					Pk            int64  `json:"pk"`
					ProfilePicURL string `json:"profile_pic_url"`
					IsPrivate     bool   `json:"is_private"`
					FullName      string `json:"full_name"`
				} `json:"user"`
			} `json:"in"`
		} `json:"usertags,omitempty"`
		MaxNumVisiblePreviewComments int `json:"max_num_visible_preview_comments"`
		LikeCount                    int `json:"like_count"`
		ImageVersions2               struct {
			Candidates []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"candidates"`
		} `json:"image_versions2"`
		ID                   string        `json:"id"`
		ClientCacheKey       string        `json:"client_cache_key"`
		Comments             []interface{} `json:"comments"`
		DeviceTimestamp      float64       `json:"device_timestamp"`
		CommentCount         int           `json:"comment_count"`
		MediaType            int           `json:"media_type"`
		OrganicTrackingToken string        `json:"organic_tracking_token"`
		CaptionIsEdited      bool          `json:"caption_is_edited"`
		OriginalHeight       int           `json:"original_height"`
		FilterType           int           `json:"filter_type"`
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
			Pk            int    `json:"pk"`
			IsPrivate     bool   `json:"is_private"`
		} `json:"user"`
		Pk              int64 `json:"pk"`
		HasLiked        bool  `json:"has_liked"`
		HasMoreComments bool  `json:"has_more_comments"`
		PhotoOfYou      bool  `json:"photo_of_you"`
		Caption         struct {
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
		} `json:"caption"`
		TakenAt       float64 `json:"taken_at"`
		OriginalWidth int     `json:"original_width"`
		Lng           float64 `json:"lng,omitempty"`
		Location      struct {
			ExternalSource   string      `json:"external_source"`
			City             string      `json:"city"`
			Name             string      `json:"name"`
			FacebookPlacesID int64       `json:"facebook_places_id"`
			ExternalID       interface{} `json:"external_id"`
			State            string      `json:"state"`
			Address          string      `json:"address"`
			Lat              float64     `json:"lat"`
			Pk               int         `json:"pk"`
			Lng              float64     `json:"lng"`
			FoursquareV2ID   string      `json:"foursquare_v2_id"`
		} `json:"location,omitempty"`
		Lat float64 `json:"lat,omitempty"`
	} `json:"ranked_items"`
	Status              string `json:"status"`
	NumResults          int    `json:"num_results"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Items               []struct {
		Code                         string `json:"code"`
		MaxNumVisiblePreviewComments int    `json:"max_num_visible_preview_comments"`
		LikeCount                    int    `json:"like_count"`
		ImageVersions2               struct {
			Candidates []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"candidates"`
		} `json:"image_versions2"`
		Lng             float64       `json:"lng,omitempty"`
		ID              string        `json:"id"`
		ClientCacheKey  string        `json:"client_cache_key"`
		Comments        []interface{} `json:"comments"`
		DeviceTimestamp float64       `json:"device_timestamp"`
		CommentCount    int           `json:"comment_count"`
		Location        struct {
			ExternalSource   string      `json:"external_source"`
			City             string      `json:"city"`
			Name             string      `json:"name"`
			FacebookPlacesID int64       `json:"facebook_places_id"`
			ExternalID       interface{} `json:"external_id"`
			State            string      `json:"state"`
			Address          string      `json:"address"`
			Lat              float64     `json:"lat"`
			Pk               int         `json:"pk"`
			Lng              float64     `json:"lng"`
			FoursquareV2ID   interface{} `json:"foursquare_v2_id"`
		} `json:"location,omitempty"`
		MediaType            int    `json:"media_type"`
		OrganicTrackingToken string `json:"organic_tracking_token"`
		CaptionIsEdited      bool   `json:"caption_is_edited"`
		OriginalHeight       int    `json:"original_height"`
		FilterType           int    `json:"filter_type"`
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
			Pk            int    `json:"pk"`
			IsPrivate     bool   `json:"is_private"`
		} `json:"user"`
		Pk              int64   `json:"pk"`
		Lat             float64 `json:"lat,omitempty"`
		HasLiked        bool    `json:"has_liked"`
		HasMoreComments bool    `json:"has_more_comments"`
		PhotoOfYou      bool    `json:"photo_of_you"`
		Caption         struct {
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
		} `json:"caption"`
		TakenAt       float64 `json:"taken_at"`
		OriginalWidth int     `json:"original_width"`
	} `json:"items"`
	MoreAvailable bool   `json:"more_available"`
	NextMaxID     string `json:"next_max_id"`
	Message       string `json:"message"` // from Error
}

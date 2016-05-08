package instagram

// Login struct
type Login struct {
	DeviceID          string `json:"device_id"`
	GUID              string `json:"guid"`
	UserName          string `json:"username"`
	Password          string `json:"password"`
	Csrftoken         string `json:"csrftoken"`
	LoginAttemptCount string `json:"login_attempt_count"`
}

// LoginResponse struct
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

package libs

// Login to Instagram.
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

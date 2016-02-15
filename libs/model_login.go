package libs

// Login to Instagram.
type Login struct {
	DeviceId          string `json:"device_id"`
	Guid              string `json:"guid"`
	UserName          string `json:"username"`
	Password          string `json:"password"`
	Csrftoken         string `json:"csrftoken"`
	LoginAttemptCount string `json:"login_attempt_count"`
}

package libs

// Search tags.
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

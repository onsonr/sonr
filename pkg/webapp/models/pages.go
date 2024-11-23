package models

type Metadata struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	OpenGraphImage string `json:"openGraphImage"`
	TwitterImage   string `json:"twitterImage"`
	Tags           string `json:"tags"`
	Category       string `json:"category"`
	Favicon        string `json:"favicon"`
	URL            string `json:"url"`
}

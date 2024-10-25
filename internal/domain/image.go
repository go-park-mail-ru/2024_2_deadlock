package domain

type ImageData struct {
	Image string `json:"image"`
}

type ImageUploadInfo struct {
	ID  ImageID  `json:"id"`
	URL ImageURL `json:"url"`
}

type ImageID string

type ImageURL string

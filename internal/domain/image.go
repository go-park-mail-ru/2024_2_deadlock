package domain

import "mime/multipart"

type ImageData struct {
	Image  multipart.File        `json:"image"`
	Header *multipart.FileHeader `json:"header"`
}

type ImageUploadInfo struct {
	ID  ImageID  `json:"id"`
	URL ImageURL `json:"url"`
}

type ImageID string

type ImageURL string

type FieldID int

package models

type Voice struct {
	Duration     int64  `json:"duration"`
	MimeType     string `json:"mime_type"`
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	Caption      string `json:"caption"`
}

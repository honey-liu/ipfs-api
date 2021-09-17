package api

type UploadFileResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	FileHash   string `json:"file_hash"`
}

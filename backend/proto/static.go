package proto

type AvatarUploadReq struct{}

type AvatarUploadResp struct{}

type DownloadRequest struct {
	Path string `json:"path"`
}

type AvatarDownloadResp struct{}

type StaticUploadRequest struct {
	Type string `json:"type" form:"type"`
}

type StaticUploadResponse struct {
	Path string `json:"path"`
}

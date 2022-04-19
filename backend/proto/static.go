package proto

type AvatarUploadReq struct{}

type AvatarUploadResp struct{}

type DownloadReq struct {
	Path string `json:"path"`
}

type AvatarDownloadResp struct{}

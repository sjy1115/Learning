package proto

type AvatarUploadReq struct{}

type AvatarUploadResp struct{}

type AvatarDownloadReq struct {
	Path string `json:"path"`
}

type AvatarDownloadResp struct{}

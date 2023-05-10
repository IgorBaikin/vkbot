package vk

type PhotosGetMessagesUploadServerResponse struct {
	Response PhotosGetMessagesUploadServer `json:"response"`
}

type PhotosGetMessagesUploadServer struct {
	AlbumID   int    `json:"album_id"`
	UploadURL string `json:"upload_url"`
	UserID    int    `json:"user_id,omitempty"`
	GroupID   int    `json:"group_id,omitempty"`
}

type PhotosMessageUploadResponse struct {
	Hash   string `json:"hash"`
	Photo  string `json:"photo"`
	Server int    `json:"server"`
}

type AttachmentsRespone struct {
	Response []Attachments `json:"response"`
}
type Attachments struct {
	OwnerID int `json:"owner_id"`
	ID      int `json:"id"`
}

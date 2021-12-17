package requests

type CreationRequest struct {
	Url string `json:"url" binding:"required"`
}

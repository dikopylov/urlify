package requests

type CreationRequest struct {
	Url string `json:"url" binding:"required"`
}

type ViewRequest struct {
	Hash string `uri:"hash" binding:"required"`
}

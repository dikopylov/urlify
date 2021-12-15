package model

type Reference struct {
	ID   int64
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

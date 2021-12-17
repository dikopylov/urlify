package model

type Reference struct {
	Url       string `json:"url" db:"url"`
	Hash      string `json:"hash" db:"hash"`
	CreatedAt string `json:"-" db:"created_at"`
}

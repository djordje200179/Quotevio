package models

type QuoteId uint64

type Quote struct {
	Id QuoteId `json:"id"`

	Text   string `json:"text"`
	Author string `json:"author"`

	Created Date `json:"created"`

	Likes    uint64 `json:"likes"`
	Dislikes uint64 `json:"dislikes"`
}

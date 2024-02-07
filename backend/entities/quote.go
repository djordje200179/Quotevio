package entities

import "time"

type Quote struct {
	Id uint `json:"id"`

	Text   string `json:"text"`
	Author string `json:"author"`

	CreatedAt time.Time `json:"created_at"`

	Likes    uint64 `json:"likes"`
	Dislikes uint64 `json:"dislikes"`
}

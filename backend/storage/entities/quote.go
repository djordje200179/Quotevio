package entities

import "time"

type Quote struct {
	Id uint `json:"id"`

	Text   string `json:"text"`
	Author string `json:"author"`

	Created time.Time `json:"created"`

	Likes    uint64 `json:"likes"`
	Dislikes uint64 `json:"dislikes"`
}

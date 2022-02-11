package models

import "time"

type Post struct {
	Id          string    `json:"id"`
	PostContent string    `json:"postContent"`
	CreatedAt   time.Time `json:"createdAt"`
	UserId      string    `json:"userId"`
}

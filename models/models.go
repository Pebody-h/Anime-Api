package models

import "gorm.io/gorm"

type Animes struct {
	gorm.Model
	Title    string  `json:"title"`
	Image    string  `json:"image"`
	Synopsis string  `json:"synopsis"`
	Trailer  string  `json:"trailer"`
	Episodes int     `json:"episodes"`
	Score    float32 `json:"score"`
	Status   string  `json:"status"`
	Year     int     `json:"year"`
}

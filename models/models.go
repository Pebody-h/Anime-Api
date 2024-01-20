package models

import "gorm.io/gorm"

type Animes struct {
	gorm.Model
	Title    string  `json:"title" gorm:"not null"`
	Image    string  `json:"image" gorm:"not null"`
	Synopsis string  `json:"synopsis" gorm:"not null"`
	Trailer  string  `json:"trailer" gorm:"not null"`
	Episodes int     `json:"episodes" gorm:"not null"`
	Score    float32 `json:"score" gorm:"not null"`
	Status   string  `json:"status" gorm:"not null"`
	Year     int     `json:"year" gorm:"not null"`
}

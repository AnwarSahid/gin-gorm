package models

import "time"

type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null; type:varchar(100)"`
	Author      string `gorm:"not null; type:varchar(100)"`
	Description string `gorm:"not null; type:varchar(255)"`
	Created_at  time.Time
	Updated_at  time.Time
}

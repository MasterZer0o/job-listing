package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"type:SERIAL;primaryKey"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	FirstName string    `gorm:"type:varchar(100)"`
	LastName  string    `gorm:"type:varchar(100)"`
	Password  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
}

type Session struct {
	ID        string `gorm:"type:varchar(255);primaryKey"`
	UserID    User   `gorm:"TYPE:integer REFERENCES users;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time
}

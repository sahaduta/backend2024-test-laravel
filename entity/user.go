package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	Slug         string
	IsProject    bool
	SelfCapture  string
	ClientPrefix string
	ClientLogo   string
	Address      string
	PhoneNumber  string
	City         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

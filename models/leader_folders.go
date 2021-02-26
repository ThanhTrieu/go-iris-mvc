package models

import (
	"time"
)

type LeaderFolders struct {
	ID       uint
	UserId   int64
	GroupId  int64
	LeaderName  string `gorm:"size:64"`
	LeaderEmail string `gorm:"size:100"`
	LeaderTelegram string `gorm:"size:100"`
	FloderName string `gorm:"size:1024"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
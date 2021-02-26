package models

import (
	"time"
)

type FileListFolder struct {
	ID          uint
	DataType    string `gorm:"size:64"`
	FolderName  string `gorm:"size:255"`
	FileName    string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
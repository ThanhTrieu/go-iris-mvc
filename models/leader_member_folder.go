package models

type LeaderMemberFolders struct {
	ID       uint
	UserId   int64
	GroupId  int64
	LeaderName  string `gorm:"size:64"`
	LeaderEmail string `gorm:"size:100"`
	LeaderTelegram string `gorm:"size:100"`
	FloderName string `gorm:"size:1024"`
	MemberID uint
	MemberName  string `gorm:"size:64"`
	MemberEmail string `gorm:"size:100"`
	MemberTelegram string `gorm:"size:100"`
	MemberFolder string `gorm:"size:1024"`
}
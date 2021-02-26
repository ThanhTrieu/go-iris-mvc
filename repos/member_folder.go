package repos

import (
	"gomvc/models"
	"time"

	"gorm.io/gorm"
)

type MemberFolderRepository interface {
	Select(query string) []models.MemberFolder
	SelectById(query string, id int64) models.MemberFolder
	SelectByLeaderId(query string, id int64) []models.MemberFolder
	CreateFolderMember(query string, leaderID int64, memberName string, memberEmail string, memberTelegram string, memberFolder string, ctTime time.Time, upTime time.Time, deTime time.Time) bool
}

type memberFolderMysqlRepository struct {
	DB *gorm.DB
}

func NewMemberFolderRepository(db *gorm.DB) MemberFolderRepository {
	return &memberFolderMysqlRepository{DB: db}
}

func (m *memberFolderMysqlRepository) Select(query string) []models.MemberFolder {
	result := []models.MemberFolder{}
	m.DB.Raw(query).Scan(&result)
	return result
}

func (m *memberFolderMysqlRepository) SelectByLeaderId(query string, id int64) []models.MemberFolder {
	result := []models.MemberFolder{}
	m.DB.Raw(query, id).Scan(&result)
	return result
}

func (m *memberFolderMysqlRepository) SelectById(query string, id int64) models.MemberFolder {
	result := models.MemberFolder{}
	m.DB.Raw(query, id).Scan(&result)
	return result
}

func (m *memberFolderMysqlRepository) CreateFolderMember(query string, leaderID int64, memberName string, memberEmail string, memberTelegram string, memberFolder string, ctTime time.Time, upTime time.Time, deTime time.Time) bool {
	tx := m.DB.Exec(query, leaderID, memberName, memberEmail, memberTelegram, memberFolder, ctTime, upTime, deTime)
	if tx.Error != nil {
    return false
	}
	return true
}

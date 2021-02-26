package repos

import (
	"gomvc/models"
	"time"

	"gorm.io/gorm"
)

type LeaderFolderRepository interface {
	Select(query string) []models.LeaderFolders
	SelectAllByRole(query string) []models.LeaderFolders
	SelectByRoleByID(query string, id int64) []models.LeaderFolders
	SelectById(query string, id int64) models.LeaderFolders
	CreateFolderLeader(userID int64, groupID int64, leaderName string, leaderEmail string, leaderTelegram string, floderName string) int64
}

type leaderFolderMysqlRepository struct {
	DB *gorm.DB
}

func NewLeaderFolderRepository(db *gorm.DB) LeaderFolderRepository {
	return &leaderFolderMysqlRepository{DB: db}
}

func (m *leaderFolderMysqlRepository) Select(query string) []models.LeaderFolders {
	result := []models.LeaderFolders{}
	m.DB.Raw(query).Scan(&result)
	return result
}

func (m *leaderFolderMysqlRepository) SelectAllByRole(query string) []models.LeaderFolders {
	result := []models.LeaderFolders{}
	m.DB.Raw(query).Scan(&result)
	return result
}

func (m *leaderFolderMysqlRepository) SelectByRoleByID(query string, id int64) []models.LeaderFolders {
	result := []models.LeaderFolders{}
	m.DB.Raw(query, id).Scan(&result)
	return result
}

func (m *leaderFolderMysqlRepository) SelectById(query string, id int64) models.LeaderFolders {
	result := models.LeaderFolders{}
	m.DB.Raw(query, id).Scan(&result)
	return result
}

func (m *leaderFolderMysqlRepository) CreateFolderLeader(userID int64, groupID int64, leaderName string, leaderEmail string, leaderTelegram string, floderName string) int64 {
	var dataInsert = models.LeaderFolders{
		UserId: userID,
		GroupId: groupID,
		LeaderName: leaderName,
		LeaderEmail: leaderEmail,
		LeaderTelegram: leaderTelegram,
		FloderName: floderName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	}
	m.DB.Create(&dataInsert)
	lastID := int64(dataInsert.ID)
	return lastID
}

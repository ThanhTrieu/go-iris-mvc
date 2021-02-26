package services

import (
	"gomvc/models"
	"gomvc/repos"
)

type LeaderFoldersService interface {
	GetAll() []models.LeaderFolders
	GetByRole(role int64, userID int64) []models.LeaderFolders
	GetByID(id int64) models.LeaderFolders
	CreateFolder(userID int64, groupID int64, leaderName string, leaderEmail string, leaderTelegram string, floderName string)  int64
}

type leaderfoldersService struct {
	repo repos.LeaderFolderRepository
}

func NewLeaderFoldersService(repo repos.LeaderFolderRepository) LeaderFoldersService {
	return &leaderfoldersService{
		repo: repo,
	}
}

func (s *leaderfoldersService) GetByRole(role int64, userID int64) []models.LeaderFolders {
	if role == -1 {
		return s.repo.SelectAllByRole("SELECT * FROM leader_folders")
	}
	return s.repo.SelectByRoleByID("SELECT * FROM leader_folders WHERE user_id=?", userID)
}

func (s *leaderfoldersService) GetAll() []models.LeaderFolders {
	return s.repo.Select("select * from leader_folders")
}

func (s *leaderfoldersService) GetByID(id int64) models.LeaderFolders {
	return s.repo.SelectById("select * from leader_folders where id=?", id)
}

func (s *leaderfoldersService) CreateFolder(userID int64, groupID int64, leaderName string, leaderEmail string, leaderTelegram string, floderName string) int64 {
	return s.repo.CreateFolderLeader(userID, groupID, leaderName, leaderEmail, leaderTelegram, floderName)
}


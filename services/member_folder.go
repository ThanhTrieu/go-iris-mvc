package services

import (
	"gomvc/models"
	"gomvc/repos"
	"time"
)

type MemberFolderService interface {
	GetAll() []models.MemberFolder
	GetByIDLeader(id int64) []models.MemberFolder
	GetByID(id int64) models.MemberFolder
	CreateFolder(leaderID int64, memberName string, memberEmail string, memberTelegram string, memberFolder string)  bool
}

type memberfoldersService struct {
	repo repos.MemberFolderRepository
}

func NewMemberFoldersService(repo repos.MemberFolderRepository) MemberFolderService {
	return &memberfoldersService{
		repo: repo,
	}
}

func (s *memberfoldersService) GetAll() []models.MemberFolder {
	return s.repo.Select("select * from member_folders")
}

func (s *memberfoldersService) GetByIDLeader(id int64) []models.MemberFolder {
	return s.repo.SelectByLeaderId("select * from member_folders where leader_id=?", id)
}

func (s *memberfoldersService) GetByID(id int64) models.MemberFolder {
	return s.repo.SelectById("select * from member_folders where id=?", id)
}

func (s *memberfoldersService) CreateFolder(leaderID int64, memberName string, memberEmail string, memberTelegram string, memberFolder string) bool {
	creatTime := time.Now()
	updateTime := time.Time{}
	delTime := time.Time{}
	return s.repo.CreateFolderMember("INSERT INTO member_folders(leader_id, member_name, member_email, member_telegram, member_folder,created_at, updated_at, deleted_at) VALUES(?,?,?,?,?,?,?,?)", leaderID, memberName, memberEmail, memberTelegram, memberFolder, creatTime, updateTime, delTime)
}


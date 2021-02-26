package services

import (
	"gomvc/models"
	"gomvc/repos"
)

type FileListFolderService interface {
	GetAll() []models.FileListFolder
	GetByID(id int64) models.FileListFolder
}

type memberListFoldersService struct {
	repo repos.FileListFolderRepository
}

func NewFileListFoldersService(repo repos.FileListFolderRepository) FileListFolderService {
	return &memberListFoldersService{
		repo: repo,
	}
}

func (s *memberListFoldersService) GetAll() []models.FileListFolder {
	return s.repo.Select("select * from file_list_folder")
}

func (s *memberListFoldersService) GetByID(id int64) models.FileListFolder {
	return s.repo.SelectById("select * from file_list_folder where id=?", id)
}


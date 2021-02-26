package repos

import (
	"gomvc/models"

	"gorm.io/gorm"
)

type FileListFolderRepository interface {
	Select(query string) []models.FileListFolder
	SelectById(query string, id int64) models.FileListFolder
}

type fileListFolderMysqlRepository struct {
	DB *gorm.DB
}

func NewFileListFolderRepository(db *gorm.DB) FileListFolderRepository {
	return &fileListFolderMysqlRepository{DB: db}
}

func (m *fileListFolderMysqlRepository) Select(query string) []models.FileListFolder {
	result := []models.FileListFolder{}
	m.DB.Raw(query).Scan(&result)
	return result
}

func (m *fileListFolderMysqlRepository) SelectById(query string, id int64) models.FileListFolder {
	result := models.FileListFolder{}
	m.DB.Raw(query, id).Scan(&result)
	return result
}

package repos

import (
	"gomvc/models"

	"gorm.io/gorm"
)

type GroupRepository interface {
	Select(query string) []models.Groups
	SelectById(query string, id int64) models.Groups
}

type groupMysqlRepository struct {
	DB *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupMysqlRepository{DB: db}
}

func (m *groupMysqlRepository) Select(query string) []models.Groups {
	result := []models.Groups{}
	m.DB.Raw(query).Scan(&result)
	return result
}

func (m *groupMysqlRepository) SelectById(query string, id int64) models.Groups {
	result := models.Groups{}
	m.DB.Raw(query, id).Scan(&result)
	return result
}
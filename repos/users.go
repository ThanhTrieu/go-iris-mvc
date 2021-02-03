package repos

import (
	"gomvc/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Select(query string) []models.Users
	SelectById(query string, id int64) models.Users
	SelectByName(query string, name string) models.Users
	LoginUser(query string, username string, password string) models.Users
}

type userMysqlRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userMysqlRepository{DB: db}
}

func (m *userMysqlRepository) Select(query string) []models.Users {
	result := []models.Users{}
	m.DB.Raw(query).Scan(&result)
	return result
}

func (m *userMysqlRepository) SelectById(query string, id int64) models.Users {
	result := models.Users{}
	m.DB.Raw(query, id).Scan(&result)
	return result
}

func (m *userMysqlRepository) LoginUser(query string, username string, password string) models.Users {
	result := models.Users{}
	m.DB.Raw(query, username, password).Scan(&result)
	return result
}

func (m *userMysqlRepository) SelectByName(query string, name string) models.Users {
	result := models.Users{}
	m.DB.Raw(query, name).Scan(&result)
	return result
}

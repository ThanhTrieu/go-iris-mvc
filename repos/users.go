package repos

import (
	"gomvc/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Select(query string) []models.Users
	SelectById(query string, id int64) models.Users
	SelectByName(query string, name string) models.Users
	SelectByEmail(query string, email string) models.Users
	LoginUser(query string, username string) models.Users
	CreateAccount(query string, username string, password string, email string, phone string, role int) bool
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

func (m *userMysqlRepository) LoginUser(query string, username string) models.Users {
	result := models.Users{}
	m.DB.Raw(query, username).Scan(&result)
	return result
}

func (m *userMysqlRepository) SelectByName(query string, name string) models.Users {
	result := models.Users{}
	m.DB.Raw(query, name).Scan(&result)
	return result
}

func (m *userMysqlRepository) SelectByEmail(query string, email string) models.Users {
	result := models.Users{}
	m.DB.Raw(query, email).Scan(&result)
	return result
}

func (m *userMysqlRepository) CreateAccount(query string, username string, password string, email string, phone string, role int) bool {
	tx := m.DB.Exec(query, username, password, email, phone, role)
	if tx.Error != nil {
    return false
	}
	return true
}

package services

import (
	"gomvc/models"
	"gomvc/repos"
)

type UsersService interface {
	GetAll() []models.Users
	GetByID(id int64) models.Users
	GetByName(name string) models.Users
	GetByEmail(email string) models.Users
	CheckUsernameExists(username string) bool
	CheckEmailExists(email string) bool
	CheckLoginUser(username string, password string) (models.Users, bool)


	CreateUser(username string, password string, email string, phone string)  bool
}

type usersService struct {
	repo repos.UserRepository
}

func NewUserService(repo repos.UserRepository) UsersService {
	return &usersService{
		repo: repo,
	}
}

func (s *usersService) GetAll() []models.Users {
	return s.repo.Select("select * from users")
}

func (s *usersService) GetByID(id int64) models.Users {
	return s.repo.SelectById("select * from users where id=?", id)
}

func (s *usersService) CheckLoginUser(username string, password string) (models.Users, bool) {
	if username == "" || password == "" {
		return models.Users{}, false
	}
	dataUser := s.repo.LoginUser("select * from users as u where u.username=? and u.password=? ", username, password)
	if (models.Users{}) == dataUser {
		return models.Users{}, false
	}
	return dataUser, true
}

func (s *usersService) GetByName(name string) models.Users {
	return s.repo.SelectByName("select * from users where username=?", name)
}

func (s *usersService) GetByEmail(email string) models.Users {
	return s.repo.SelectByEmail("select * from users where email=?", email)
}

func (s *usersService) CheckUsernameExists(username string) bool {
	if username == "" {
		return true
	}
	info := s.GetByName(username)
	if (models.Users{}) == info {
		return false
	}
	return true
}

func (s *usersService) CheckEmailExists(email string) bool  {
	if email == "" {
		return true
	}
	infoEmail := s.GetByEmail(email)
	if (models.Users{}) == infoEmail {
		return false
	}
	return true
}

func (s *usersService) CreateUser(username string, password string, email string, phone string) bool {
	// return s.repo.CreateAccount("INSERT INTO users(username, password, email, phone, authen_key, fullname, address, brirthday, status, gender, last_login, created_at, updated_at, deleted_at) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)", username, password, email, phone, "", "", "", "", 1, 1, "", "", "", "")
	return s.repo.CreateAccount("INSERT INTO users(username, password, email, phone) VALUES(?,?,?,?)", username, password, email, phone)
}


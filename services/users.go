package services

import (
	"gomvc/models"
	"gomvc/repos"
)

type UsersService interface {
	GetAll() []models.Users
	GetByID(id int64) models.Users
	GetByName(name string) models.Users
	CheckLoginUser(username string, password string) (models.Users, bool)
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


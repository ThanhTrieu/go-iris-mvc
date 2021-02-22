package services

import (
	"gomvc/models"
	"gomvc/repos"
)

type GroupsService interface {
	GetAll() []models.Groups
	GetByID(id int64) models.Groups
}

type groupsService struct {
	repo repos.GroupRepository
}

func NewGroupService(repo repos.GroupRepository) GroupsService {
	return &groupsService{
		repo: repo,
	}
}

func (s *groupsService) GetAll() []models.Groups {
	return s.repo.Select("select * from groups")
}

func (s *groupsService) GetByID(id int64) models.Groups {
	return s.repo.SelectById("select * from groups where id=?", id)
}

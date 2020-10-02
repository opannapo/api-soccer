package services

import (
	"app/app/entities"
	"app/app/injection/repositoryinjection"
)

type PlayerService struct {
	Repo *repositoryinjection.RepositoryInjection
}

func (instance *PlayerService) GetAll() (result []*entities.PlayerEntity, err error) {
	result, err = instance.Repo.PlayerRepositories.GetAll()
	return
}

func (instance *PlayerService) GetByTeam(teamId int) (result []*entities.PlayerEntity, err error) {
	result, err = instance.Repo.PlayerRepositories.GetByTeam(teamId)
	return
}

func (instance *PlayerService) GetOne(playerId int) (result entities.PlayerEntity, err error) {
	result, err = instance.Repo.PlayerRepositories.GetOne(playerId)
	return
}

func NewInstancePlayerService(repositoryInjected *repositoryinjection.RepositoryInjection) PlayerServices {
	return &PlayerService{
		Repo: repositoryInjected,
	}
}

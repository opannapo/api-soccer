package services

import (
	"app/app/entities"
	"app/app/injection/repositoryinjection"
)

type TeamService struct {
	Repo *repositoryinjection.RepositoryInjection
}

func (instance *TeamService) GetById(id int) (result []*entities.TeamEntity, err error) {
	result, err = instance.Repo.TeamRepositories.GetById(id)
	return
}

func (instance *TeamService) GetAll() (result []*entities.TeamEntity, err error) {
	result, err = instance.Repo.TeamRepositories.GetAll()
	return
}

func NewInstanceTimeService(repositoryInjected *repositoryinjection.RepositoryInjection) TeamServices {
	return &TeamService{
		Repo: repositoryInjected,
	}
}

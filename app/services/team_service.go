package services

import (
	"app/app/entities"
	"app/app/injection/repositoryinjection"
)

type TeamService struct {
	Repo *repositoryinjection.RepositoryInjection
}

func (instance *TeamService) GetAll() (result []*entities.TeamEntity, err error) {
	result, err = instance.Repo.TeamRepositories.GetAll()
	return
}

func (instance *TeamService) GetOne(id int) (result entities.TeamEntity, err error) {
	panic("implement me")
}

func NewInstanceTimeService(repositoryInjected *repositoryinjection.RepositoryInjection) TeamServices {
	return &TeamService{
		Repo: repositoryInjected,
	}
}

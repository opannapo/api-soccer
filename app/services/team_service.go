package services

import (
	"app/app/apis/param"
	"app/app/entities"
	"app/app/injection/repositoryinjection"
)

type TeamService struct {
	Repo *repositoryinjection.RepositoryInjection
}

func (instance *TeamService) Create(param *param.TeamCreateParameter) (err error) {
	txCreateTeams, err := instance.Repo.TeamRepositories.Create(param)
	if err != nil && txCreateTeams != nil {
		txCreateTeams.Rollback()
		return
	}

	//modify data player
	var players = param.Players
	for i := range players {
		players[i].TeamId = param.Team.ID
	}
	txCreatePlayer, err := instance.Repo.PlayerRepositories.Create(param.Players)
	if err != nil && txCreatePlayer != nil {
		txCreateTeams.Rollback()
		txCreatePlayer.Rollback()
		return
	}

	if txCreateTeams != nil && txCreatePlayer != nil {
		txCreateTeams.Commit()
		txCreatePlayer.Commit()
		return
	}

	return
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

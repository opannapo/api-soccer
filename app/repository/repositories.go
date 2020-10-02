package repository

import "app/app/entities"

type TeamRepositories interface {
	GetAll() (result []*entities.TeamEntity, err error)
	GetById(teamId int) (result []*entities.TeamEntity, err error)
}

type PlayerRepositories interface {
	GetAll() (result []*entities.PlayerEntity, err error)
	GetByTeam(teamId int) (result []*entities.PlayerEntity, err error)
	GetOne(playerId int) (result entities.PlayerEntity, err error)
}

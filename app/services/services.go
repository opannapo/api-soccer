package services

import (
	"app/app/apis/param"
	"app/app/entities"
)

type TeamServices interface {
	GetAll() (result []*entities.TeamEntity, err error)
	GetById(teamId int) (result []*entities.TeamEntity, err error)
	Create(param *param.TeamCreateParameter) (err error)
}

type PlayerServices interface {
	GetAll() (result []*entities.PlayerEntity, err error)
	GetByTeam(playerId int) (result []*entities.PlayerEntity, err error)
	GetOne(playerId int) (result entities.PlayerEntity, err error)
}

package services

import "app/app/entities"

type TeamServices interface {
	GetAll() (result []*entities.TeamEntity, err error)
	GetOne(id int) (result entities.TeamEntity, err error)
}

type PlayerServices interface {
	GetAll() (result []*entities.PlayerEntity, err error)
	GetByTeam(teamId int) (result []*entities.PlayerEntity, err error)
	GetOne(id int) (result entities.PlayerEntity, err error)
}

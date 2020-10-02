package repository

import (
	"app/app/apis/param"
	"app/app/entities"
	"github.com/jinzhu/gorm"
)

type TeamRepositories interface {
	GetAll() (result []*entities.TeamEntity, err error)
	GetById(teamId int) (result []*entities.TeamEntity, err error)
	Create(param *param.TeamCreateParameter) (tx *gorm.DB, err error)
}

type PlayerRepositories interface {
	GetAll() (result []*entities.PlayerEntity, err error)
	GetByTeam(teamId int) (result []*entities.PlayerEntity, err error)
	GetOne(playerId int) (result entities.PlayerEntity, err error)
	Create(param []*entities.PlayerEntity) (tx *gorm.DB, err error)
}

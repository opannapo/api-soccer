package repository

import (
	"app/app/entities"
	"github.com/jinzhu/gorm"
)

type PlayerRepository struct {
	Db *gorm.DB
}

func (instance *PlayerRepository) GetAll() (result []*entities.PlayerEntity, err error) {
	err = instance.Db.
		Preload("Team").
		Find(&result).
		Error
	return
}

func (instance *PlayerRepository) GetByTeam(teamId int) (result []*entities.PlayerEntity, err error) {
	err = instance.Db.
		Preload("Team").
		Where("team_id=?", teamId).
		Find(&result).
		Error
	return
}

func (instance *PlayerRepository) GetOne(playerId int) (result entities.PlayerEntity, err error) {
	err = instance.Db.
		Preload("Team").
		Where("id=?", playerId).
		First(&result).
		Error
	return
}

func NewInstancePlayerRepository(db *gorm.DB) PlayerRepositories {
	return &PlayerRepository{Db: db}
}

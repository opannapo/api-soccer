package repository

import (
	"app/app/entities"
	"github.com/jinzhu/gorm"
)

type PlayerRepository struct {
	Db *gorm.DB
}

func (instance *PlayerRepository) Create(param []*entities.PlayerEntity) (tx *gorm.DB, err error) {
	tx = instance.Db.Begin()

	for i := range param {
		err = tx.Create(&param[i]).Error
		if err != nil {
			break
		}
	}

	return
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

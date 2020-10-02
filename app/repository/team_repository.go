package repository

import (
	"app/app/entities"
	"github.com/jinzhu/gorm"
)

type TeamRepository struct {
	Db *gorm.DB
}

func (instance *TeamRepository) GetById(id int) (result []*entities.TeamEntity, err error) {
	err = instance.Db.
		Preload("Players").
		Where("id=?", id).
		First(&result).
		Error

	return
}

func (instance *TeamRepository) GetAll() (result []*entities.TeamEntity, err error) {
	err = instance.Db.
		Find(&result).
		Error
	return
}

func NewInstanceTeamRepository(db *gorm.DB) TeamRepositories {
	return &TeamRepository{Db: db}
}

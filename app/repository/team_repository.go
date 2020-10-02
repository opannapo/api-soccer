package repository

import (
	"app/app/entities"
	"github.com/jinzhu/gorm"
)

type TeamRepository struct {
	Db *gorm.DB
}

func (instance *TeamRepository) GetAll() (result []*entities.TeamEntity, err error) {
	err = instance.Db.Preload("Players").Find(&result).Error
	return
}

func (instance *TeamRepository) GetOne(id int) (result entities.TeamEntity, err error) {
	panic("implement me")
}

func NewInstanceTeamRepository(db *gorm.DB) TeamRepositories {
	return &TeamRepository{Db: db}
}

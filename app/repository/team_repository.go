package repository

import (
	"app/app/apis/param"
	"app/app/entities"
	"github.com/jinzhu/gorm"
)

type TeamRepository struct {
	Db *gorm.DB
}

func (instance *TeamRepository) Create(param *param.TeamCreateParameter) (tx *gorm.DB, err error) {
	data := param.Team
	tx = instance.Db.Begin()
	err = tx.Create(data).Error
	return
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

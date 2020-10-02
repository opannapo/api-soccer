package repository

import "app/app/entities"

type TeamRepositories interface {
	GetAll() (result []*entities.TeamEntity, err error)
	GetOne(id int) (result entities.TeamEntity, err error)
}

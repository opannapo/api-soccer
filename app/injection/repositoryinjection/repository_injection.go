package repositoryinjection

import (
	"app/app/repository"
	"github.com/jinzhu/gorm"
)

type RepositoryInjection struct {
	*mysqlRepoInjected
	db *gorm.DB
}

type mysqlRepoInjected struct {
	TeamRepositories   repository.TeamRepositories
	PlayerRepositories repository.PlayerRepositories
}

//NewInstanceRepositoryInjection new instance of RepositoryInjection struct
func NewInstanceRepositoryInjection(db *gorm.DB) *RepositoryInjection {
	msql := mysqlRepoInjected{
		TeamRepositories:   repository.NewInstanceTeamRepository(db),
		PlayerRepositories: repository.NewInstancePlayerRepository(db),
	}

	return &RepositoryInjection{
		db:                db,
		mysqlRepoInjected: &msql,
	}
}

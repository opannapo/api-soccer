package repositoryinjection

import "github.com/jinzhu/gorm"

type RepositoryInjection struct {
	*mysqlRepoInjected
	db *gorm.DB
}

type mysqlRepoInjected struct {
}

//NewInstanceRepositoryInjection new instance of RepositoryInjection struct
func NewInstanceRepositoryInjection(db *gorm.DB) *RepositoryInjection {
	return &RepositoryInjection{
		db: db,
	}
}

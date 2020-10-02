package servicesinjection

import (
	"app/app/injection/repositoryinjection"
	"app/app/services"
)

//ServiceInjection struct
type ServiceInjection struct {
	*servicesInjected
	repository *repositoryinjection.RepositoryInjection
}

type servicesInjected struct {
	TeamService   services.TeamServices
	PlayerService services.PlayerServices
}

//NewInstanceServiceInjection new instance of ServiceInjection, & generate all services Instance
func NewInstanceServiceInjection(repository *repositoryinjection.RepositoryInjection) *ServiceInjection {
	service := servicesInjected{
		TeamService:   services.NewInstanceTimeService(repository),
		PlayerService: services.NewInstancePlayerService(repository),
	}

	return &ServiceInjection{
		servicesInjected: &service,
		repository:       repository,
	}
}

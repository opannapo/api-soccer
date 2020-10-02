package servicesinjection

import "app/app/injection/repositoryinjection"

//ServiceInjection struct
type ServiceInjection struct {
	*servicesInjected
	repository *repositoryinjection.RepositoryInjection
}

type servicesInjected struct {
}

//NewInstanceServiceInjection new instance of ServiceInjection, & generate all services Instance
func NewInstanceServiceInjection(repository *repositoryinjection.RepositoryInjection) *ServiceInjection {
	ms := servicesInjected{}

	return &ServiceInjection{
		servicesInjected: &ms,
		repository:       repository,
	}
}

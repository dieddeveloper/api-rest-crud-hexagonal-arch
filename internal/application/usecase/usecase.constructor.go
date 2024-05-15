package usecase

import "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/application/port"

type ServiceStruc struct {
	servicesStruc port.IServices
}

func NewUseCaseConstructor(servicesStruc port.IServices) *ServiceStruc {
	return &ServiceStruc{
		servicesStruc: servicesStruc,
	}
}

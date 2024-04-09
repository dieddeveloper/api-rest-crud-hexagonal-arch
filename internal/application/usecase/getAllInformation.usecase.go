package usecase

import "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"

func (services *ServiceStruc) GetAllPersonInformationUseCase() ([]*dtos.PersonDTO, error) {
	return services.servicesStruc.GetAllPersonInformationServices()
}

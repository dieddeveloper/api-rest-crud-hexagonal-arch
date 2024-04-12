package usecase

import "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"

func (services *ServiceStruc) GetPersonByIDUseCase(personID int64) (*dtos.PersonDTO, error) {
	return services.servicesStruc.GetPersonByIDService(personID)
}

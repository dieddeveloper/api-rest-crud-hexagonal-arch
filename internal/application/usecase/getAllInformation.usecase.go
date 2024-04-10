package usecase

import "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"

func (services *ServiceStruc) GetAllPersonInformationUseCase(requestMetadata dtos.RequestInformationMetadata) ([]*dtos.PersonDTO, dtos.PaginationMetadataResponse, error) {
	return services.servicesStruc.GetAllPersonInformationServices(requestMetadata)
}

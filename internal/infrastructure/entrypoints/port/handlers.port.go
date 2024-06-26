package port

import "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"

type IHandlersUseCases interface {
	GetAllPersonInformationUseCase(requestMetadata dtos.RequestInformationMetadata) ([]*dtos.PersonDTO, dtos.PaginationMetadataResponse, error)
	GetPersonByIDUseCase(cardNumber string) (*dtos.PersonDTO, error)
	CreatePersonUseCase(personDTO *dtos.PersonDTO) (int64, error)
}

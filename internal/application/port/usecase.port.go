package port

import "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"

type IServices interface {
	GetAllPersonInformationService(requestMetadata dtos.RequestInformationMetadata) ([]*dtos.PersonDTO, dtos.PaginationMetadataResponse, error)
	GetPersonByIDService(cardNumber string) (*dtos.PersonDTO, error)
	CreatePersonService(personDTO dtos.PersonDTO) (int64, error)
}

package port

import (
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/adapters/models"
)

type IAdapters interface {
	GetAllPersonInformationAdapter(requestMetadata dtos.RequestInformationMetadata) ([]*models.PersonModel, dtos.PaginationMetadataResponse, error)
	GetPersonByIDAdapter(cardNumber string) (*models.PersonModel, error)
	CreatePersonAdapter(personDTO dtos.PersonDTO) (int64, error)
}

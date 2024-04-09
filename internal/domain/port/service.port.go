package port

import "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/adapters/models"

type IAdapters interface {
	GetAllPersonInformationAdapter() ([]*models.PersonModel, error)
}

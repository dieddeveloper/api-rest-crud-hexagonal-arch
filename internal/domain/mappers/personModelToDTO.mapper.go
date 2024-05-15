package mappers

import (
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/adapters/models"
)

func PersonModelToDTOMapper(modelToMap *models.PersonModel) *dtos.PersonDTO {
	return &dtos.PersonDTO{
		ID:   int64(modelToMap.ID),
		Name: modelToMap.Name,
		Age:  int64(modelToMap.Age),
	}
}

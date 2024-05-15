package mappers

import (
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/adapters/models"
)

func PersonModelToDTOMapperArray(modelToMap []*models.PersonModel) []*dtos.PersonDTO {
	var resultDtoSlice []*dtos.PersonDTO
	var resultDTO dtos.PersonDTO
	for _, item := range modelToMap {
		resultDTO.ID = int64(item.ID)
		resultDTO.Name = item.Name
		resultDTO.Age = int64(item.Age)

		resultDtoSlice = append(resultDtoSlice, &resultDTO)
	}
	return resultDtoSlice
}

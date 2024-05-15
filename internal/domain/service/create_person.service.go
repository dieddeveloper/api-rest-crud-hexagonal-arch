package service

import (
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
)

func (adapters *AdaptersStruc) CreatePersonService(personDTO dtos.PersonDTO) (int64, error) {
	return adapters.adaptersStruct.CreatePersonAdapter(personDTO)
}

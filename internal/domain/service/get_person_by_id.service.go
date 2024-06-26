package service

import (
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/mappers"
	"github.com/sirupsen/logrus"
)

func (adapters *AdaptersStruc) GetPersonByIDService(cardNumber string) (*dtos.PersonDTO, error) {
	result, err := adapters.adaptersStruct.GetPersonByIDAdapter(cardNumber)
	if err != nil {
		logrus.Error("there is an error getting person information in service adapter ", err)
		return nil, err
	}
	return mappers.PersonModelToDTOMapper(result), nil
}

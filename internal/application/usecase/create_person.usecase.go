package usecase

import (
	"errors"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/sirupsen/logrus"
)

func (services *ServiceStruc) CreatePersonUseCase(personDTO *dtos.PersonDTO) (int64, error) {

	personValidated, err := services.servicesStruc.GetPersonByIDService(personDTO.CardNumber)
	if personValidated != nil || err != nil {
		logrus.Error("there is an error getting person information in service adapter ", err)
		return 0, errors.New("person already exists")
	}
	resultID, err := services.servicesStruc.CreatePersonService(personDTO)
	if err != nil {
		logrus.Error("there is an error getting person information in service adapter ", err)
		return 0, err
	}
	return resultID, nil
}

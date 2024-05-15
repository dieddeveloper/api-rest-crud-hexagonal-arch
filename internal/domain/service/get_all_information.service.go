package service

import (
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/mappers"
	"github.com/sirupsen/logrus"
)

func (adapters *AdaptersStruc) GetAllPersonInformationService(requestMetadata dtos.RequestInformationMetadata) ([]*dtos.PersonDTO, dtos.PaginationMetadataResponse, error) {
	result, pagination, err := adapters.adaptersStruct.GetAllPersonInformationAdapter(requestMetadata)
	if err != nil {
		logrus.Error("there is an error getting person information in service adapter ", err)
		return nil, dtos.PaginationMetadataResponse{}, err
	}
	return mappers.PersonModelToDTOMapperArray(result), pagination, nil
}

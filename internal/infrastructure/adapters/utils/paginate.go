package utils

import (
	"math"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func PaginateScope(paginationDto dtos.RequestInformationMetadata) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (paginationDto.Offset - 1) * paginationDto.Limit
		return db.Offset(int(offset)).Limit(int(paginationDto.Limit))
	}
}

func Paginate(query *gorm.DB, paginationDto dtos.RequestInformationMetadata, result interface{}) (dtos.PaginationMetadataResponse, error) {
	var totalItems int64
	err := query.Count(&totalItems).Error
	if err != nil {
		logrus.Error("error counting total elements: ", err)
		return dtos.PaginationMetadataResponse{}, err
	}

	pageNumber := paginationDto.Offset
	if pageNumber == 0 {
		pageNumber = 1
	}

	pageSize := paginationDto.Limit
	if pageSize == 0 {
		pageSize = 20
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	hasNextPage := int64(pageNumber*pageSize) < totalItems
	hasPreviousPage := pageNumber > 1

	err = query.
		Scopes(PaginateScope(dtos.RequestInformationMetadata{
			Limit:  pageNumber,
			Offset: pageSize,
		})).
		Find(result).Error
	if err != nil {
		logrus.Error("error looking for paginated results: ", err)
		return dtos.PaginationMetadataResponse{}, err
	}
	return dtos.PaginationMetadataResponse{
		TotalItems:      int(totalItems),
		TotalPages:      totalPages,
		CurrentPage:     int(pageNumber),
		PageSize:        int(pageSize),
		HasNextPage:     hasNextPage,
		HasPreviousPage: hasPreviousPage,
	}, nil
}

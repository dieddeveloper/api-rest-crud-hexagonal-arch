package helpers

import (
	"math"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Paginator struct{}

func NewPaginator() *Paginator {
	return &Paginator{}
}

// PaginateScope crea un closure que puede ser utilizado por GORM para aplicar la paginación en la consulta.
func (p *Paginator) PaginateScope(paginationDto dtos.PaginationMetadataResponse) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (paginationDto.PageNumber - 1) * paginationDto.PageSize
		return db.Offset(offset).Limit(paginationDto.PageSize)
	}
}

// Paginate aplica la paginación a la consulta proporcionada y la ejecuta, luego devuelve un objeto de respuesta de paginación.
func (p *Paginator) Paginate(query *gorm.DB, paginationDto dtos.RequestInformationMetadata, result interface{}) (dtos.PaginationMetadataResponse, error) {
	var totalItems int64
	err := query.Count(&totalItems).Error
	if err != nil {
		logrus.Error("Error al contar los elementos totales: ", err)
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

	// Resultados paginados
	err = query.
		Scopes(p.PaginateScope(dtos.PaginationRequest{
			PageNumber: pageNumber,
			PageSize:   pageSize,
		})).
		Find(result).Error
	if err != nil {
		logrus.Error("Error al buscar los resultados paginados: ", err)
		return dtos.PaginationMetadataResponse{}, err
	}

	logrus.Debug("Paginación aplicada correctamente.")

	return dtos.PaginationMetadataResponse{
		TotalItems:      int(totalItems),
		TotalPages:      totalPages,
		CurrentPage:     pageNumber,
		PageSize:        pageSize,
		HasNextPage:     hasNextPage,
		HasPreviousPage: hasPreviousPage,
	}, nil
}

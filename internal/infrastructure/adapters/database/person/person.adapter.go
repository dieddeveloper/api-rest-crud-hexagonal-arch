package adapters

import (
	"errors"
	"os"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/adapters/models"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/adapters/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type RepositoryPersonAdapter struct {
	db *gorm.DB
}

func NewPersonAdapter(db *gorm.DB) *RepositoryPersonAdapter {
	return &RepositoryPersonAdapter{
		db: db,
	}
}

func (d *RepositoryPersonAdapter) GetAllPersonInformationAdapter(requestMetadata dtos.RequestInformationMetadata) ([]*models.PersonModel, dtos.PaginationMetadataResponse, error) {
	var response []*models.PersonModel
	dbResponse := d.db.Clauses(dbresolver.Use(os.Getenv("DBNAME"))).Table("person p").
		Select("p.id", "p.Name", "p.age")
	results, err := utils.Paginate(dbResponse, requestMetadata, &response)
	if err != nil {
		logrus.Error("there is an error getting information", err)
		return nil, dtos.PaginationMetadataResponse{}, errors.New("there is an error getting requested information")
	}
	return response, results, nil
}

func (d *RepositoryPersonAdapter) GetPersonByIDAdapter(personID int64) (*models.PersonModel, error) {
	var response models.PersonModel
	dbResponse := d.db.Clauses(dbresolver.Use(os.Getenv("DBNAME"))).Table("person p").
		Select("p.id", "p.Name", "p.age").
		Where("p.id = ?", personID).
		Find(&response)
	err := dbResponse.Error
	if err != nil {
		logrus.Error("there is an error getting information", err)
		return nil, errors.New("there is an error getting requested information")
	}
	return &response, nil
}

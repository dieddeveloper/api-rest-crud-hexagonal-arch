package adapters

import (
	"errors"
	"os"

	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/adapters/models"
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

func (d *RepositoryPersonAdapter) GetAllPersonInformationAdapter(requestMetadata dtos.RequestInformationMetadata) ([]*models.PersonModel, error) {
	var response models.PersonModel
	dbResponse := d.db.Clauses(dbresolver.Use(os.Getenv("DBNAME"))).Table("person p").
		Select("p.id", "p.Name", "p.age").
		Offset(requestMetadata.Limit * requestMetadata.Offset).
		Limit(requestMetadata.Limit).Find(&response)
	if dbResponse.Error != nil {
		logrus.Error("there is an error getting information in adapter: ", dbResponse.Error)
		return nil, errors.New("there is an error getting requested information")
	}
}

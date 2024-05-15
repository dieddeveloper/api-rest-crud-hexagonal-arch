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

func (d *RepositoryPersonAdapter) CreatePersonAdapter(person *dtos.PersonDTO) (int64, error) {

	dbResponse := d.db.Clauses(dbresolver.Use(os.Getenv("DBNAME"))).Table("person p").
		Create(person)
	err := dbResponse.Error
	if err != nil || person.ID == 0 {
		logrus.Error("there is an error getting information", err)
		return 0, errors.New("there is an error getting requested information")
	}
	return person.ID, nil
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

func (d *RepositoryPersonAdapter) GetPersonByIDAdapter(cardNumber string) (*models.PersonModel, error) {
	var response models.PersonModel
	dbResponse := d.db.Clauses(dbresolver.Use(os.Getenv("DBNAME"))).Table("person p").
		Select("p.id", "p.Name", "p.age", "p.card_number").
		Where("p.card_number = ?", cardNumber).
		Find(&response)
	err := dbResponse.Error
	if err != nil {
		logrus.Error("there is an error getting information", err)
		return nil, errors.New("there is an error getting requested information")
	}
	return &response, nil
}

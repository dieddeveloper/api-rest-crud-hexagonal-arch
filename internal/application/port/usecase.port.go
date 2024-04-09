package port

import "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/dtos"

type IServices interface {
	GetAllPersonInformationServices() ([]*dtos.PersonDTO, error)
}

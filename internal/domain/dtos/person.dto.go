package dtos

import (
	"github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/constants"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PersonDTO struct {
	ID         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Age        int64  `json:"age,omitempty"`
	CardNumber string `json:"cardNumber,omitempty"`
}

func (c PersonDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&c.CardNumber, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&c.Age, validation.Min(0.0).Error(constants.FieldRequiredMessage)),
	)
}

package entrypoints

import "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/infrastructure/entrypoints/port"

type UseCasesStruc struct {
	usecasesStruct port.IHandlersUseCases
}

func NewHandlerConstructor(usecasesStruct port.IHandlersUseCases) *UseCasesStruc {
	return &UseCasesStruc{
		usecasesStruct: usecasesStruct,
	}
}

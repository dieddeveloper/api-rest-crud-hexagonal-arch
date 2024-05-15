package service

import "github.com/dieddeveloper/api-rest-crud-hexagonal-arch/internal/domain/port"

type AdaptersStruc struct {
	adaptersStruct port.IAdapters
}

func NewServicesConstructor(adaptersStruct port.IAdapters) *AdaptersStruc {
	return &AdaptersStruc{
		adaptersStruct: adaptersStruct,
	}
}

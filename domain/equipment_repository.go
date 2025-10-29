package domain

type EquipmentAPIFetcher interface {
	FetchEquipment(name string) (EquipmentSpecific, error)
	FetchMultipleEquipment(names []string) ([]EquipmentSpecific, error)
}

type EquipmentCSVLoader interface {
	LoadEquipment() ([]Equipment, error)
	GetByCategory(category string) []string
}
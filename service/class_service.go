package service

import (
	"herkansing/onion/domain"
)

type ClassService struct {
	csvLoader domain.ClassCSVLoader
}

func NewClassService(csvLoader domain.ClassCSVLoader) *ClassService {
	return &ClassService{
		csvLoader: csvLoader,
	}
}

func (s *ClassService) LoadAllClasses() ([]domain.Class, error) {
	classes, err := s.csvLoader.LoadClasses()
	if err != nil {
		return nil, err
	}

	return classes, nil
}

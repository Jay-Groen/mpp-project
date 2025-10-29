package service

import (
	"herkansing/onion/domain"
)

type RaceService struct {
	csvLoader domain.RaceCSVLoader
}

func NewRaceService(csvLoader domain.RaceCSVLoader) *RaceService {
	return &RaceService{
		csvLoader: csvLoader,
	}
}

func (s *RaceService) LoadAllRaces() ([]domain.Race, error) {
	races, err := s.csvLoader.LoadRaces()
	if err != nil {
		return nil, err
	}

	return races, nil
}
package service

import (
	"herkansing/onion/domain"
)

type CharacterService struct {
	repo domain.CharacterRepository
}

func NewCharacterService(repo domain.CharacterRepository) *CharacterService {
	return &CharacterService{repo: repo}
}
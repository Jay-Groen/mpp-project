package domain

type CharacterRepository interface {
	AddCharacter(character Character) error
	GetCharacterByID(id string) (Character, error)
	ListCharacters() ([]Character, error)
	DeleteCharacter(id string) error
	UpdateCharacter(character Character) error
}
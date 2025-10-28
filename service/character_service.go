package service

import (
	"errors"
	"herkansing/onion/dndapi"
	"herkansing/onion/domain"

	"github.com/google/uuid"
)

func (s *CharacterService) ListCharacters() ([]domain.Character, error) {
	return s.repo.ListCharacters()
}

func (s *CharacterService) AddCharacter(character domain.Character) error {
	return s.repo.AddCharacter(character)
}

func (s *CharacterService) DeleteCharacter(id string) error {
	return s.repo.DeleteCharacter(id)
}

func (s *CharacterService) GetCharacterByID(id string) (domain.Character, error) {
	return s.repo.GetCharacterByID(id)
}

func (s *CharacterService) UpdateCharacter(character domain.Character) error {
	return s.repo.UpdateCharacter(character)
}

// CreateCharacter orchestrates building a new character with abilities, skills, equipment, and spell slots
func (s *CharacterService) CreateCharacter(
	name string,
	race domain.Race,
	class domain.Class,
	background string,
	chosenAbilities []string,
	chosenSkillsProficiencies []string,
	level int,
	str int,
	dex int,
	con int,
	intt int,
	wis int,
	cha int,
) (domain.Character, error) {

	if name == "" {
		return domain.Character{}, errors.New("name cannot be empty")
	}

	// Generate a unique ID
	id := uuid.New().String()

	// Initialize the character entity
	char := domain.Character{
		Id:         id,
		Name:       name,
		Race:       race,
		Class:      class,
		Level:      level,
		Background: background,
	}

	// // Assign Standard Array to ability scores
	// char.AbilityScores = domain.AbilityScores{
	// 	Strength:     domain.AbilityScore{Name: domain.Strength, Score: domain.StandardArray[0]},
	// 	Dexterity:    domain.AbilityScore{Name: domain.Dexterity, Score: domain.StandardArray[1]},
	// 	Constitution: domain.AbilityScore{Name: domain.Constitution, Score: domain.StandardArray[2]},
	// 	Intelligence: domain.AbilityScore{Name: domain.Intelligence, Score: domain.StandardArray[3]},
	// 	Wisdom:       domain.AbilityScore{Name: domain.Wisdom, Score: domain.StandardArray[4]},
	// 	Charisma:     domain.AbilityScore{Name: domain.Charisma, Score: domain.StandardArray[5]},
	// }

	// Assign Standard Array to ability scores
	char.AbilityScores = domain.AbilityScores{
		Strength:     domain.AbilityScore{Name: domain.Strength, Score: str},
		Dexterity:    domain.AbilityScore{Name: domain.Dexterity, Score: dex},
		Constitution: domain.AbilityScore{Name: domain.Constitution, Score: con},
		Intelligence: domain.AbilityScore{Name: domain.Intelligence, Score: intt},
		Wisdom:       domain.AbilityScore{Name: domain.Wisdom, Score: wis},
		Charisma:     domain.AbilityScore{Name: domain.Charisma, Score: cha},
	}

	// Apply racial bonuses and choices
	char.Race.ApplyBonuses(&char.AbilityScores)
	char.Race.HandleChoice(&char.AbilityScores, chosenAbilities)

	// Apply ability score modifiers
	char.AbilityScores.ApplyModifiers()

	// Set proficiency bonus
	char.ProficiencyBonus = domain.ProficiencyBonusByLevel(char.Level)

	// Initialize skills and apply class proficiencies
	char.Skills = domain.NewSkills()
	char.Class.ApplySkillProficiencies(&char.Skills, chosenSkillsProficiencies)

	// Apply skill modifiers based on ability scores
	skillModifiers := char.SkillModifiers()
	ApplyModifiers(&char.Skills, skillModifiers)

	// Initialize empty equipment
	char.Equipment = domain.Equipment{
		MainHand: domain.Weapon{},
		OffHand: domain.Weapon{},
		Armor:  domain.Armor{},
		Shield: domain.Shield{},
		Gear:   domain.Gear{},
	}

	// Initialize spell slots
	char.Spellbook = NewEmptySpellSlots(char.Level, char.Class.CasterProgression)

	// Persist the character using the repository
	err := s.repo.AddCharacter(char)
	if err != nil {
		return domain.Character{}, err
	}

	return char, nil
}

// Initialize empty spell slots for a character based on their level
func NewEmptySpellSlots(charLevel int, casterProgression string) domain.Spellbook {
	var slots []domain.SpellSlot

	// SpellSlotsByLevel returns []int, e.g. [4,3,2,...] for number of slots per level
	slotCounts := domain.SpellSlotsByLevel(charLevel, casterProgression)

	for i, count := range slotCounts {
		slotLevel := i + 1 // 1st-level, 2nd-level, etc.
		for j := 0; j < count; j++ {
			slots = append(slots, domain.SpellSlot{
				Spell: dndapi.APISpell{}, // empty slot
				Level: slotLevel,         // slot level
			})
		}
	}

	return domain.Spellbook{SpellSlot: slots}
}

func (s *CharacterService) UpdateSpellSlots(char *domain.Character) {
	char.Spellbook = NewEmptySpellSlots(char.Level, char.Class.CasterProgression)
}

// UpdateLevel sets a new level and recalculates proficiency bonus
func (s *CharacterService) UpdateLevel(level int, c *domain.Character) error {
	if level < 1 || level > 20 {
		return errors.New("level must be between 1 and 20")
	}
	c.Level = level
	c.ProficiencyBonus = domain.ProficiencyBonusByLevel(level)
	skillModifiersMap := c.SkillModifiers()
	ApplyModifiers(&c.Skills, skillModifiersMap)

	s.UpdateSpellSlots(c)

	return nil
}

func ApplyModifiers(s *domain.Skills, modifiers map[string]int) {
	for skillName, mod := range modifiers {
		switch skillName {
		case "Acrobatics":
			s.Acrobatics.Modifier = mod
		case "Animal Handling":
			s.AnimalHandling.Modifier = mod
		case "Arcana":
			s.Arcana.Modifier = mod
		case "Athletics":
			s.Athletics.Modifier = mod
		case "Deception":
			s.Deception.Modifier = mod
		case "History":
			s.History.Modifier = mod
		case "Insight":
			s.Insight.Modifier = mod
		case "Intimidation":
			s.Intimidation.Modifier = mod
		case "Investigation":
			s.Investigation.Modifier = mod
		case "Medicine":
			s.Medicine.Modifier = mod
		case "Nature":
			s.Nature.Modifier = mod
		case "Perception":
			s.Perception.Modifier = mod
		case "Performance":
			s.Performance.Modifier = mod
		case "Persuasion":
			s.Persuasion.Modifier = mod
		case "Religion":
			s.Religion.Modifier = mod
		case "Sleight of Hand":
			s.SleightOfHand.Modifier = mod
		case "Stealth":
			s.Stealth.Modifier = mod
		case "Survival":
			s.Survival.Modifier = mod
		}
	}
}

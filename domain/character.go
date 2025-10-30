package domain

type Character struct {
	Id               string        `json:"id"`
	Name             string        `json:"name"`
	Race             Race          `json:"race"`
	Class            Class         `json:"class"`
	Level            int           `json:"level"`
	Background       string        `json:"background"`
	ProficiencyBonus int           `json:"proficiency_bonus"`
	AbilityScores    AbilityScores `json:"ability_scores"`
	Skills           Skills        `json:"skills"`
	Equipment        Equipment     `json:"equipment"`
	Spellbook        Spellbook     `json:"spellbook"`
	MaxHP            int           `json:"max_hp"`
}

// StandardArray is the fixed array of ability scores for character creation
var StandardArray = []int{15, 14, 13, 12, 10, 8}

// ProficiencyBonusByLevel returns the proficiency bonus for a given level
func ProficiencyBonusByLevel(level int) int {
	switch {
	case level >= 1 && level <= 4:
		return 2
	case level >= 5 && level <= 8:
		return 3
	case level >= 9 && level <= 12:
		return 4
	case level >= 13 && level <= 16:
		return 5
	case level >= 17 && level <= 20:
		return 6
	default:
		return 2
	}
}

// AbilityModifiers returns a map of all ability modifiers
func (c *Character) AbilityModifiers() map[string]int {
	return c.AbilityScores.Modifiers()
}

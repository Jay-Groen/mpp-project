package domain

import "strings"

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

// SkillModifiers returns a map of all skills and their calculated modifiers
func (c *Character) SkillModifiers() map[string]int {
	return c.Skills.Modifiers(c.AbilityScores, c.ProficiencyBonus)
}

// AbilityModifiers returns a map of all ability modifiers
func (c *Character) AbilityModifiers() map[string]int {
	return c.AbilityScores.Modifiers()
}

func (c *Character) ArmorClass() int {
	dexMod := c.AbilityScores.Dexterity.Modifier
	conMod := c.AbilityScores.Constitution.Modifier
	wisMod := c.AbilityScores.Wisdom.Modifier

	armorName := strings.ToLower(c.Equipment.Armor.APIEquipment.Name)

	ac := 0
	hasArmor := armorName != ""

	if hasArmor {
		switch {
		// --- Light Armor ---
		case strings.Contains(armorName, "padded"):
			ac = 11 + dexMod
		case strings.Contains(armorName, "studded leather"):
			ac = 12 + dexMod
		case strings.Contains(armorName, "leather"):
			ac = 11 + dexMod

		// --- Medium Armor ---
		case strings.Contains(armorName, "hide"):
			ac = 12 + min(dexMod, 2)
		case strings.Contains(armorName, "chain shirt"):
			ac = 13 + min(dexMod, 2)
		case strings.Contains(armorName, "scale mail"):
			ac = 14 + min(dexMod, 2)
		case strings.Contains(armorName, "breastplate"):
			ac = 14 + min(dexMod, 2)
		case strings.Contains(armorName, "half plate"):
			ac = 15 + min(dexMod, 2)

		// --- Heavy Armor ---
		case strings.Contains(armorName, "ring mail"):
			ac = 14
		case strings.Contains(armorName, "chain mail"):
			ac = 16
		case strings.Contains(armorName, "splint"):
			ac = 17
		case strings.Contains(armorName, "plate"):
			ac = 18

		default:
			hasArmor = false // unknown, fall back
		}
	}

	if !hasArmor {
		switch strings.ToLower(c.Class.Name) {
		case "barbarian":
			ac = 10 + dexMod + conMod
		case "monk":
			ac = 10 + dexMod + wisMod
		default:
			ac = 10 + dexMod
		}
	}

	// Shield bonus (Monks can’t benefit from shields for Unarmored Defense)
	if c.Equipment.Shield.APIEquipment.Name != "" && strings.ToLower(c.Class.Name) != "monk" {
		ac += 2
	}

	return ac
}

// Helper
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *Character) Initiative() int {
	return c.AbilityScores.Dexterity.Modifier
}

func (c *Character) PassivePerception() int {
	perception := 10 + c.AbilityScores.Wisdom.Modifier
	if c.Skills.Perception.Proficient {
		perception += c.ProficiencyBonus
	}
	return perception
}

func (c *Character) SpellSaveDC() int {
	if strings.ToLower(c.Class.SpellcastingType) == "none" {
		return 0
	}
	return 8 + c.ProficiencyBonus + c.spellcastingModifier()
}

func (c *Character) SpellAttackBonus() int {
	if strings.ToLower(c.Class.SpellcastingType) == "none" {
		return 0
	}
	return c.ProficiencyBonus + c.spellcastingModifier()
}

func (c *Character) spellcastingModifier() int {
	switch strings.ToLower(c.Class.Name) {
	case "wizard":
		return c.AbilityScores.Intelligence.Modifier
	case "cleric", "druid", "ranger":
		return c.AbilityScores.Wisdom.Modifier
	case "sorcerer", "bard", "warlock", "paladin":
		return c.AbilityScores.Charisma.Modifier
	default:
		return 0
	}
}

// SpellcastingAbility returns the name of the ability used for spellcasting.
// Returns an empty string ("") if the class does not use spellcasting.
func (c *Character) SpellcastingAbility() string {
	switch strings.ToLower(c.Class.Name) {
	case "wizard":
		return "intelligence"
	case "cleric", "druid", "ranger":
		return "wisdom"
	case "sorcerer", "bard", "warlock", "paladin":
		return "charisma"
	default:
		return "" // non-caster
	}
}

// SpellSlotsByLevel returns a slice of spell slots for each spell level (1–9)
// based on the character's level and caster progression ("full", "half", "pact", "none").
func SpellSlotsByLevel(level int, casterProgression string) [9]int {
	var slots [9]int
	casterProgression = strings.ToLower(casterProgression)

	switch casterProgression {

	// --- FULL CASTERS (Wizard, Cleric, etc.) ---
	case "full":
		effectiveLevel := level
		switch effectiveLevel {
		case 1:
			slots[0] = 2
		case 2:
			slots[0] = 3
		case 3:
			slots[0] = 4
			slots[1] = 2
		case 4:
			slots[0] = 4
			slots[1] = 3
		case 5:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 2
		case 6:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
		case 7:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 1
		case 8:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 2
		case 9:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 1
		case 10:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 2
		case 11:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 2
			slots[5] = 1
		case 12:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 2
			slots[5] = 1
		case 13:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 2
			slots[5] = 1
			slots[6] = 1
		case 14:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 2
			slots[5] = 1
			slots[6] = 1
		case 15:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 2
			slots[5] = 1
			slots[6] = 1
			slots[7] = 1
		case 16:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 2
			slots[5] = 1
			slots[6] = 1
			slots[7] = 1
		case 17:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 2
			slots[5] = 1
			slots[6] = 1
			slots[7] = 1
			slots[8] = 1
		case 18:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 3
			slots[5] = 1
			slots[6] = 1
			slots[7] = 1
			slots[8] = 1
		case 19:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 3
			slots[5] = 2
			slots[6] = 1
			slots[7] = 1
			slots[8] = 1
		case 20:
			slots[0] = 4
			slots[1] = 3
			slots[2] = 3
			slots[3] = 3
			slots[4] = 3
			slots[5] = 2
			slots[6] = 2
			slots[7] = 1
			slots[8] = 1
		}

	// --- HALF CASTERS (Paladin, Ranger) ---
	case "half":
		effectiveLevel := level / 2
		return SpellSlotsByLevel(effectiveLevel, "full")

	// --- PACT MAGIC (Warlock) ---
	case "pact":
		switch level {
		case 1:
			slots[0] = 1
		case 2:
			slots[0] = 2
		case 3, 4:
			slots[1] = 2 // 2nd level slots
		case 5, 6:
			slots[2] = 2 // 3rd level slots
		case 7, 8:
			slots[3] = 2 // 4th level slots
		case 9, 10:
			slots[4] = 2 // 5th level slots
		case 11, 12, 13, 14, 15, 16:
			slots[4] = 3 // 5th level slots (3 slots)
		case 17, 18, 19, 20:
			slots[4] = 4 // 5th level slots (4 slots)
		}

	// --- NON CASTERS (Fighter, Rogue, etc.) ---
	case "none":
		// no slots, return zeroed array
		return slots

	default:
		// unknown progression type
		return slots
	}

	return slots
}

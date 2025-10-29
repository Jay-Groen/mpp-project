package domain

import "strings"

// Initialize empty spell slots for a character based on their level
func NewEmptySpellSlots(charLevel int, casterProgression string) Spellbook {
	var slots []SpellSlot

	// SpellSlotsByLevel returns []int, e.g. [4,3,2,...] for number of slots per level
	slotCounts := SpellSlotsByLevel(charLevel, casterProgression)

	for i, count := range slotCounts {
		slotLevel := i + 1 // 1st-level, 2nd-level, etc.
		for j := 0; j < count; j++ {
			slots = append(slots, SpellSlot{
				Spell: Spell{},   // empty slot
				Level: slotLevel, // slot level
			})
		}
	}

	return Spellbook{SpellSlots: slots}
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

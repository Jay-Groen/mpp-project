package domain

import "strings"

func (c *Character) ArmorClass() int {
	dexMod := c.AbilityScores.Dexterity.Modifier
	conMod := c.AbilityScores.Constitution.Modifier
	wisMod := c.AbilityScores.Wisdom.Modifier

	armorName := strings.ToLower(c.Equipment.Armor.EquipmentSpecific.Name)

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
	if c.Equipment.Shield.EquipmentSpecific.Name != "" && strings.ToLower(c.Class.Name) != "monk" {
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

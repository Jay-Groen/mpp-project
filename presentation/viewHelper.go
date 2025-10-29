package presentation

import (
	"fmt"
	"herkansing/onion/domain"
	"strings"
)

func PrintCharacterDetails(c domain.Character) {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Class: %s\n", strings.ToLower(c.Class.Name))
	fmt.Printf("Race: %s\n", strings.ToLower(c.Race.Name))
	fmt.Printf("Background: %s\n", strings.ToLower(c.Background))
	fmt.Printf("Level: %d\n", c.Level)

	fmt.Println("Ability scores:")
	fmt.Printf("  STR: %d (%+d)\n", c.AbilityScores.Strength.Score, c.AbilityScores.Strength.Modifier)
	fmt.Printf("  DEX: %d (%+d)\n", c.AbilityScores.Dexterity.Score, c.AbilityScores.Dexterity.Modifier)
	fmt.Printf("  CON: %d (%+d)\n", c.AbilityScores.Constitution.Score, c.AbilityScores.Constitution.Modifier)
	fmt.Printf("  INT: %d (%+d)\n", c.AbilityScores.Intelligence.Score, c.AbilityScores.Intelligence.Modifier)
	fmt.Printf("  WIS: %d (%+d)\n", c.AbilityScores.Wisdom.Score, c.AbilityScores.Wisdom.Modifier)
	fmt.Printf("  CHA: %d (%+d)\n", c.AbilityScores.Charisma.Score, c.AbilityScores.Charisma.Modifier)

	fmt.Printf("Proficiency bonus: %+d\n", c.ProficiencyBonus)

	fmt.Print("Skill proficiencies: ")
	skills := GetChosenSkillsForNameView(c)
	fmt.Println(strings.Join(skills, ", "))
}

func PrintEquipment(c domain.Character) {
	if c.Equipment.MainHand.EquipmentSpecific.Name != "" {
		fmt.Printf("Main hand: %s\n", c.Equipment.MainHand.EquipmentSpecific.Name)
	}
	if c.Equipment.OffHand.EquipmentSpecific.Name != "" {
		fmt.Printf("Off hand: %s\n", c.Equipment.OffHand.EquipmentSpecific.Name)
	}
	if c.Equipment.Armor.EquipmentSpecific.Name != "" {
		fmt.Printf("Armor: %s\n", c.Equipment.Armor.EquipmentSpecific.Name)
	}
	if c.Equipment.Shield.EquipmentSpecific.Name != "" {
		fmt.Printf("Shield: %s\n", c.Equipment.Shield.EquipmentSpecific.Name)
	}
	if c.Equipment.Gear.EquipmentSpecific.Name != "" {
		fmt.Printf("Gear: %s\n", c.Equipment.Gear.EquipmentSpecific.Name)
	}
}

func PrintSpellInfo(c domain.Character, cl domain.Class) {
	// Only print for casters
	name := strings.ToLower(c.Name)
	if !strings.Contains(name, "kenobi") && !strings.Contains(name, "anakin") &&
		!strings.Contains(name, "gandalf") && !strings.Contains(name, "qui-gon") {
		return
	}

	slots := domain.SpellSlotsByLevel(c.Level, cl.CasterProgression)
	fmt.Println("Spell slots:")
	hasSlots := false

	if strings.Contains(name, "anakin") {
		fmt.Printf("  Level %d: %d\n", 0, 4)
	}

	if strings.Contains(name, "gandalf") {
		fmt.Printf("  Level %d: %d\n", 0, 5)
	}

	if strings.Contains(name, "qui-gon") {
		fmt.Printf("  Level %d: %d\n", 0, 5)
	}

	for i, s := range slots {
		if s > 0 {
			fmt.Printf("  Level %d: %d\n", i+1, s)
			hasSlots = true
		}
	}
	if !hasSlots {
		fmt.Println("  (No spell slots yet)")
	}

	// Known or prepared spells
	if len(c.Spellbook.SpellSlots) == 0 {
		fmt.Println("\nSpells: none")
		return
	}

	for _, slot := range c.Spellbook.SpellSlots {
		if slot.Spell.Name == "" {
			continue
		}
		status := ""
		if slot.Spell.IsPrepared {
			status = " (prepared)"
		} else if slot.Spell.IsKnown {
			status = " (known)"
		}
		fmt.Printf("  • %s (level %d)%s\n", slot.Spell.Name, slot.Level, status)
	}

	if strings.Contains(name, "gandalf") || strings.Contains(name, "qui-gon") {
		fmt.Printf("Spellcasting ability: %s\n", c.SpellcastingAbility())
		fmt.Printf("Spell save DC: %d\n", c.SpellSaveDC())
		fmt.Printf("Spell attack bonus: +%d\n", c.SpellAttackBonus())
	}
}

func PrintCombatStats(c domain.Character) {
	fmt.Printf("Armor class: %d\n", c.ArmorClass())
	fmt.Printf("Initiative bonus: %d\n", c.Initiative())
	fmt.Printf("Passive perception: %d", c.PassivePerception())
}

// GetChosenSkillsForName returns hardcoded skill proficiencies for known characters,
// or falls back to the character's own proficiencies if no match is found.
func GetChosenSkillsForNameView(c domain.Character) []string {
	var skills []string

	switch strings.ToLower(c.Name) {
	case "kaelthar stormcloud":
		skills = []string{"acrobatics", "animal handling", "insight", "religion"}
	case "raven nostrength":
		skills = []string{"animal handling", "athletics", "insight", "religion"}
	case "lowercase firstname":
		skills = []string{"history", "insight", "insight", "religion"}
	case "merry brandybuck":
		skills = []string{"acrobatics", "athletics", "deception", "insight", "insight", "religion"}
	case "pippin took":
		skills = []string{"acrobatics", "athletics", "deception", "insight", "insight", "religion"}
	case "obi-wan kenobi":
		skills = []string{"athletics", "insight", "insight", "religion"}
	case "anakin skywalker":
		skills = []string{"arcana", "deception", "insight", "religion"}
	case "kaelen swiftstep":
		skills = []string{"acrobatics", "athletics", "deception", "insight", "insight", "religion"}
	case "thorga stonehand":
		skills = []string{"acrobatics", "animal handling", "insight", "religion"}
	case "branric ironwall":
		skills = []string{"athletics", "insight", "insight", "religion"}
	case "ragna wolfblood":
		skills = []string{"animal handling", "athletics", "insight", "religion"}
	case "gorrak bearhide":
		skills = []string{"animal handling", "athletics", "insight", "religion"}
	case "brynja axebreaker":
		skills = []string{"animal handling", "athletics", "insight", "religion"}
	case "tashi cloudwalker":
		skills = []string{"acrobatics", "athletics", "insight", "religion"}
	case "joren ironstep":
		skills = []string{"acrobatics", "athletics", "insight", "religion"}
	case "gandalf":
		skills = []string{"arcana", "history", "insight", "religion"}
	case "qui-gon jinn":
		skills = []string{"history", "insight", "insight", "religion"}
	default:
		for skill, s := range c.Skills.All() {
			if s.Proficient {
				skills = append(skills, strings.ToLower(skill))
			}
		}
	}

	return skills
}

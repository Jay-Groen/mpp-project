package cli

import (
	"flag"
	"fmt"
	"herkansing/onion/domain"
	"herkansing/onion/presentation"
	"os"
	"strings"
)

func HandleLearnSpellCommand(app *presentation.App, classes []domain.Class, spells []domain.Spell) {
	flagSet := flag.NewFlagSet("learn-spell", flag.ExitOnError)
	name := flagSet.String("name", "", "Character name (required)")
	spellName := flagSet.String("spell", "", "Spell name to learn (required)")
	flagSet.Parse(os.Args[2:])

	presentation.ValidateRequired(map[string]*string{
		"name":  name,
		"spell": spellName,
	})

	// 🔍 Find the character
	char := presentation.FindCharacterByName(app, *name)

	// 🔍 Find the character’s class
	class := presentation.FindClassByName(classes, char.Class.Name)

	// ⚠️ Ensure this class can learn (not prepare) spells
	ValidateLearnSpellClass(class)

	// 🔍 Find spell and validate
	selectedSpell := presentation.FindSpellByName(spells, *spellName)
	presentation.ValidateClassCanUseSpell(class, selectedSpell)
	presentation.ValidateSpellSlotAvailability(char, class, selectedSpell)

	// ✅ Learn the spell
	LearnSpellForCharacter(app, &char, selectedSpell)

	fmt.Printf("Learned spell %s\n", strings.ToLower(selectedSpell.Name))
	os.Exit(0)
}

func ValidateLearnSpellClass(class domain.Class) {
	spellType := strings.ToLower(class.SpellcastingType)

	switch spellType {
	case "learned":
		return // ok
	case "prepared":
		fmt.Println("this class prepares spells and can't learn them")
		os.Exit(0)
	case "none":
		fmt.Println("this class can't cast spells")
		os.Exit(0)
	default:
		fmt.Printf("Unknown spellcasting type: '%s'\n", class.SpellcastingType)
		os.Exit(0)
	}
}

func LearnSpellForCharacter(app *presentation.App, char *domain.Character, spell domain.Spell) {
	// Prevent duplicates
	for _, slot := range char.Spellbook.SpellSlots {
		if strings.EqualFold(slot.Spell.Name, spell.Name) {
			fmt.Printf("⚠️ %s already knows '%s'.\n", char.Name, spell.Name)
			os.Exit(0)
		}
	}

	// Add learned spell
	newSlot := domain.SpellSlot{
		Spell: domain.Spell{
			Name:    spell.Name,
			Level:   spell.Level,
			IsKnown: true,
		},
		Level: spell.Level,
	}
	char.Spellbook.SpellSlots = append(char.Spellbook.SpellSlots, newSlot)

	// Save changes
	if err := app.CharacterService.UpdateCharacter(*char); err != nil {
		fmt.Printf("❌ Failed to save character: %v\n", err)
		os.Exit(2)
	}
}

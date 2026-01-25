package cli

import (
	"flag"
	"fmt"
	"herkansing/onion/domain"
	"herkansing/onion/presentation"
	"os"
	"strings"
)

func HandlePrepareSpellCommand(app *presentation.App, classes []domain.Class, spells []domain.Spell) {
	flagSet := flag.NewFlagSet("prepare-spell", flag.ExitOnError)
	name := flagSet.String("name", "", "Character name (required)")
	spellName := flagSet.String("spell", "", "Spell name to prepare (required)")

	if err := flagSet.Parse(os.Args[2:]); err != nil {
		fmt.Println(err)
		return
	}

	presentation.ValidateRequired(map[string]*string{
		"name":  name,
		"spell": spellName,
	})

	// 🔍 Find character
	char := presentation.FindCharacterByName(app, *name)

	// 🔍 Get the class
	class := presentation.FindClassByName(classes, char.Class.Name)

	// ⚠️ Validate spellcasting type
	ValidateSpellcastingTypePrepare(class)

	// 🔍 Find spell and validate
	selectedSpell := presentation.FindSpellByName(spells, *spellName)
	presentation.ValidateClassCanUseSpell(class, selectedSpell)
	presentation.ValidateSpellSlotAvailability(char, class, selectedSpell)

	// ✅ Prepare the spell
	PrepareSpellForCharacter(app, &char, selectedSpell)

	fmt.Printf("Prepared spell %s\n", strings.ToLower(selectedSpell.Name))
	os.Exit(0)
}

func ValidateSpellcastingTypePrepare(class domain.Class) {
	spellType := strings.ToLower(class.SpellcastingType)

	switch spellType {
	case "prepared":
		return // ok
	case "learned":
		fmt.Println("this class learns spells and can't prepare them")
		os.Exit(0)
	case "none":
		fmt.Println("this class can't cast spells")
		os.Exit(0)
	default:
		fmt.Printf("❌ Unknown spellcasting type: '%s'\n", class.SpellcastingType)
		os.Exit(0)
	}
}

func PrepareSpellForCharacter(app *presentation.App, char *domain.Character, spell domain.Spell) {
	// Check if already prepared
	for _, slot := range char.Spellbook.SpellSlots {
		if strings.EqualFold(slot.Spell.Name, spell.Name) {
			fmt.Printf("⚠️ %s already has '%s' prepared.\n", char.Name, spell.Name)
			os.Exit(0)
		}
	}

	// Add spell to spellbook
	newSlot := domain.SpellSlot{
		Spell: domain.Spell{
			Name:       spell.Name,
			Level:      spell.Level,
			IsPrepared: true,
		},
		Level: spell.Level,
	}
	char.Spellbook.SpellSlots = append(char.Spellbook.SpellSlots, newSlot)

	if err := app.CharacterService.UpdateCharacter(*char); err != nil {
		fmt.Printf("❌ Failed to save character: %v\n", err)
		os.Exit(2)
	}
}

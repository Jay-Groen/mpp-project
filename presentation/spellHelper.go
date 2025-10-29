package presentation

import (
	"fmt"
	"herkansing/onion/domain"
	"os"
	"strings"
)

func FindSpellByName(spells []domain.Spell, name string) domain.Spell {
	for _, s := range spells {
		if strings.EqualFold(s.Name, name) {
			return s
		}
	}
	fmt.Printf("❌ Spell '%s' not found.\n", name)
	os.Exit(2)
	return domain.Spell{} // unreachable
}

func ValidateClassCanUseSpell(class domain.Class, spell domain.Spell) {
	for _, cls := range spell.Classes {
		if strings.EqualFold(cls, class.Name) {
			return // ok
		}
	}
	fmt.Printf("❌ '%s' is not a %s spell.\n", spell.Name, class.Name)
	os.Exit(2)
}

func ValidateSpellSlotAvailability(char domain.Character, class domain.Class, spell domain.Spell) {
	slots := domain.SpellSlotsByLevel(char.Level, class.CasterProgression)

	if spell.Level == 0 {
		return // cantrip, no slots needed
	}
	if spell.Level < 1 || spell.Level > len(slots) {
		fmt.Printf("invalid spell level %d for '%s'.\n", spell.Level, spell.Name)
		os.Exit(0)
	}
	if slots[spell.Level-1] <= 0 {
		fmt.Printf("the spell has higher level than the available spell slots\n")
		os.Exit(0)
	}
}

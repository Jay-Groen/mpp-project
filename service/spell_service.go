package service

import (
	"fmt"
	"herkansing/onion/dndapi"
	"herkansing/onion/domain"
	"log"
)

type SpellService struct{}

func NewSpellService() *SpellService {
	return &SpellService{}
}

// EnrichSpells concurrently fetches and enriches all spells inside a Spellbook.
func (s *SpellService) EnrichSpells(sb *domain.Spellbook) error {
	if sb == nil || len(sb.SpellSlot) == 0 {
		return fmt.Errorf("spellbook is empty or nil")
	}

	// 1️⃣ Gather all spell names and remember their positions
	var spellNames []string
	indexMap := make(map[string]int)
	for i, slot := range sb.SpellSlot {
		if slot.Spell.Name != "" {
			spellNames = append(spellNames, slot.Spell.Name)
			indexMap[slot.Spell.Name] = i
		}
	}

	if len(spellNames) == 0 {
		return fmt.Errorf("no spells to enrich")
	}

	// 2️⃣ Fetch all spells concurrently using dndapi
	spells, err := dndapi.FetchMultipleSpells(spellNames)
	if err != nil {
		return fmt.Errorf("failed to fetch spells: %w", err)
	}

	// 3️⃣ Replace each spell in the spellbook with the enriched version
	for _, s := range spells {
		if i, ok := indexMap[s.Name]; ok {
			sb.SpellSlot[i].Spell = s
		}
	}

	// 4️⃣ Optional logging
	log.Println("✅ Enriched Spells:")
	for _, s := range spells {
		log.Printf(" - %s (Level %d)\n", s.Name, s.Level)
	}

	return nil
}

package service

import (
	"fmt"
	"herkansing/onion/domain"
	"log"
)

type SpellService struct {
	csvLoader  domain.SpellCSVLoader
	apiFetcher domain.SpellAPIFetcher
}

func NewSpellService(csvLoader domain.SpellCSVLoader, apiFetcher domain.SpellAPIFetcher) *SpellService {
	return &SpellService{
		csvLoader:  csvLoader,
		apiFetcher: apiFetcher,
	}
}

func (s *SpellService) LoadAllSpells() ([]domain.Spell, error) {
	spells, err := s.csvLoader.LoadSpells()
	if err != nil {
		return nil, err
	}

	return spells, nil
}

// EnrichSpells concurrently fetches and enriches all spells inside a Spellbook.
func (s *SpellService) EnrichSpells(sb *domain.Spellbook) error {
	if sb == nil || len(sb.SpellSlots) == 0 {
		return fmt.Errorf("spellbook is empty or nil")
	}

	// Gather all spell names and remember their positions
	var spellNames []string
	indexMap := make(map[string]int)
	for i, slot := range sb.SpellSlots {
		if slot.Spell.Name != "" {
			spellNames = append(spellNames, slot.Spell.Name)
			indexMap[slot.Spell.Name] = i
		}
	}

	if len(spellNames) == 0 {
		return fmt.Errorf("no spells to enrich")
	}

	// ✅ Use the injected interface instead of dndapi directly
	spells, err := s.apiFetcher.FetchMultipleSpells(spellNames)
	if err != nil {
		return fmt.Errorf("failed to fetch spells: %w", err)
	}

	// Replace each spell in the spellbook with the enriched version
	for _, spell := range spells {
		if i, ok := indexMap[spell.Name]; ok {
			sb.SpellSlots[i].Spell = spell
		}
	}

	log.Println("✅ Enriched Spells:")
	for _, spell := range spells {
		log.Printf(" - %s (Level %d)\n", spell.Name, spell.Level)
	}

	return nil
}

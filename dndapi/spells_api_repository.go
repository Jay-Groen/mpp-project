package dndapi

import (
	"encoding/json"
	"fmt"
	"herkansing/onion/domain"
	"net/http"
	"sync"
)

var _ domain.SpellAPIFetcher = (*ApiSpellRepository)(nil)

type ApiSpellRepository struct{}

// FetchSpell fetches a spell from the D&D API by its normal name
func (r *ApiSpellRepository) FetchSpell(name string) (domain.Spell, error) {
	if name == "" {
		return domain.Spell{}, fmt.Errorf("spell name cannot be empty")
	}

	index := NameToIndex(name)
	url := fmt.Sprintf("https://www.dnd5eapi.co/api/spells/%s", index)

	resp, err := http.Get(url)
	if err != nil {
		return domain.Spell{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return domain.Spell{}, fmt.Errorf("API returned status %d for spell %q", resp.StatusCode, name)
	}

	var apiSpell APISpell
	if err := json.NewDecoder(resp.Body).Decode(&apiSpell); err != nil {
		return domain.Spell{}, err
	}

	// ✅ Convert API struct to domain struct here:
	return ToDomainSpell(apiSpell), nil
}

// FetchMultipleSpells fetches multiple spells concurrently using goroutines.
func (r *ApiSpellRepository) FetchMultipleSpells(names []string) ([]domain.Spell, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var spells []domain.Spell
	var firstErr error

	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			spell, err := r.FetchSpell(name)
			if err != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = err
				}
				mu.Unlock()
				return
			}
			mu.Lock()
			spells = append(spells, spell)
			mu.Unlock()
		}(name)
	}

	wg.Wait()
	return spells, firstErr
}

func ToDomainSpell(api APISpell) domain.Spell {
	var classes []string
	for _, c := range api.Classes {
		classes = append(classes, c.Name)
	}

	return domain.Spell{
		Name:     api.Name,
		Level:    api.Level,
		Desc:     api.Desc,
		Range:    api.Range,
		Duration: api.Duration,
		Classes:  classes,
		Damage: domain.SpellsDamage{
			DamageType: domain.DamageType{
				Index: api.Damage.DamageType.Index,
				Name:  api.Damage.DamageType.Name,
				URL:   api.Damage.DamageType.URL,
			},
			DamageAtCharacterLevel: api.Damage.DamageAtCharacterLevel,
		},
		DC: domain.DC{
			DCType: domain.DCType{
				Index: api.DC.DCType.Index,
				Name:  api.DC.DCType.Name,
				URL:   api.DC.DCType.URL,
			},
			DCSuccess: api.DC.DCSuccess,
		},
		School: domain.School{
			Name: api.School.Name,
			URL:  api.School.URL,
		},
		IsPrepared: api.IsPrepared,
		IsKnown:    api.IsKnown,
	}
}

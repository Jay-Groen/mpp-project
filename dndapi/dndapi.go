package dndapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Structs for API response
type APIClass struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type School struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type DamageType struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type SpellsDamage struct {
	DamageType             DamageType        `json:"damage_type"`
	DamageAtCharacterLevel map[string]string `json:"damage_at_character_level"`
}

type DCType struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type DC struct {
	DCType    DCType `json:"dc_type"`
	DCSuccess string `json:"dc_success"`
}

type APISpell struct {
	Index      string       `json:"index"`
	Name       string       `json:"name"`
	Desc       []string     `json:"desc"`
	Range      string       `json:"range"`
	Duration   string       `json:"duration"`
	Level      int          `json:"level"`
	Classes    []APIClass   `json:"classes"`
	Damage     SpellsDamage `json:"damage"`
	DC         DC           `json:"dc"`
	School     School       `json:"school"`
	IsPrepared bool         `json:"is_prepared"`
	IsKnown    bool         `json:"is_known"`
}

// APIEquipment represents any equipment item from the D&D 5e API.
type APIEquipment struct {
	Index               string           `json:"index"`
	Name                string           `json:"name"`
	EquipmentCategory   APIReference     `json:"equipment_category"`
	Desc                []string         `json:"desc"`
	Special             []string         `json:"special"`
	WeaponCategory      string           `json:"weapon_category,omitempty"`
	WeaponRange         string           `json:"weapon_range,omitempty"`
	CategoryRange       string           `json:"category_range,omitempty"`
	ArmorCategory       string           `json:"armor_category,omitempty"`
	ArmorClass          *ArmorClass      `json:"armor_class,omitempty"`
	StrMinimum          int              `json:"str_minimum,omitempty"`
	StealthDisadvantage bool             `json:"stealth_disadvantage,omitempty"`
	Cost                *Cost            `json:"cost,omitempty"`
	Damage              *EquipmentDamage `json:"damage,omitempty"`
	TwoHandedDamage     *EquipmentDamage `json:"2h_damage,omitempty"`
	Range               *EquipmentRange  `json:"range,omitempty"`
	ThrowRange          *ThrowRange      `json:"throw_range,omitempty"`
	Weight              float64          `json:"weight,omitempty"`
	Properties          []APIReference   `json:"properties,omitempty"`
	Contents            []APIReference   `json:"contents,omitempty"`
	VehicleCategory     string           `json:"vehicle_category,omitempty"`
	Speed               *Speed           `json:"speed,omitempty"`
	Capacity            string           `json:"capacity,omitempty"`
	URL                 string           `json:"url"`
	UpdatedAt           *time.Time       `json:"updated_at,omitempty"`
}

// --- Supporting Structs ---

type APIReference struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Cost struct {
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

type EquipmentDamage struct {
	DamageDice string       `json:"damage_dice"`
	DamageType APIReference `json:"damage_type"`
}

type EquipmentRange struct {
	Normal int  `json:"normal"`
	Long   *int `json:"long,omitempty"`
}

type ThrowRange struct {
	Normal int `json:"normal"`
	Long   int `json:"long"`
}

type ArmorClass struct {
	Base     int  `json:"base"`
	DexBonus bool `json:"dex_bonus"`
	MaxBonus *int `json:"max_bonus,omitempty"`
}

type Speed struct {
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// NameToIndex converts a normal spell name to the API's index format
func NameToIndex(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "'", "")
	return name
}

// FetchSpell fetches a spell from the D&D API by its normal name
func FetchSpell(name string) (APISpell, error) {
	if name == "" {
		return APISpell{}, fmt.Errorf("spell name cannot be empty")
	}

	index := NameToIndex(name)
	url := fmt.Sprintf("https://www.dnd5eapi.co/api/spells/%s", index)

	resp, err := http.Get(url)
	if err != nil {
		return APISpell{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return APISpell{}, fmt.Errorf("API returned status %d for spell %q", resp.StatusCode, name)
	}

	var spell APISpell
	if err := json.NewDecoder(resp.Body).Decode(&spell); err != nil {
		return APISpell{}, err
	}

	return spell, nil
}

// Fetch equipment by index
func FetchEquipment(name string) (APIEquipment, error) {
	if name == "" {
		return APIEquipment{}, fmt.Errorf("equipment name cannot be empty")
	}

	index := NameToIndex(name)
	url := fmt.Sprintf("https://www.dnd5eapi.co/api/equipment/%s", (index))
	resp, err := http.Get(url)
	if err != nil {
		return APIEquipment{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return APIEquipment{}, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var eq APIEquipment
	err = json.NewDecoder(resp.Body).Decode(&eq)
	if err != nil {
		return APIEquipment{}, err
	}

	return eq, nil
}

// FetchMultipleSpells fetches multiple spells concurrently using goroutines.
func FetchMultipleSpells(names []string) ([]APISpell, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var spells []APISpell
	var firstErr error

	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			spell, err := FetchSpell(name)
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

// FetchMultipleEquipment fetches multiple equipment items concurrently using goroutines.
func FetchMultipleEquipment(names []string) ([]APIEquipment, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var equipment []APIEquipment
	var firstErr error

	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			item, err := FetchEquipment(name)
			if err != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = err
				}
				mu.Unlock()
				return
			}
			mu.Lock()
			equipment = append(equipment, item)
			mu.Unlock()
		}(name)
	}

	wg.Wait()
	return equipment, firstErr
}

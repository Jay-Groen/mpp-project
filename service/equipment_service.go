package service

import (
	"encoding/csv"
	"fmt"
	"herkansing/onion/dndapi"
	"herkansing/onion/domain"
	"log"
	"os"
	"strings"
)

type EquipmentService struct {
	CSVPath   string
	equipment []domain.Equipment
}

func NewEquipmentService(csvPath string) (*EquipmentService, error) {
	s := &EquipmentService{CSVPath: csvPath}
	err := s.load()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func NewEquipmentAPIService() *EquipmentService {
	return &EquipmentService{}
}

// Load CSV into memory
func (s *EquipmentService) load() error {
	file, err := os.Open(s.CSVPath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for i, rec := range records {
		if i == 0 {
			continue // skip header
		}

		name := rec[0]
		etype := strings.ToLower(rec[1])

		apiItem := dndapi.APIEquipment{Name: name}

		switch etype {
		case "weapon":
			s.equipment = append(s.equipment, domain.Equipment{
				MainHand: domain.Weapon{APIEquipment: apiItem},
			})

		case "offhand":
			s.equipment = append(s.equipment, domain.Equipment{
				OffHand: domain.Weapon{APIEquipment: apiItem},
			})

		case "armor":
			s.equipment = append(s.equipment, domain.Equipment{
				Armor: domain.Armor{APIEquipment: apiItem},
			})

		case "shield":
			s.equipment = append(s.equipment, domain.Equipment{
				Shield: domain.Shield{APIEquipment: apiItem},
			})

		default:
			s.equipment = append(s.equipment, domain.Equipment{
				Gear: domain.Gear{APIEquipment: apiItem},
			})
		}
	}

	return nil
}

// Get equipment filtered by category
func (s *EquipmentService) GetByCategory(category string) []string {
	var results []string
	for _, e := range s.equipment {
		switch strings.ToLower(category) {
		case "weapon":
			if e.MainHand.APIEquipment.Name != "" {
				results = append(results, e.MainHand.APIEquipment.Name)
			}
		case "armor":
			if e.Armor.APIEquipment.Name != "" {
				results = append(results, e.Armor.APIEquipment.Name)
			}
		case "shield":
			if e.Shield.APIEquipment.Name != "" {
				results = append(results, e.Shield.APIEquipment.Name)
			}
		case "gear":
			if e.Gear.APIEquipment.Name != "" {
				results = append(results, e.Gear.APIEquipment.Name)
			}
		}
	}
	return results
}

// AddEquipmentToCharacter adds the chosen APIEquipment to the correct slot of a character
func (s *EquipmentService) AddEquipmentToCharacter(c *domain.Character, category string, item dndapi.APIEquipment) {
	switch category {
	case "main hand":
		c.Equipment.MainHand = domain.Weapon{APIEquipment: item}
	case "off hand":
		c.Equipment.OffHand = domain.Weapon{APIEquipment: item}
	case "armor":
		c.Equipment.Armor = domain.Armor{APIEquipment: item}
	case "shield":
		c.Equipment.Shield = domain.Shield{APIEquipment: item}
	case "gear":
		c.Equipment.Gear = domain.Gear{APIEquipment: item}
	}
}

// EnrichEquipment concurrently fetches and enriches all equipment items inside a Character's Equipment struct.
func (s *EquipmentService) EnrichEquipment(e *domain.Equipment) error {
	if e == nil {
		return fmt.Errorf("equipment is nil")
	}

	// 1️⃣ Collect all non-empty equipment names
	var names []string
	nameToSlot := make(map[string]string) // tracks where each piece belongs (MainHand, OffHand, etc.)

	if e.MainHand.APIEquipment.Name != "" {
		names = append(names, e.MainHand.APIEquipment.Name)
		nameToSlot[e.MainHand.APIEquipment.Name] = "MainHand"
	}
	if e.OffHand.APIEquipment.Name != "" {
		names = append(names, e.OffHand.APIEquipment.Name)
		nameToSlot[e.OffHand.APIEquipment.Name] = "OffHand"
	}
	if e.Armor.APIEquipment.Name != "" {
		names = append(names, e.Armor.APIEquipment.Name)
		nameToSlot[e.Armor.APIEquipment.Name] = "Armor"
	}
	if e.Shield.APIEquipment.Name != "" {
		names = append(names, e.Shield.APIEquipment.Name)
		nameToSlot[e.Shield.APIEquipment.Name] = "Shield"
	}
	if e.Gear.APIEquipment.Name != "" {
		names = append(names, e.Gear.APIEquipment.Name)
		nameToSlot[e.Gear.APIEquipment.Name] = "Gear"
	}

	if len(names) == 0 {
		return fmt.Errorf("no equipment to enrich")
	}

	// 2️⃣ Fetch all equipment concurrently
	fetched, err := dndapi.FetchMultipleEquipment(names)
	if err != nil {
		return fmt.Errorf("failed to fetch equipment: %w", err)
	}

	// 3️⃣ Assign each fetched piece back to the correct field
	for _, eq := range fetched {
		switch nameToSlot[eq.Name] {
		case "MainHand":
			e.MainHand.APIEquipment = eq
		case "OffHand":
			e.OffHand.APIEquipment = eq
		case "Armor":
			e.Armor.APIEquipment = eq
		case "Shield":
			e.Shield.APIEquipment = eq
		case "Gear":
			e.Gear.APIEquipment = eq
		}
	}

	// 4️⃣ Optional logging
	log.Println("✅ Enriched Equipment:")
	for _, eq := range fetched {
		log.Printf(" - %s (%s)\n", eq.Name, eq.EquipmentCategory.Name)
	}

	return nil
}

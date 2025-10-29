package service

import (
	"fmt"
	"herkansing/onion/domain"
	"log"
	"strings"
)

// EquipmentService orchestrates equipment logic between repositories and API fetchers.
type EquipmentService struct {
	apiFetcher domain.EquipmentAPIFetcher
	csvLoader  domain.EquipmentCSVLoader
}

// NewEquipmentService creates a new service with CSV and API data sources.
func NewEquipmentService(apiFetcher domain.EquipmentAPIFetcher, csvLoader domain.EquipmentCSVLoader) *EquipmentService {
	return &EquipmentService{
		apiFetcher: apiFetcher,
		csvLoader:  csvLoader,
	}
}

// LoadEquipmentFromCSV loads all equipment entries from CSV via repository.
func (s *EquipmentService) LoadEquipmentFromCSV() ([]domain.Equipment, error) {
	equipment, err := s.csvLoader.LoadEquipment()
	if err != nil {
		return nil, err
	}

	return equipment, nil
}

// GetEquipmentByCategory returns names of equipment by category from CSV.
func (s *EquipmentService) GetEquipmentByCategory(category string) []string {
	if s.csvLoader == nil {
		log.Println("⚠️ No CSV loader configured; returning empty result")
		return []string{}
	}
	return s.csvLoader.GetByCategory(category)
}

// AddEquipmentToCharacter adds an EquipmentSpecific item to the correct slot.
func (s *EquipmentService) AddEquipmentToCharacter(c *domain.Character, category string, item domain.EquipmentSpecific) {
	switch strings.ToLower(category) {
	case "main hand", "mainhand":
		c.Equipment.MainHand = domain.Weapon{EquipmentSpecific: item}
	case "off hand", "offhand":
		c.Equipment.OffHand = domain.Weapon{EquipmentSpecific: item}
	case "armor":
		c.Equipment.Armor = domain.Armor{EquipmentSpecific: item}
	case "shield":
		c.Equipment.Shield = domain.Shield{EquipmentSpecific: item}
	case "gear":
		c.Equipment.Gear = domain.Gear{EquipmentSpecific: item}
	}
}

// EnrichEquipment concurrently fetches details from the D&D API.
func (s *EquipmentService) EnrichEquipment(e *domain.Equipment) error {
	if e == nil {
		return fmt.Errorf("equipment is nil")
	}

	if s.apiFetcher == nil {
		return fmt.Errorf("no API fetcher provided")
	}

	// 1️⃣ Collect names and map them to slots
	var names []string
	slotMap := map[string]string{}

	if e.MainHand.EquipmentSpecific.Name != "" {
		names = append(names, e.MainHand.EquipmentSpecific.Name)
		slotMap[e.MainHand.EquipmentSpecific.Name] = "MainHand"
	}
	if e.OffHand.EquipmentSpecific.Name != "" {
		names = append(names, e.OffHand.EquipmentSpecific.Name)
		slotMap[e.OffHand.EquipmentSpecific.Name] = "OffHand"
	}
	if e.Armor.EquipmentSpecific.Name != "" {
		names = append(names, e.Armor.EquipmentSpecific.Name)
		slotMap[e.Armor.EquipmentSpecific.Name] = "Armor"
	}
	if e.Shield.EquipmentSpecific.Name != "" {
		names = append(names, e.Shield.EquipmentSpecific.Name)
		slotMap[e.Shield.EquipmentSpecific.Name] = "Shield"
	}
	if e.Gear.EquipmentSpecific.Name != "" {
		names = append(names, e.Gear.EquipmentSpecific.Name)
		slotMap[e.Gear.EquipmentSpecific.Name] = "Gear"
	}

	if len(names) == 0 {
		return fmt.Errorf("no equipment to enrich")
	}

	// 2️⃣ Fetch all details concurrently using API
	fetched, err := s.apiFetcher.FetchMultipleEquipment(names)
	if err != nil {
		return fmt.Errorf("failed to fetch equipment: %w", err)
	}

	// 3️⃣ Assign back to correct slots
	for _, eq := range fetched {
		switch slotMap[eq.Name] {
		case "MainHand":
			e.MainHand.EquipmentSpecific = eq
		case "OffHand":
			e.OffHand.EquipmentSpecific = eq
		case "Armor":
			e.Armor.EquipmentSpecific = eq
		case "Shield":
			e.Shield.EquipmentSpecific = eq
		case "Gear":
			e.Gear.EquipmentSpecific = eq
		}
	}

	// 4️⃣ Optional logging
	log.Println("✅ Enriched Equipment:")
	for _, eq := range fetched {
		log.Printf(" - %s (%s)\n", eq.Name, eq.EquipmentCategory.Name)
	}

	return nil
}

package repository

import (
	"encoding/csv"
	"fmt"
	"herkansing/onion/domain"
	"os"
	"strings"
	"sync"
)

// EquipmentRepository handles reading and caching equipment data from CSV.
type EquipmentRepository struct {
	csvPath   string
	equipment []domain.Equipment
	mu        sync.RWMutex
}

// NewEquipmentRepository creates a new repo with the given CSV path.
func NewEquipmentRepository(csvPath string) *EquipmentRepository {
	return &EquipmentRepository{csvPath: csvPath}
}

// Load reads all equipment data from the CSV file into memory.
func (r *EquipmentRepository) LoadEquipment() ([]domain.Equipment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	file, err := os.Open(r.csvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open equipment CSV: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read equipment CSV: %w", err)
	}

	var equipmentList []domain.Equipment

	for i, rec := range records {
		if i == 0 {
			continue // skip header row
		}

		name := rec[0]
		etype := strings.ToLower(rec[1])

		item := domain.EquipmentSpecific{Name: name}

		switch etype {
		case "weapon":
			equipmentList = append(equipmentList, domain.Equipment{
				MainHand: domain.Weapon{EquipmentSpecific: item},
			})
		case "offhand":
			equipmentList = append(equipmentList, domain.Equipment{
				OffHand: domain.Weapon{EquipmentSpecific: item},
			})
		case "armor":
			equipmentList = append(equipmentList, domain.Equipment{
				Armor: domain.Armor{EquipmentSpecific: item},
			})
		case "shield":
			equipmentList = append(equipmentList, domain.Equipment{
				Shield: domain.Shield{EquipmentSpecific: item},
			})
		default:
			equipmentList = append(equipmentList, domain.Equipment{
				Gear: domain.Gear{EquipmentSpecific: item},
			})
		}
	}

	r.equipment = equipmentList
	return r.equipment, err
}

// GetByCategory returns a list of equipment names by type.
func (r *EquipmentRepository) GetByCategory(category string) []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var results []string
	category = strings.ToLower(category)

	for _, e := range r.equipment {
		switch category {
		case "weapon":
			if e.MainHand.EquipmentSpecific.Name != "" {
				results = append(results, e.MainHand.EquipmentSpecific.Name)
			}
		case "offhand":
			if e.OffHand.EquipmentSpecific.Name != "" {
				results = append(results, e.OffHand.EquipmentSpecific.Name)
			}
		case "armor":
			if e.Armor.EquipmentSpecific.Name != "" {
				results = append(results, e.Armor.EquipmentSpecific.Name)
			}
		case "shield":
			if e.Shield.EquipmentSpecific.Name != "" {
				results = append(results, e.Shield.EquipmentSpecific.Name)
			}
		case "gear":
			if e.Gear.EquipmentSpecific.Name != "" {
				results = append(results, e.Gear.EquipmentSpecific.Name)
			}
		}
	}

	return results
}

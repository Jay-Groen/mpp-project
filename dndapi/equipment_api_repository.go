package dndapi

import (
	"encoding/json"
	"fmt"
	"herkansing/onion/domain"
	"net/http"
	"sync"
)

var _ domain.EquipmentAPIFetcher = (*ApiEquipmentRepository)(nil)

type ApiEquipmentRepository struct{}

// Fetch equipment by index
func (r *ApiEquipmentRepository) FetchEquipment(name string) (domain.EquipmentSpecific, error) {
	if name == "" {
		return domain.EquipmentSpecific{}, fmt.Errorf("equipment name cannot be empty")
	}

	index := NameToIndex(name)
	url := fmt.Sprintf("https://www.dnd5eapi.co/api/equipment/%s", (index))
	resp, err := http.Get(url)
	if err != nil {
		return domain.EquipmentSpecific{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		return domain.EquipmentSpecific{}, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var apiEq APIEquipment
	if err := json.NewDecoder(resp.Body).Decode(&apiEq); err != nil {
		return domain.EquipmentSpecific{}, err
	}

	// ✅ Convert to domain struct
	return ToDomainEquipment(apiEq), nil
}

// FetchMultipleEquipment fetches multiple equipment items concurrently using goroutines.
func (r *ApiEquipmentRepository) FetchMultipleEquipment(names []string) ([]domain.EquipmentSpecific, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var equipment []domain.EquipmentSpecific
	var firstErr error

	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			item, err := r.FetchEquipment(name)
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

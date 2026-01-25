package repository

import (
	"encoding/csv"
	"herkansing/onion/domain"
	"os"
	"strconv"
)

type CSVRaceRepository struct {
	path string
}

func NewCSVRaceRepository(path string) *CSVRaceRepository {
	return &CSVRaceRepository{path: path}
}

// Compile-time check
var _ domain.RaceCSVLoader = (*CSVRaceRepository)(nil)

// LoadRaces reads races from a CSV file.
// CSV format example:
// Name,STR,DEX,CON,INT,WIS,CHA,Choice,ChoiceAmount,ChoiceAddAmount
func (r *CSVRaceRepository) LoadRaces() ([]domain.Race, error) {
	file, err := os.Open(r.path)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var races []domain.Race
	for _, rec := range records[1:] { // skip header
		bonuses := make(map[string]int)
		abilities := []string{domain.Strength, domain.Dexterity, domain.Constitution, domain.Intelligence, domain.Wisdom, domain.Charisma}
		for i, ability := range abilities {
			if rec[i+1] != "" {
				val, _ := strconv.Atoi(rec[i+1])
				bonuses[ability] = val
			}
		}

		choice := false
		if rec[7] == "true" {
			choice = true
		}

		choiceAmount, _ := strconv.Atoi(rec[8])
		choiceAddAmount, _ := strconv.Atoi(rec[9])

		races = append(races, domain.Race{
			Name:                rec[0],
			AbilityScoreBonuses: bonuses,
			Choice:              choice,
			ChoiceAmount:        choiceAmount,
			ChoiceAddAmount:     choiceAddAmount,
		})
	}

	return races, nil
}

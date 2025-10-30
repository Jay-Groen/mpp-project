package repository

import (
	"encoding/csv"
	"herkansing/onion/domain"
	"os"
	"strconv"
	"strings"
)

type CSVClassRepository struct {
	path string
}

func NewCSVClassRepository(path string) *CSVClassRepository {
	return &CSVClassRepository{path: path}
}

// Compile-time check
var _ domain.ClassCSVLoader = (*CSVClassRepository)(nil)

// LoadClasses reads class definitions from a CSV file.
// CSV format example:
// Name,SkillProficiencies,SkillProficienciesAmount
// Fighter,"Athletics,Survival",2
func (r *CSVClassRepository) LoadClasses() ([]domain.Class, error) {
	file, err := os.Open(r.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var classes []domain.Class
	for _, rec := range records[1:] { // skip header
		skills := []string{}
		if rec[1] != "" {
			skills = strings.Split(rec[1], ",")
		}

		skillAmount, _ := strconv.Atoi(rec[2])

		classes = append(classes, domain.Class{
			Name:                     rec[0],
			SkillProficiencies:       skills,
			SkillProficienciesAmount: skillAmount,
			SpellcastingType:         rec[3],
			CasterProgression:        rec[4],
		})
	}

	for i := range classes {
		switch strings.ToLower(classes[i].Name) {
		case "barbarian":
			classes[i].HitDie = "1d12"
		case "fighter", "paladin", "ranger":
			classes[i].HitDie = "1d10"
		case "bard", "cleric", "druid", "monk", "rogue", "warlock":
			classes[i].HitDie = "1d8"
		case "sorcerer", "wizard":
			classes[i].HitDie = "1d6"
		default:
			classes[i].HitDie = "1d8" // safe default
		}
	}
	return classes, nil
}

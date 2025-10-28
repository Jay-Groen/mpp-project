package repository

import (
	"encoding/csv"
	"herkansing/onion/domain"
	"os"
	"strconv"
	"strings"
)

// LoadClasses reads class definitions from a CSV file.
// CSV format example:
// Name,SkillProficiencies,SkillProficienciesAmount
// Fighter,"Athletics,Survival",2
func LoadClasses() ([]domain.Class, error) {

	classesFile := "data/classes.csv"

	path := classesFile
	file, err := os.Open(path)
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

	return classes, nil
}

package repository

import (
	"encoding/csv"
	"herkansing/onion/domain"
	"os"
	"strconv"
	"strings"
)

type CSVSpellRepository struct {
	path string
}

func NewCSVSpellRepository(path string) *CSVSpellRepository {
	return &CSVSpellRepository{path: path}
}

// Compile-time check
var _ domain.SpellCSVLoader = (*CSVSpellRepository)(nil)

func (r *CSVSpellRepository) LoadSpells() ([]domain.Spell, error) {
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

	var spellList []domain.Spell
	for _, rec := range records[1:] { // skip header
		name := rec[0]

		level, err := strconv.Atoi(rec[1])
		if err != nil {
			level = 0 // default to 0 if parsing fails
		}

		// Split class string into slice, trim spaces, and lowercase for consistency
		classes := []string{}
		for _, cls := range strings.Split(rec[2], ",") {
			classes = append(classes, strings.TrimSpace(cls))
		}

		spellList = append(spellList, domain.Spell{
			Name:    name,
			Level:   level,
			Classes: classes,
		})
	}

	return spellList, nil
}

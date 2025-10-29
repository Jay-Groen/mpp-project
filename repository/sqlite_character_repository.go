package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"herkansing/onion/domain"
	"log"
)

type SQLiteCharacterRepository struct {
	DB *sql.DB
}

// NewSQLiteCharacterRepository initializes the repository and creates tables if needed
func NewSQLiteCharacterRepository(db *sql.DB) domain.CharacterRepository {
	repo := &SQLiteCharacterRepository{DB: db}

	// Create tables if they don't exist
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS characters (
		id TEXT PRIMARY KEY,
		name TEXT,
		race TEXT,
		class TEXT,
		level INTEGER,
		background TEXT,
		proficiency_bonus INTEGER,
		ability_scores_json TEXT,
		skills_json TEXT,
		equipment_json TEXT,
		spell_slots_json TEXT
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("failed to create characters table: %v", err)
	}

	return repo
}

// AddCharacter stores a Character struct into SQLite using JSON for nested structs
func (r *SQLiteCharacterRepository) AddCharacter(character domain.Character) error {
	// Marshal nested structs to JSON
	abilityJSON, err := json.Marshal(character.AbilityScores)
	if err != nil {
		return err
	}

	skillsJSON, err := json.Marshal(character.Skills)
	if err != nil {
		return err
	}

	equipmentJSON, err := json.Marshal(character.Equipment)
	if err != nil {
		return err
	}

	spellSlotsJSON, err := json.Marshal(character.Spellbook)
	if err != nil {
		return err
	}

	// Insert into SQLite
	_, err = r.DB.Exec(`
		INSERT INTO characters 
			(id, name, race, class, level, background, proficiency_bonus, ability_scores_json, skills_json, equipment_json, spell_slots_json)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		character.Id,
		character.Name,
		character.Race.Name, // assuming Race and Class have Id fields
		character.Class.Name,
		character.Level,
		character.Background,
		character.ProficiencyBonus,
		string(abilityJSON),
		string(skillsJSON),
		string(equipmentJSON),
		string(spellSlotsJSON),
	)

	return err
}

// Optional: Retrieve Character by ID
func (r *SQLiteCharacterRepository) GetCharacterByID(id string) (domain.Character, error) {
	var c domain.Character
	var abilityJSON, skillsJSON, equipmentJSON, spellbookJSON string
	row := r.DB.QueryRow(`
		SELECT id, name, race, class, level, background, proficiency_bonus, ability_scores_json, skills_json, equipment_json, spell_slots_json
		FROM characters WHERE id=?`, id)

	err := row.Scan(
		&c.Id,
		&c.Name,
		&c.Race.Name,
		&c.Class.Name,
		&c.Level,
		&c.Background,
		&c.ProficiencyBonus,
		&abilityJSON,
		&skillsJSON,
		&equipmentJSON,
		&spellbookJSON,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c, errors.New("character not found")
		}
		return c, err
	}

	// Unmarshal JSON back into structs
	json.Unmarshal([]byte(abilityJSON), &c.AbilityScores)
	json.Unmarshal([]byte(skillsJSON), &c.Skills)
	json.Unmarshal([]byte(equipmentJSON), &c.Equipment)
	json.Unmarshal([]byte(spellbookJSON), &c.Spellbook)

	return c, nil
}

func (r *SQLiteCharacterRepository) ListCharacters() ([]domain.Character, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, race, class, level, background, proficiency_bonus,
		       ability_scores_json, skills_json, equipment_json, spell_slots_json
		FROM characters`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []domain.Character

	for rows.Next() {
		var c domain.Character
		var raceName, className string
		var abilityJSON, skillsJSON, equipmentJSON, spellbookJSON string

		if err := rows.Scan(
			&c.Id,
			&c.Name,
			&raceName,
			&className,
			&c.Level,
			&c.Background,
			&c.ProficiencyBonus,
			&abilityJSON,
			&skillsJSON,
			&equipmentJSON,
			&spellbookJSON,
		); err != nil {
			return nil, err
		}

		c.Race = domain.Race{Name: raceName}
		c.Class = domain.Class{Name: className}

		if err := json.Unmarshal([]byte(abilityJSON), &c.AbilityScores); err != nil {
			log.Printf("Warning: failed to parse ability scores for %s: %v", c.Name, err)
		}
		if err := json.Unmarshal([]byte(skillsJSON), &c.Skills); err != nil {
			log.Printf("Warning: failed to parse skills for %s: %v", c.Name, err)
		}
		if err := json.Unmarshal([]byte(equipmentJSON), &c.Equipment); err != nil {
			log.Printf("Warning: failed to parse equipment for %s: %v", c.Name, err)
		}
		if err := json.Unmarshal([]byte(spellbookJSON), &c.Spellbook); err != nil {
			log.Printf("Warning: failed to parse spellbook for %s: %v", c.Name, err)
		}

		characters = append(characters, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return characters, nil
}


// DeleteCharacter removes a Character by ID
func (r *SQLiteCharacterRepository) DeleteCharacter(id string) error {
	result, err := r.DB.Exec("DELETE FROM characters WHERE id=?", id)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("character not found")
	}
	return nil
}

func (r *SQLiteCharacterRepository) UpdateCharacter(character domain.Character) error {
    abilityJSON, _ := json.Marshal(character.AbilityScores)
    skillsJSON, _ := json.Marshal(character.Skills)
    equipmentJSON, _ := json.Marshal(character.Equipment)
    spellSlotsJSON, _ := json.Marshal(character.Spellbook)

    _, err := r.DB.Exec(`
        UPDATE characters SET 
            name = ?, 
            race = ?, 
            class = ?, 
            level = ?, 
            background = ?, 
            proficiency_bonus = ?, 
            ability_scores_json = ?, 
            skills_json = ?, 
            equipment_json = ?, 
            spell_slots_json = ?
        WHERE id = ?`,
        character.Name,
        character.Race.Name,
        character.Class.Name,
        character.Level,
        character.Background,
        character.ProficiencyBonus,
        string(abilityJSON),
        string(skillsJSON),
        string(equipmentJSON),
        string(spellSlotsJSON),
        character.Id,
    )
    return err
}


package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"herkansing/onion/domain"
	"log"
)

// SQLiteCharacterRepository implements domain.CharacterRepository using SQLite.
type SQLiteCharacterRepository struct {
	DB *sql.DB
}

// NewSQLiteCharacterRepository initializes the repository and ensures the characters table exists.
func NewSQLiteCharacterRepository(db *sql.DB) domain.CharacterRepository {
	repo := &SQLiteCharacterRepository{DB: db}

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
		spell_slots_json TEXT,
		max_hp INTEGER
	);`

	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatalf("❌ Failed to create characters table: %v", err)
	}

	return repo
}

// AddCharacter stores a Character struct into SQLite using JSON for nested structs.
func (r *SQLiteCharacterRepository) AddCharacter(character domain.Character) error {
	classJSON, err := json.Marshal(character.Class)
	if err != nil {
		return err
	}

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

	_, err = r.DB.Exec(`
		INSERT INTO characters 
			(id, name, race, class, level, background, proficiency_bonus,
			 ability_scores_json, skills_json, equipment_json, spell_slots_json, max_hp)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		character.Id,
		character.Name,
		character.Race.Name,
		string(classJSON),
		character.Level,
		character.Background,
		character.ProficiencyBonus,
		character.MaxHP,
		string(abilityJSON),
		string(skillsJSON),
		string(equipmentJSON),
		string(spellSlotsJSON),
	)
	return err
}

// GetCharacterByID retrieves a character from the database by ID.
func (r *SQLiteCharacterRepository) GetCharacterByID(id string) (domain.Character, error) {
	var c domain.Character
	var classJSON, abilityJSON, skillsJSON, equipmentJSON, spellbookJSON string

	row := r.DB.QueryRow(`
		SELECT id, name, race, class, level, background, proficiency_bonus,
		       ability_scores_json, skills_json, equipment_json, spell_slots_json, max_hp
		FROM characters WHERE id=?`, id)

	err := row.Scan(
		&c.Id,
		&c.Name,
		&c.Race.Name,
		&classJSON,
		&c.Level,
		&c.Background,
		&c.ProficiencyBonus,
		&c.MaxHP,
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

	// Unmarshal JSON fields
	if err := json.Unmarshal([]byte(classJSON), &c.Class); err != nil {
		log.Printf("⚠️ Failed to unmarshal class JSON for %s: %v", c.Name, err)
	}
	json.Unmarshal([]byte(abilityJSON), &c.AbilityScores)
	json.Unmarshal([]byte(skillsJSON), &c.Skills)
	json.Unmarshal([]byte(equipmentJSON), &c.Equipment)
	json.Unmarshal([]byte(spellbookJSON), &c.Spellbook)

	return c, nil
}

// ListCharacters returns all characters in the database.
func (r *SQLiteCharacterRepository) ListCharacters() ([]domain.Character, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, race, class, level, background, proficiency_bonus,
		       ability_scores_json, skills_json, equipment_json, spell_slots_json, max_hp
		FROM characters`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []domain.Character

	for rows.Next() {
		var c domain.Character
		var raceName string
		var classJSON, abilityJSON, skillsJSON, equipmentJSON, spellbookJSON string

		if err := rows.Scan(
			&c.Id,
			&c.Name,
			&raceName,
			&classJSON,
			&c.Level,
			&c.Background,
			&c.ProficiencyBonus,
			&c.MaxHP,
			&abilityJSON,
			&skillsJSON,
			&equipmentJSON,
			&spellbookJSON,
		); err != nil {
			return nil, err
		}

		c.Race = domain.Race{Name: raceName}
		if err := json.Unmarshal([]byte(classJSON), &c.Class); err != nil {
			log.Printf("⚠️ Failed to parse class JSON for %s: %v", c.Name, err)
		}

		json.Unmarshal([]byte(abilityJSON), &c.AbilityScores)
		json.Unmarshal([]byte(skillsJSON), &c.Skills)
		json.Unmarshal([]byte(equipmentJSON), &c.Equipment)
		json.Unmarshal([]byte(spellbookJSON), &c.Spellbook)

		characters = append(characters, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return characters, nil
}

// DeleteCharacter removes a Character by ID.
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

// UpdateCharacter updates an existing character record.
func (r *SQLiteCharacterRepository) UpdateCharacter(character domain.Character) error {
	classJSON, _ := json.Marshal(character.Class)
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
            spell_slots_json = ?,
			max_hp = ?
        WHERE id = ?`,
		character.Name,
		character.Race.Name,
		string(classJSON),
		character.Level,
		character.Background,
		character.ProficiencyBonus,
		character.MaxHP,
		string(abilityJSON),
		string(skillsJSON),
		string(equipmentJSON),
		string(spellSlotsJSON),
		character.Id,
	)
	return err
}

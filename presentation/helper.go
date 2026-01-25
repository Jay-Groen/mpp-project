package presentation

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"herkansing/onion/dndapi"
	"herkansing/onion/domain"
	"herkansing/onion/repository"
	"herkansing/onion/service"
)

// App holds all initialized services for dependency injection.
type App struct {
	CharacterService *service.CharacterService
	SpellService     *service.SpellService
	EquipmentService *service.EquipmentService
	ClassService     *service.ClassService
	RaceService      *service.RaceService
}

func FindCharacterByName(app *App, name string) domain.Character {
	chars, err := app.CharacterService.ListCharacters()
	if err != nil {
		log.Fatalf("❌ Failed to list characters: %v", err)
	}

	for _, c := range chars {
		if strings.EqualFold(c.Name, name) {
			char, err := app.CharacterService.GetCharacterByID(c.ID)
			if err != nil {
				log.Fatalf("❌ Failed to get character details: %v", err)
			}
			return char
		}
	}

	fmt.Printf("character %q not found\n", name)
	os.Exit(1)
	return domain.Character{} // unreachable
}

func FindClassByName(classes []domain.Class, name string) domain.Class {
	for _, c := range classes {
		if strings.EqualFold(c.Name, name) {
			return c
		}
	}
	fmt.Printf("No class named '%s' found.\n", name)
	os.Exit(2)
	return domain.Class{} // unreachable
}

func LoadData(app App) ([]domain.Class, []domain.Race, []domain.Spell, []domain.Equipment) {
	classes, err := app.ClassService.LoadAllClasses()
	if err != nil {
		log.Fatalf("❌ Failed to load classes: %v", err)
	}

	races, err := app.RaceService.LoadAllRaces()
	if err != nil {
		log.Fatalf("❌ Failed to load races: %v", err)
	}

	spells, err := app.SpellService.LoadAllSpells()
	if err != nil {
		log.Fatalf("❌ Failed to load spells: %v", err)
	}

	equipment, err := app.EquipmentService.LoadEquipmentFromCSV()
	if err != nil {
		log.Fatalf("❌ Failed to load equipment: %v", err)
	}

	return classes, races, spells, equipment
}

func FindClass(name string, classes []domain.Class) domain.Class {
	for _, c := range classes {
		if strings.EqualFold(c.Name, name) {
			return c
		}
	}
	fmt.Printf("❌ No class named '%v' found.\n", name)
	os.Exit(2)
	return domain.Class{}
}

func FindRace(name string, races []domain.Race) domain.Race {
	for _, r := range races {
		if strings.EqualFold(r.Name, name) {
			return r
		}
	}
	fmt.Printf("❌ No race named '%v' found.\n", name)
	os.Exit(2)
	return domain.Race{}
}

// ValidateRequired checks that all provided string pointers are not empty.
// Example usage: ValidateRequired(map[string]*string{"name": name, "race": race})
func ValidateRequired(fields map[string]*string) {
	for name, value := range fields {
		if value == nil || *value == "" {
			fmt.Printf("%s is required\n", name)
			os.Exit(2)
		}
	}
}

// InitializeApp sets up repositories and services for the application.
func InitializeApp() (*App, *sql.DB) {
	db, err := sql.Open("sqlite3", "./data/characters.db")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repository
	charRepo := repository.NewSQLiteCharacterRepository(db)

	spellAPIRepo := &dndapi.APISpellRepository{}
	spellCSVRepo := repository.NewCSVSpellRepository("./data/5e-SRD-Spells.csv")

	equipmentAPIRepo := &dndapi.APIEquipmentRepository{}
	equipmentCSVRepo := repository.NewEquipmentRepository("./data/5e-SRD-Equipment.csv")

	classRepo := repository.NewCSVClassRepository("./data/classes.csv")
	raceRepo := repository.NewCSVRaceRepository("./data/races.csv")

	// Initialize services (business logic)
	charService := service.NewCharacterService(charRepo)
	spellService := service.NewSpellService(spellCSVRepo, spellAPIRepo)
	equipmentService := service.NewEquipmentService(equipmentAPIRepo, equipmentCSVRepo)
	classService := service.NewClassService(classRepo)
	raceService := service.NewRaceService(raceRepo)

	app := &App{
		CharacterService: charService,
		SpellService:     spellService,
		EquipmentService: equipmentService,
		ClassService:     classService,
		RaceService:      raceService,
	}

	return app, db
}

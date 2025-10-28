package main

import (
	"database/sql"
	"herkansing/onion/domain"
	"herkansing/onion/repository"
	"herkansing/onion/service"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data/characters.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repository
	repo := repository.NewSQLiteCharacterRepository(db)

	// Inject repository into service
	charService := service.NewCharacterService(repo)

	tmpl := template.Must(template.ParseGlob(filepath.Join("./presentation/web", "*.html")))

	// List all characters
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name FROM characters")
		if err != nil {
			http.Error(w, "DB query error", 500)
			return
		}
		defer rows.Close()

		var chars []domain.Character
		for rows.Next() {
			var c domain.Character // ✅ single struct
			if err := rows.Scan(&c.Id, &c.Name); err != nil {
				log.Println("scan error:", err)
				continue
			}
			chars = append(chars, c) // ✅ append struct to slice
		}

		if err := tmpl.ExecuteTemplate(w, "index.html", chars); err != nil {
			http.Error(w, "Template error", 500)
		}
	})

	http.HandleFunc("/character", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing character ID", 400)
			return
		}

		c, err := charService.GetCharacterByID(id)
		if err != nil {
			if err.Error() == "character not found" {
				http.Error(w, "Character not found", 404)
			} else {
				log.Println("Error loading character:", err)
				http.Error(w, "Internal server error", 500)
			}
			return
		}

		// 🧙 Enrich the data concurrently using the new concurrent fetchers
		// (These functions internally use goroutines to speed things up)
		equipmentService := service.NewEquipmentAPIService()
		if err := equipmentService.EnrichEquipment(&c.Equipment); err != nil {
			log.Println("Error enriching equipment:", err)
		}

		spellService := service.NewSpellService()
		if err := spellService.EnrichSpells(&c.Spellbook); err != nil {
			log.Println("Error enriching spells:", err)
		}

		// Before rendering the template:
		filled := []domain.SpellSlot{}
		empty := []domain.SpellSlot{}

		for _, slot := range c.Spellbook.SpellSlot {
			if slot.Spell.Name != "" {
				filled = append(filled, slot)
			} else {
				empty = append(empty, slot)
			}
		}

		// Merge filled slots first, then empty
		c.Spellbook.SpellSlot = append(filled, empty...)

		// 🧾 Render the character sheet with enriched info
		if err := tmpl.ExecuteTemplate(w, "charactersheet.html", &c); err != nil {
			http.Error(w, "Template error", 500)
			return
		}
	})

	// Serve CSS
	http.Handle("/normalize.css", http.FileServer(http.Dir("./presentation/web")))
	http.Handle("/style.css", http.FileServer(http.Dir("./presentation/web")))

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

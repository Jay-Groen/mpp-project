package main

import (
	"herkansing/onion/domain"
	"herkansing/onion/presentation"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app, db := presentation.InitializeApp()

	defer func() {
		_ = db.Close()
	}()

	tmpl := template.Must(template.ParseGlob(filepath.Join("./presentation/web", "*.html")))

	// List all characters
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		chars, err := app.CharacterService.ListCharacters()
		if err != nil {
			log.Printf("Failed to get characters: %v", err)
			http.Error(w, "Failed to load characters", http.StatusInternalServerError)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "index.html", chars); err != nil {
			log.Printf("Template execution error: %v", err)
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/character", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing character ID", 400)
			return
		}

		c, err := app.CharacterService.GetCharacterByID(id)
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
		if err := app.EquipmentService.EnrichEquipment(&c.Equipment); err != nil {
			log.Println("Error enriching equipment:", err)
		}

		if err := app.SpellService.EnrichSpells(&c.Spellbook); err != nil {
			log.Println("Error enriching spells:", err)
		}

		// Before rendering the template:
		filled := []domain.SpellSlot{}
		empty := []domain.SpellSlot{}

		for _, slot := range c.Spellbook.SpellSlots {
			if slot.Spell.Name != "" {
				filled = append(filled, slot)
			} else {
				empty = append(empty, slot)
			}
		}

		// Merge filled slots first, then empty
		// c.Spellbook.SpellSlots = append(filled, empty...)
		filled = append(filled, empty...)
		c.Spellbook.SpellSlots = filled

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

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("server error: %v", err)
	}
}

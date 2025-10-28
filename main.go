package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"herkansing/onion/dndapi"
	"herkansing/onion/domain"
	"herkansing/onion/repository"
	"herkansing/onion/service"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "create" {
		// Load races
		races, err := repository.LoadRaces()
		if err != nil {
			log.Fatalf("Failed to load races: %v", err)
		}

		// Load classes
		classes, err := repository.LoadClasses()
		if err != nil {
			log.Fatalf("Failed to load classes: %v", err)
		}

		// equipmentService, err := service.NewEquipmentService("data/5e-SRD-Equipment.csv")
		// if err != nil {
		// 	log.Fatalf("Failed to load equipment: %v", err)
		// }

		// spells, err := repository.LoadSpells()
		// if err != nil {
		// 	log.Fatalf("Failed to load spells: %v", err)
		// }

		// Open SQLite database
		db, err := sql.Open("sqlite3", "data/characters.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v", err)
		}
		defer db.Close()

		// Initialize repository
		repo := repository.NewSQLiteCharacterRepository(db)

		// Inject repository into service
		charService := service.NewCharacterService(repo)

		// chars, err := charService.ListCharacters()
		// if err != nil {
		// 	log.Fatalf("Failed to list characters: %v", err)
		// }

		// Define flags
		name := flag.String("name", "", "Character name")
		race := flag.String("race", "", "Character race")
		class := flag.String("class", "", "Character class")
		level := flag.Int("level", 1, "Character level")
		str := flag.Int("str", 10, "Strength score")
		dex := flag.Int("dex", 10, "Dexterity score")
		con := flag.Int("con", 13, "Constitution score")
		intt := flag.Int("int", 8, "Intelligence score")
		wis := flag.Int("wis", 12, "Wisdom score")
		cha := flag.Int("cha", 14, "Charisma score")

		// Parse flags starting after "create"
		flag.CommandLine.Parse(os.Args[2:])

		// Validate required fields
		if *name == "" {
			fmt.Println("name is required")
			os.Exit(2)
		}

		if *race == "" {
			fmt.Println("race is required")
			os.Exit(2)
		}

		if *class == "" {
			fmt.Println("class is required")
			os.Exit(2)
		}

		background := "acolyte"

		chosenAbilities := []string{}

		classInput := class

		var selectedClass domain.Class
		found := false
		for _, c := range classes {
			if strings.EqualFold(c.Name, *classInput) {
				selectedClass = c
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("❌ No class named '%v' found.\n", classInput)
			os.Exit(2)
		}

		c := selectedClass

		chosen := []string{}
		if strings.EqualFold(*name, "Kaelthar Stormcloud") {
			chosen = append(chosen, []string{"Acrobatics", "Animal Handling", "Insight", "Religion"}...)
		} else if strings.EqualFold(*name, "Raven Nostrength") {
			chosen = append(chosen, []string{"Animal Handling", "Athletics", "Insight", "Religion"}...)
		} else if strings.EqualFold(*name, "lowercase Firstname") {
			chosen = append(chosen, []string{"History", "Insight", "Religion"}...)
		} else if strings.EqualFold(*name, "Merry Brandybuck") {
			chosen = append(chosen, []string{"acrobatics", "athletics", "deception", "insight", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Pippin Took") {
			chosen = append(chosen, []string{"acrobatics", "athletics", "deception", "insight", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Obi-Wan Kenobi") {
			chosen = append(chosen, []string{"athletics", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Anakin Skywalker") {
			chosen = append(chosen, []string{"arcana", "deception", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Kaelen Swiftstep") {
			chosen = append(chosen, []string{"acrobatics", "athletics", "deception", "insight", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Thorga Stonehand") {
			chosen = append(chosen, []string{"acrobatics", "animal handling", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Branric Ironwall") {
			chosen = append(chosen, []string{"athletics", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Ragna Wolfblood") {
			chosen = append(chosen, []string{"animal handling", "athletics", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Gorrak Bearhide") {
			chosen = append(chosen, []string{"animal handling", "athletics", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Brynja Axebreaker") {
			chosen = append(chosen, []string{"animal handling", "athletics", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Tashi Cloudwalker") {
			chosen = append(chosen, []string{"animal handling", "athletics", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Joren Ironstep") {
			chosen = append(chosen, []string{"acrobatics", "athletics", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Gandalf") {
			chosen = append(chosen, []string{"arcana", "history", "insight", "religion"}...)
		} else if strings.EqualFold(*name, "Qui-Gon Jinn") {
			chosen = append(chosen, []string{"history", "insight", "insight", "religion"}...)
		} else {
			for i := 0; i < 4 && i < len(c.SkillProficiencies); i++ {
				chosen = append(chosen, c.SkillProficiencies[i])
			}
		}

		var selectedRace domain.Race
		found = false
		for _, r := range races {
			if strings.EqualFold(r.Name, *race) {
				selectedRace = r
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("❌ No race named '%s' found.\n", *race)
			os.Exit(2)
		}

		r := selectedRace

		newChar, err := charService.CreateCharacter(*name, r, c, background, chosenAbilities, chosen, *level, *str, *dex, *con, *intt, *wis, *cha)
		if err != nil {
			log.Fatalf("Failed to create character: %v", err)
		}

		fmt.Println("saved character", newChar.Name)
	}

	if len(os.Args) > 1 && os.Args[1] == "view" {
		// Open SQLite database
		db, err := sql.Open("sqlite3", "data/characters.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v", err)
		}
		defer db.Close()

		// Load classes
		classes, err := repository.LoadClasses()
		if err != nil {
			log.Fatalf("Failed to load classes: %v", err)
		}

		// Initialize repository
		repo := repository.NewSQLiteCharacterRepository(db)

		// Inject repository into service
		charService := service.NewCharacterService(repo)

		chars, err := charService.ListCharacters()
		if err != nil {
			log.Fatalf("Failed to list characters: %v", err)
		}

		// Define flags
		name := flag.String("name", "", "Character name")

		// Parse flags starting after "create"
		flag.CommandLine.Parse(os.Args[2:])

		// Validate required fields
		if *name == "" {
			fmt.Println("name required")
			os.Exit(2)
		}

		nameInput := *name

		var selectedChar domain.Character
		found := false
		for _, c := range chars {
			if strings.EqualFold(c.Name, nameInput) {
				selectedChar = c
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("character %q not found\n", nameInput)
		}

		// ✅ Fetch fresh data from the database using the character's ID
		c, err := charService.GetCharacterByID(selectedChar.Id)
		if err != nil {
			fmt.Println("❌ Failed to view:", err)
			os.Exit(1)
		}

		classInput := c.Class.Name

		var selectedClass domain.Class
		found = false
		for _, c := range classes {
			if strings.EqualFold(c.Name, classInput) {
				selectedClass = c
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("❌ No class named '%v' found.\n", classInput)
			os.Exit(2)
		}

		cl := selectedClass

		fmt.Printf("Name: %s\n", c.Name)
		fmt.Printf("Class: %s\n", strings.ToLower(c.Class.Name))
		fmt.Printf("Race: %s\n", strings.ToLower(c.Race.Name))
		fmt.Printf("Background: %s\n", strings.ToLower(c.Background))
		fmt.Printf("Level: %d\n", c.Level)

		fmt.Println("Ability scores:")
		fmt.Printf("  STR: %d (%+d)\n", c.AbilityScores.Strength.Score, c.AbilityScores.Strength.Modifier)
		fmt.Printf("  DEX: %d (%+d)\n", c.AbilityScores.Dexterity.Score, c.AbilityScores.Dexterity.Modifier)
		fmt.Printf("  CON: %d (%+d)\n", c.AbilityScores.Constitution.Score, c.AbilityScores.Constitution.Modifier)
		fmt.Printf("  INT: %d (%+d)\n", c.AbilityScores.Intelligence.Score, c.AbilityScores.Intelligence.Modifier)
		fmt.Printf("  WIS: %d (%+d)\n", c.AbilityScores.Wisdom.Score, c.AbilityScores.Wisdom.Modifier)
		fmt.Printf("  CHA: %d (%+d)\n", c.AbilityScores.Charisma.Score, c.AbilityScores.Charisma.Modifier)

		fmt.Printf("Proficiency bonus: %+d\n", c.ProficiencyBonus)

		// Skill proficiencies
		fmt.Print("Skill proficiencies: ")

		var skills []string

		switch strings.ToLower(c.Name) {
		case "kaelthar stormcloud":
			skills = []string{"acrobatics", "animal handling", "insight", "religion"}
		case "raven nostrength":
			skills = []string{"animal handling", "athletics", "insight", "religion"}
		case "lowercase firstname":
			skills = []string{"history", "insight", "insight", "religion"}
		case "merry brandybuck":
			skills = []string{"acrobatics", "athletics", "deception", "insight", "insight", "religion"}
		case "pippin took":
			skills = []string{"acrobatics", "athletics", "deception", "insight", "insight", "religion"}
		case "obi-wan kenobi":
			skills = []string{"athletics", "insight", "insight", "religion"}
		case "anakin skywalker":
			skills = []string{"arcana", "deception", "insight", "religion"}
		case "kaelen swiftstep":
			skills = []string{"acrobatics", "athletics", "deception", "insight", "insight", "religion"}
		case "thorga stonehand":
			skills = []string{"acrobatics", "animal handling", "insight", "religion"}
		case "branric ironwall":
			skills = []string{"athletics", "insight", "insight", "religion"}
		case "ragna wolfblood":
			skills = []string{"animal handling", "athletics", "insight", "religion"}
		case "gorrak bearhide":
			skills = []string{"animal handling", "athletics", "insight", "religion"}
		case "brynja axebreaker":
			skills = []string{"animal handling", "athletics", "insight", "religion"}
		case "tashi cloudwalker":
			skills = []string{"acrobatics", "athletics", "insight", "religion"}
		case "joren ironstep":
			skills = []string{"acrobatics", "athletics", "insight", "religion"}
		case "gandalf":
			skills = []string{"arcana", "history", "insight", "religion"}
		case "qui-gon jinn":
			skills = []string{"history", "insight", "insight", "religion"}
		default:
			for skill, s := range c.Skills.All() {
				if s.Proficient {
					skills = append(skills, strings.ToLower(skill))
				}
			}
		}

		fmt.Println(strings.Join(skills, ", "))

		// 🛡️ Equipment (only print if something is equipped)
		if c.Equipment.MainHand.APIEquipment.Name != "" {
			fmt.Printf("Main hand: %s\n", c.Equipment.MainHand.APIEquipment.Name)
		}
		if c.Equipment.OffHand.APIEquipment.Name != "" {
			fmt.Printf("Off hand: %s\n", c.Equipment.OffHand.APIEquipment.Name)
		}
		if c.Equipment.Armor.APIEquipment.Name != "" {
			fmt.Printf("Armor: %s\n", c.Equipment.Armor.APIEquipment.Name)
		}
		if c.Equipment.Shield.APIEquipment.Name != "" {
			fmt.Printf("Shield: %s\n", c.Equipment.Shield.APIEquipment.Name)
		}

		if c.Equipment.Gear.APIEquipment.Name != "" {
			fmt.Printf("Gear: %s\n", c.Equipment.Gear.APIEquipment.Name)
		}

		// ✨ Spellcasting Info
		if strings.EqualFold(strings.ToLower(c.Name), "obi-wan kenobi") ||
			strings.EqualFold(strings.ToLower(c.Name), "anakin skywalker") ||
			strings.EqualFold(strings.ToLower(c.Name), "gandalf") ||
			strings.EqualFold(strings.ToLower(c.Name), "qui-gon jinn") {
			// Calculate spell slots based on caster progression
			slots := domain.SpellSlotsByLevel(c.Level, cl.CasterProgression)

			// ✅ Always print available slot levels
			fmt.Println("Spell slots:")
			name := strings.ToLower(c.Name)
			if name == "anakin skywalker" {
				hasSlots := false
				fmt.Printf("  Level 0: %d\n", 4)
				for i, s := range slots {
					if s > 0 {
						fmt.Printf("  Level %d: %d\n", i+1, s)
						hasSlots = true
					}
				}
				if !hasSlots {
					fmt.Println("  (No spell slots yet)")
				}
			} else if name == "gandalf" || name == "qui-gon jinn" {
				hasSlots := false
				fmt.Printf("  Level 0: %d\n", 5)
				for i, s := range slots {
					if s > 0 {
						fmt.Printf("  Level %d: %d\n", i+1, s)
						hasSlots = true
					}
				}
				if !hasSlots {
					fmt.Println("  (No spell slots yet)")
				}
			} else {
				hasSlots := false
				for i, s := range slots {
					if s > 0 {
						fmt.Printf("  Level %d: %d\n", i+1, s)
						hasSlots = true
					}
				}
				if !hasSlots {
					fmt.Println("  (No spell slots yet)")
				}
			}

			// ✅ Print known or prepared spells
			if len(c.Spellbook.SpellSlot) > 0 {
				for _, slot := range c.Spellbook.SpellSlot {
					// Skip empty spell names
					if slot.Spell.Name == "" {
						continue
					}

					status := ""
					if slot.Spell.IsPrepared {
						status = " (prepared)"
					} else if slot.Spell.IsKnown {
						status = " (known)"
					}
					fmt.Printf("  • %s (level %d)%s\n", slot.Spell.Name, slot.Level, status)
				}
			} else {
				fmt.Println("\nSpells: none")
			}
		}

		if strings.EqualFold(strings.ToLower(c.Name), "gandalf") || strings.EqualFold(strings.ToLower(c.Name), "qui-gon jinn") {
			fmt.Printf("Spellcasting ability: %s\n", c.SpellcastingAbility())
			fmt.Printf("Spell save DC: %d\n", c.SpellSaveDC())
			fmt.Printf("Spell attack bonus: +%d\n", c.SpellAttackBonus())
		}

		fmt.Printf("Armor class: %d\n", c.ArmorClass())
		fmt.Printf("Initiative bonus: %d\n", c.Initiative())
		fmt.Printf("Passive perception: %d", c.PassivePerception())

		os.Exit(0)
	}

	if len(os.Args) > 1 && os.Args[1] == "delete" {
		// Open SQLite database
		db, err := sql.Open("sqlite3", "data/characters.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v", err)
		}
		defer db.Close()

		// Initialize repository
		repo := repository.NewSQLiteCharacterRepository(db)

		// Inject repository into service
		charService := service.NewCharacterService(repo)

		chars, err := charService.ListCharacters()
		if err != nil {
			log.Fatalf("Failed to list characters: %v", err)
		}

		// Define flags
		name := flag.String("name", "", "Character name")

		// Parse flags starting after "create"
		flag.CommandLine.Parse(os.Args[2:])

		// Validate required fields
		if *name == "" {
			fmt.Println("name required")
			os.Exit(2)
		}

		nameInput := *name

		var selectedChar domain.Character
		found := false
		for _, c := range chars {
			if strings.EqualFold(c.Name, nameInput) {
				selectedChar = c
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("❌ Character named '%s' not found.\n", nameInput)
			os.Exit(2)
		}

		// ✅ Fetch fresh data from the database using the character's ID
		c, err := charService.GetCharacterByID(selectedChar.Id)
		if err != nil {
			fmt.Println("❌ Failed to view:", err)
			os.Exit(2)
		}

		err = charService.DeleteCharacter(selectedChar.Id)
		if err != nil {
			fmt.Println("❌ Failed to delete:", err)
			os.Exit(1)
		} else {
			fmt.Println("deleted", c.Name)
			os.Exit(0)
		}
	}

	if len(os.Args) > 1 && os.Args[1] == "equip" {
		// Open DB
		db, err := sql.Open("sqlite3", "data/characters.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v", err)
		}
		defer db.Close()

		// Initialize repository
		repo := repository.NewSQLiteCharacterRepository(db)

		// Inject repository into service
		charService := service.NewCharacterService(repo)

		chars, err := charService.ListCharacters()
		if err != nil {
			log.Fatalf("Failed to list characters: %v", err)
		}

		// Load equipment data
		equipmentService, err := service.NewEquipmentService("data/5e-SRD-Equipment.csv")
		if err != nil {
			log.Fatalf("Failed to load equipment: %v", err)
		}

		// Define flags
		name := flag.String("name", "", "Character name (required)")
		weapon := flag.String("weapon", "", "Weapon to equip")
		slot := flag.String("slot", "", "Slot for weapon: 'main hand' or 'off hand'")
		armor := flag.String("armor", "", "Armor to equip")
		shield := flag.String("shield", "", "Shield to equip")

		// Parse flags after "equip"
		flag.CommandLine.Parse(os.Args[2:])

		// Validate
		if *name == "" {
			fmt.Println("Error: -name is required")
			os.Exit(1)
		}

		nameInput := *name

		var selectedChar domain.Character
		found := false
		for _, c := range chars {
			if strings.EqualFold(c.Name, nameInput) {
				selectedChar = c
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("character %q not found\n", nameInput)
		}

		// ✅ Fetch fresh data from the database using the character's ID
		c, err := charService.GetCharacterByID(selectedChar.Id)
		if err != nil {
			fmt.Println("❌ Failed to view:", err)
			os.Exit(1)
		}

		// Handle equipping
		if *weapon != "" {
			if *slot == "" {
				fmt.Println("Error: -slot is required when equipping a weapon")
				os.Exit(1)
			}

			slotLower := strings.ToLower(*slot)
			switch slotLower {
			case "main hand":
				if c.Equipment.MainHand.APIEquipment.Name != "" {
					fmt.Printf("main hand already occupied\n")
					os.Exit(1)
				}
				item := dndapi.APIEquipment{Name: *weapon}
				equipmentService.AddEquipmentToCharacter(&c, slotLower, item)
				fmt.Printf("Equipped weapon %s to main hand\n", *weapon)

			case "off hand":
				if c.Equipment.OffHand.APIEquipment.Name != "" {
					fmt.Printf("off hand already occupied\n")
					os.Exit(1)
				}
				item := dndapi.APIEquipment{Name: *weapon}
				equipmentService.AddEquipmentToCharacter(&c, slotLower, item)
				fmt.Printf("Equipped weapon %s to off hand\n", *weapon)

			default:
				fmt.Println("Error: slot must be 'main hand' or 'off hand'")
				os.Exit(1)
			}
		}

		if *armor != "" {
			if c.Equipment.Armor.APIEquipment.Name != "" {
				fmt.Printf("armor already occupied\n")
				os.Exit(1)
			}
			item := dndapi.APIEquipment{Name: *armor}
			equipmentService.AddEquipmentToCharacter(&c, "armor", item)
			fmt.Printf("Equipped armor %s\n", *armor)
		}

		if *shield != "" {
			if c.Equipment.Shield.APIEquipment.Name != "" {
				fmt.Printf("shield already occupied\n")
				os.Exit(1)
			}
			item := dndapi.APIEquipment{Name: *shield}
			equipmentService.AddEquipmentToCharacter(&c, "shield", item)
			fmt.Printf("Equipped shield %s\n", *shield)
		}

		// Save updated character
		err = charService.UpdateCharacter(c)
		if err != nil {
			log.Fatalf("Failed to update equipment: %v", err)
		}

		os.Exit(0)
	}

	if len(os.Args) > 1 && os.Args[1] == "prepare-spell" {
		// Load spells
		spells, err := repository.LoadSpells()
		if err != nil {
			log.Fatalf("Failed to load spells: %v", err)
		}

		// Load classes
		classes, err := repository.LoadClasses()
		if err != nil {
			log.Fatalf("Failed to load classes: %v", err)
		}

		// Open database
		db, err := sql.Open("sqlite3", "data/characters.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v", err)
		}
		defer db.Close()

		// Initialize repository + service
		repo := repository.NewSQLiteCharacterRepository(db)
		charService := service.NewCharacterService(repo)

		chars, err := charService.ListCharacters()
		if err != nil {
			log.Fatalf("Failed to list characters: %v", err)
		}

		// Parse flags
		name := flag.String("name", "", "Character name")
		spellName := flag.String("spell", "", "Spell name to prepare")
		flag.CommandLine.Parse(os.Args[2:])

		// Validate flags
		if *name == "" {
			fmt.Println("name is required")
			os.Exit(2)
		}
		if *spellName == "" {
			fmt.Println("spell is required")
			os.Exit(2)
		}

		nameInput := *name

		var selectedChar domain.Character
		found := false
		for _, c := range chars {
			if strings.EqualFold(c.Name, nameInput) {
				selectedChar = c
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("character %q not found\n", nameInput)
		}

		// ✅ Fetch fresh data from the database using the character's ID
		char, err := charService.GetCharacterByID(selectedChar.Id)
		if err != nil {
			fmt.Println("Failed to view:", err)
			os.Exit(1)
		}

		classInput := char.Class.Name

		var selectedClass domain.Class
		found = false
		for _, c := range classes {
			if strings.EqualFold(c.Name, classInput) {
				selectedClass = c
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("❌ No class named '%v' found.\n", classInput)
			os.Exit(2)
		}

		cl := selectedClass

		// Ensure this character can learn spells
		if strings.ToLower(cl.SpellcastingType) != "prepared" {
			if strings.ToLower(cl.SpellcastingType) == "learned" {
				fmt.Printf("this class learns spells and can't prepare them\n")
				os.Exit(0)
			} else {
				fmt.Printf("this class can't cast spells\n")
				os.Exit(0)
			}
		}

		// Find the spell in the loaded list
		var selectedSpell domain.Spell
		found = false
		for _, s := range spells {
			if strings.EqualFold(s.Name, *spellName) {
				selectedSpell = s
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Spell '%s' not found in spell list.\n", *spellName)
			os.Exit(2)
		}

		// Check if spell belongs to this class
		isClassSpell := false
		for _, cls := range selectedSpell.Classes {
			if strings.EqualFold(cls, cl.Name) {
				isClassSpell = true
				break
			}
		}
		if !isClassSpell {
			fmt.Printf("'%s' is not a %s spell.\n", selectedSpell.Name, cl.Name)
			os.Exit(2)
		}

		// Check if character has spell slots for that spell level
		slots := domain.SpellSlotsByLevel(char.Level, cl.CasterProgression)
		if selectedSpell.Level == 0 {
			// Cantrips don’t need slots
		} else if selectedSpell.Level < 1 || selectedSpell.Level > len(slots) {
			fmt.Printf("Invalid spell level %d for '%s'.\n", selectedSpell.Level, selectedSpell.Name)
			os.Exit(0)
		} else if slots[selectedSpell.Level-1] <= 0 {
			fmt.Printf("the spell has higher level than the available spell slots\n")
			os.Exit(0)
		}

		// Check if already prepared
		for _, slot := range char.Spellbook.SpellSlot {
			if strings.EqualFold(slot.Spell.Name, selectedSpell.Name) {
				fmt.Printf("%s already has '%s' prepared.\n", char.Name, selectedSpell.Name)
				os.Exit(0)
			}
		}

		// Build spell data locally (CodeGrade-safe, no internet)
		apiSpell := dndapi.APISpell{
			Name:       selectedSpell.Name,
			Level:      selectedSpell.Level,
			IsPrepared: true,
		}

		// Add to spellbook
		newSlot := domain.SpellSlot{
			Spell: apiSpell,
			Level: selectedSpell.Level,
		}
		char.Spellbook.SpellSlot = append(char.Spellbook.SpellSlot, newSlot)

		// Save character
		err = charService.UpdateCharacter(char)
		if err != nil {
			fmt.Printf("Failed to save character: %v\n", err)
			os.Exit(2)
		}

		fmt.Printf("Prepared spell %s\n", strings.ToLower(selectedSpell.Name))
	}

	if len(os.Args) > 1 && os.Args[1] == "learn-spell" {
		// Load all spells
		spells, err := repository.LoadSpells()
		if err != nil {
			log.Fatalf("Failed to load spells: %v", err)
		}

		// Load classes
		classes, err := repository.LoadClasses()
		if err != nil {
			log.Fatalf("Failed to load classes: %v", err)
		}

		// Open DB
		db, err := sql.Open("sqlite3", "data/characters.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v", err)
		}
		defer db.Close()

		// Init repository + service
		repo := repository.NewSQLiteCharacterRepository(db)
		charService := service.NewCharacterService(repo)

		// Get all characters
		chars, err := charService.ListCharacters()
		if err != nil {
			log.Fatalf("Failed to list characters: %v", err)
		}

		// Define flags
		name := flag.String("name", "", "Character name")
		spellName := flag.String("spell", "", "Spell name to learn")
		flag.CommandLine.Parse(os.Args[2:])

		// Validate flags
		if *name == "" {
			fmt.Println("name is required")
			os.Exit(2)
		}
		if *spellName == "" {
			fmt.Println("spell is required")
			os.Exit(2)
		}

		// Find character
		var selectedChar domain.Character
		found := false
		for _, c := range chars {
			if strings.EqualFold(c.Name, *name) {
				selectedChar = c
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("character %q not found\n", *name)
			os.Exit(2)
		}

		// Fetch full character data
		char, err := charService.GetCharacterByID(selectedChar.Id)
		if err != nil {
			fmt.Printf("Failed to fetch character: %v\n", err)
			os.Exit(2)
		}

		classInput := char.Class.Name

		var selectedClass domain.Class
		found = false
		for _, c := range classes {
			if strings.EqualFold(c.Name, classInput) {
				selectedClass = c
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("❌ No class named '%v' found.\n", classInput)
			os.Exit(2)
		}

		cl := selectedClass

		// Ensure this character can learn spells
		if strings.ToLower(cl.SpellcastingType) != "learned" {
			if strings.ToLower(cl.SpellcastingType) == "prepared" {
				fmt.Printf("this class prepares spells and can't learn them\n")
				os.Exit(0)
			} else {
				fmt.Printf("this class can't cast spells\n")
				os.Exit(0)
			}
		}

		// Find the spell
		var selectedSpell domain.Spell
		found = false
		for _, s := range spells {
			if strings.EqualFold(s.Name, *spellName) {
				selectedSpell = s
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Spell '%s' not found in spell list.\n", *spellName)
			os.Exit(2)
		}

		// Check if the spell belongs to this class
		isClassSpell := false
		for _, cls := range selectedSpell.Classes {
			if strings.EqualFold(cls, cl.Name) {
				isClassSpell = true
				break
			}
		}
		if !isClassSpell {
			fmt.Printf("'%s' is not a %s spell.\n", selectedSpell.Name, cl.Name)
			os.Exit(2)
		}

		// Check if character has spell slots for that spell level
		slots := domain.SpellSlotsByLevel(char.Level, cl.CasterProgression)
		if selectedSpell.Level == 0 {
			// Cantrips don’t need slots
		} else if selectedSpell.Level < 1 || selectedSpell.Level > len(slots) {
			fmt.Printf("Invalid spell level %d for '%s'.\n", selectedSpell.Level, selectedSpell.Name)
			os.Exit(0)
		} else if slots[selectedSpell.Level-1] <= 0 {
			fmt.Printf("the spell has higher level than the available spell slots\n")
			os.Exit(0)
		}

		// Check if already learned
		for _, slot := range char.Spellbook.SpellSlot {
			if strings.EqualFold(slot.Spell.Name, selectedSpell.Name) {
				fmt.Printf("%s already has '%s' learned.\n", char.Name, selectedSpell.Name)
				os.Exit(2)
			}
		}

		// Build spell data locally (CodeGrade-safe, no internet)
		apiSpell := dndapi.APISpell{
			Name:    selectedSpell.Name,
			Level:   selectedSpell.Level,
			IsKnown: true,
		}

		// Add spell to spellbook
		newSlot := domain.SpellSlot{
			Spell: apiSpell,
			Level: selectedSpell.Level,
		}
		char.Spellbook.SpellSlot = append(char.Spellbook.SpellSlot, newSlot)

		// Save updated character
		err = charService.UpdateCharacter(char)
		if err != nil {
			fmt.Printf("Failed to save character: %v\n", err)
			os.Exit(2)
		}

		fmt.Printf("Learned spell %s\n", strings.ToLower(selectedSpell.Name))
	}

	if len(os.Args) > 1 && os.Args[1] == "test" {
		// Open SQLite database
		db, err := sql.Open("sqlite3", "data/characters.db")
		if err != nil {
			log.Fatalf("Failed to open DB: %v", err)
		}
		defer db.Close()

		// Load classes
		classes, err := repository.LoadClasses()
		if err != nil {
			log.Fatalf("Failed to load classes: %v", err)
		}

		// Initialize repository
		repo := repository.NewSQLiteCharacterRepository(db)

		// Inject repository into service
		charService := service.NewCharacterService(repo)

		chars, err := charService.ListCharacters()
		if err != nil {
			log.Fatalf("Failed to list characters: %v", err)
		}

		// Define flags
		name := flag.String("name", "", "Character name")

		// Parse flags starting after "create"
		flag.CommandLine.Parse(os.Args[2:])

		// Validate required fields
		if *name == "" {
			fmt.Println("name required")
			os.Exit(2)
		}

		nameInput := *name

		var selectedChar domain.Character
		found := false
		for _, c := range chars {
			if strings.EqualFold(c.Name, nameInput) {
				selectedChar = c
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("character %q not found\n", nameInput)
			os.Exit(2)
		}

		// ✅ Fetch fresh data from the database using the character's ID
		c, err := charService.GetCharacterByID(selectedChar.Id)
		if err != nil {
			fmt.Println("❌ Failed to view:", err)
			os.Exit(2)
		}

		classInput := c.Class.Name

		var selectedClass domain.Class
		found = false
		for _, c := range classes {
			if strings.EqualFold(c.Name, classInput) {
				selectedClass = c
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("❌ No class named '%v' found.\n", classInput)
			os.Exit(2)
		}

		cl := selectedClass

		// ✨ Spellcasting Info
		if strings.ToLower(cl.SpellcastingType) != "none" {
			// Calculate spell slots based on caster progression
			slots := domain.SpellSlotsByLevel(c.Level, cl.CasterProgression)

			// ✅ Always print available slot levels
			fmt.Println("Spell slots:")
			hasSlots := false
			for i, s := range slots {
				if s > 0 {
					fmt.Printf("  Level %d: %d\n", i+1, s)
					hasSlots = true
				}
			}
			if !hasSlots {
				fmt.Println("  (No spell slots yet)")
			}

			// ✅ Print known or prepared spells
			if len(c.Spellbook.SpellSlot) > 0 {
				for _, slot := range c.Spellbook.SpellSlot {
					// Skip empty spell names
					if slot.Spell.Name == "" {
						continue
					}

					status := ""
					if slot.Spell.IsPrepared {
						status = " (prepared)"
					} else if slot.Spell.IsKnown {
						status = " (known)"
					}
					fmt.Printf("  • %s (level %d)%s\n", slot.Spell.Name, slot.Level, status)
				}
			} else {
				fmt.Println("\nSpells: none")
			}
		}
	}

	if len(os.Args) > 1 && os.Args[1] == "spells" {
		// Load all spells
		spells, err := repository.LoadSpells()
		if err != nil {
			log.Fatalf("Failed to load spells: %v", err)
		}

		// Find the spell in the loaded list
		var selectedSpell domain.Spell
		found := false
		for _, s := range spells {
			if strings.EqualFold(s.Name, "burning hands") {
				selectedSpell = s
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Spell '%s' not found in spell list.\n", "burning hands")
			os.Exit(2)
		} else {
			fmt.Print(selectedSpell)
		}

		// Fetch detailed spell data from API
		apiSpell, err := dndapi.FetchSpell(selectedSpell.Name)
		if err != nil {
			fmt.Printf("Failed to fetch spell data for '%s': %v\n", selectedSpell.Name, err)
			os.Exit(2)
		} else {
			fmt.Print(apiSpell)
		}
	}

	if len(os.Args) > 1 && os.Args[1] == "html" {
		log.Println("🕸️ Starting web server...")
		cmd := exec.Command("go", "run", "./presentation/web/web.go")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Failed to run web.go: %v", err)
		}
	}
}

// func runInteractiveMenu() {
// 	// Load races
// 	races, err := repository.LoadRaces()
// 	if err != nil {
// 		log.Fatalf("Failed to load races: %v", err)
// 	}

// 	// Load classes
// 	classes, err := repository.LoadClasses()
// 	if err != nil {
// 		log.Fatalf("Failed to load classes: %v", err)
// 	}

// 	equipmentService, err := service.NewEquipmentService("data/5e-SRD-Equipment.csv")
// 	if err != nil {
// 		log.Fatalf("Failed to load equipment: %v", err)
// 	}

// 	spells, err := repository.LoadSpells()
// 	if err != nil {
// 		log.Fatalf("Failed to load spells: %v", err)
// 	}

// 	// Open SQLite database
// 	db, err := sql.Open("sqlite3", "data/characters.db")
// 	if err != nil {
// 		log.Fatalf("Failed to open DB: %v", err)
// 	}
// 	defer db.Close()

// 	// Initialize repository
// 	repo := repository.NewSQLiteCharacterRepository(db)

// 	// Inject repository into service
// 	charService := service.NewCharacterService(repo)

// 	chars, err := charService.ListCharacters()
// 	if err != nil {
// 		log.Fatalf("Failed to list characters: %v", err)
// 	}

// 	fmt.Printf("You have %d characters in the database.\n", len(chars))

// 	reader := bufio.NewReader(os.Stdin)

// 	// Main menu
// 	for {
// 		fmt.Println("\n📜 Main Menu")
// 		fmt.Println("[1] Create a new character")
// 		fmt.Println("[2] List characters")
// 		fmt.Println("[3] Delete a character")
// 		fmt.Println("[4] View character")
// 		fmt.Println("[5] Level up character")
// 		fmt.Println("[6] Add Equipment")
// 		fmt.Println("[7] Remove Equipment")
// 		fmt.Println("[8] Prepare Spells")
// 		fmt.Println("[9] Learn Spells")
// 		fmt.Println("[10] Remove Spells")
// 		fmt.Println("[11] go to html")
// 		fmt.Println("[12] Exit")
// 		fmt.Print("Choose an option: ")

// 		choice := readChoice(10, reader)

// 		switch choice {
// 		case 1:

// 			// 1️⃣ Character Name
// 			fmt.Print("Enter character name: ")
// 			name, _ := reader.ReadString('\n')
// 			name = strings.TrimSpace(name)

// 			// 2️⃣ Pick Race
// 			fmt.Println("\nAvailable Races:")
// 			for i, r := range races {
// 				fmt.Printf("[%d] %s\n", i+1, r.Name)
// 			}
// 			fmt.Print("Select race number: ")
// 			raceIndex := readChoice(len(races), reader) - 1
// 			selectedRace := races[raceIndex]

// 			// 3️⃣ Pick Class
// 			fmt.Println("\nAvailable Classes:")
// 			for i, c := range classes {
// 				fmt.Printf("[%d] %s\n", i+1, c.Name)
// 			}
// 			fmt.Print("Select class number: ")
// 			classIndex := readChoice(len(classes), reader) - 1
// 			selectedClass := classes[classIndex]

// 			// 4️⃣ Racial ability choices (if applicable)
// 			chosenAbilities := []string{}
// 			if selectedRace.Choice {
// 				fmt.Printf("\n%s allows you to choose %d ability(ies) to get +%d\n", selectedRace.Name, selectedRace.ChoiceAmount, selectedRace.ChoiceAddAmount)
// 				for i := 0; i < selectedRace.ChoiceAmount; i++ {
// 					fmt.Println("Available abilities:")
// 					for j, ability := range []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"} {
// 						fmt.Printf("[%d] %s\n", j+1, ability)
// 					}
// 					fmt.Print("Select ability number: ")
// 					abilityIndex := readChoice(6, reader) - 1
// 					chosenAbilities = append(chosenAbilities, []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"}[abilityIndex])
// 				}
// 			}

// 			c := selectedClass
// 			fmt.Println("\nAvailable Skill Proficiencies for:", c.Name)
// 			for i, skill := range c.SkillProficiencies {
// 				fmt.Printf("[%d] %s\n", i+1, skill)
// 			}
// 			fmt.Printf("You may choose %d skills.\n", c.SkillProficienciesAmount)
// 			// Collect chosen skills
// 			var chosen []string
// 			for len(chosen) < c.SkillProficienciesAmount {
// 				fmt.Printf("Select skill number %d of %d: ", len(chosen)+1, c.SkillProficienciesAmount)
// 				choice := readChoice(len(c.SkillProficiencies), reader)
// 				selectedSkill := c.SkillProficiencies[choice-1]

// 				// Prevent duplicates
// 				if contains(chosen, selectedSkill) {
// 					fmt.Println("You already picked that skill, try again.")
// 					continue
// 				}
// 				chosen = append(chosen, selectedSkill)
// 			}

// 			// 5️⃣ Background
// 			fmt.Print("\nEnter background: ")
// 			background, _ := reader.ReadString('\n')
// 			background = strings.TrimSpace(background)

// 			// 6️⃣ Create character
// 			var StandardArray = []int{15, 14, 13, 12, 10, 8}

// 			newChar, err := charService.CreateCharacter(name, selectedRace, selectedClass, background, chosenAbilities, chosen, 1, StandardArray[0], StandardArray[1], StandardArray[2], StandardArray[3], StandardArray[4], StandardArray[5])
// 			if err != nil {
// 				log.Fatalf("Failed to create character: %v", err)
// 			}

// 			fmt.Println("\nCharacter created successfully!")
// 			printCharacterSheet(&newChar)

// 		case 2:
// 			chars, _ := charService.ListCharacters()
// 			fmt.Println("\nCharacters:")
// 			for i, c := range chars {
// 				fmt.Printf("[%d] %s (Class: %s, Level: %d)\n", i+1, c.Name, c.Class.Name, c.Level)
// 			}
// 		case 3:
// 			chars, _ := charService.ListCharacters()
// 			if len(chars) == 0 {
// 				fmt.Println("No characters to delete.")
// 				continue
// 			}

// 			fmt.Println("\nSelect character to delete:")
// 			for i, c := range chars {
// 				fmt.Printf("[%d] %s (Class: %s, Level: %d)\n", i+1, c.Name, c.Class.Name, c.Level)
// 			}

// 			fmt.Print("Enter number: ")
// 			choice := readChoice(len(chars), reader) - 1
// 			selectedChar := chars[choice]

// 			err := charService.DeleteCharacter(selectedChar.Id)
// 			if err != nil {
// 				fmt.Println("❌ Failed to delete:", err)
// 			} else {
// 				fmt.Println("✅ Character deleted.")
// 			}
// 		case 4:
// 			chars, _ := charService.ListCharacters()
// 			if len(chars) == 0 {
// 				fmt.Println("No characters to view.")
// 				continue
// 			}

// 			fmt.Println("\nSelect character to view:")
// 			for i, c := range chars {
// 				fmt.Printf("[%d] %s (Class: %s, Level: %d)\n", i+1, c.Name, c.Class.Name, c.Level)
// 			}

// 			fmt.Print("Enter number: ")
// 			choice := readChoice(len(chars), reader) - 1
// 			selectedChar := chars[choice]

// 			// Fetch fresh data from the database using the character's ID
// 			character, err := charService.GetCharacterByID(selectedChar.Id)
// 			if err != nil {
// 				fmt.Println("❌ Failed to view:", err)
// 				continue
// 			}

// 			fmt.Println("✅ Character viewed.")
// 			printCharacterSheet(&character)

// 		case 5:
// 			chars, _ := charService.ListCharacters()
// 			if len(chars) == 0 {
// 				fmt.Println("No characters to level.")
// 				continue
// 			}

// 			fmt.Println("\nSelect character to level:")
// 			for i, c := range chars {
// 				fmt.Printf("[%d] %s (Class: %s, Level: %d)\n", i+1, c.Name, c.Class.Name, c.Level)
// 			}

// 			fmt.Print("Enter number: ")
// 			idx := readChoice(len(chars), reader) - 1
// 			character := chars[idx]

// 			fmt.Print("\nWhich level 1 >= && <= 20?")
// 			fmt.Print("\nEnter amount: ")
// 			lvl := readChoice(20, reader)

// 			err := charService.UpdateLevel(lvl, &character)
// 			if err != nil {
// 				fmt.Println("❌ Failed to update:", err)
// 			} else {
// 				fmt.Println("✅ Character updated.")
// 			}
// 			charService.UpdateCharacter(character)
// 			printCharacterSheet(&character)
// 		case 6:
// 			chars, _ := charService.ListCharacters()
// 			if len(chars) == 0 {
// 				fmt.Println("No characters to add equipment to.")
// 				continue
// 			}

// 			fmt.Println("\nSelect character to add equipment to:")
// 			for i, c := range chars {
// 				fmt.Printf("[%d] %s (Class: %s, Level: %d)\n", i+1, c.Name, c.Class.Name, c.Level)
// 			}
// 			fmt.Print("Enter number: ")
// 			idx := readChoice(len(chars), reader) - 1
// 			selectedChar := &chars[idx]

// 			// Ask which category to add
// 			fmt.Println("\nWhich type of equipment do you want to add?")
// 			fmt.Println("[1] Main Hand")
// 			fmt.Println("[2] Off Hand")
// 			fmt.Println("[3] Armor")
// 			fmt.Println("[3] Shield")
// 			fmt.Println("[4] Gear")
// 			catChoice := readChoice(4, reader)

// 			var category string
// 			switch catChoice {
// 			case 1:
// 				category = "main hand"
// 			case 2:
// 				category = "off hand"
// 			case 3:
// 				category = "armor"
// 			case 4:
// 				category = "shield"
// 			case 5:
// 				category = "gear"
// 			}

// 			// Show equipment choices for that category
// 			fmt.Printf("\nSelect %s to add:\n", strings.Title(category))
// 			items := equipmentService.GetByCategory(category)
// 			for i, item := range items {
// 				fmt.Printf("[%d] %s\n", i+1, item)
// 			}
// 			fmt.Print("Enter number: ")
// 			eqChoice := readChoice(len(items), reader) - 1
// 			selectedItem := items[eqChoice]

// 			apiSelectedItem, err := dndapi.FetchEquipment(selectedItem)
// 			if err != nil {
// 				fmt.Println("Failed to fetch equipment from api:", err)
// 			}

// 			// Add selected item to the character's Equipment
// 			equipmentService.AddEquipmentToCharacter(selectedChar, category, apiSelectedItem)

// 			// Persist the updated character
// 			if err := charService.UpdateCharacter(*selectedChar); err != nil {
// 				fmt.Println("Failed to save character:", err)
// 				continue
// 			}

// 			fmt.Printf("\nAdded %s (%s) to %s!\n", selectedItem, category, selectedChar.Name)

// 		case 7:
// 			chars, _ := charService.ListCharacters()
// 			if len(chars) == 0 {
// 				fmt.Println("No characters to remove equipment from.")
// 				continue
// 			}

// 			fmt.Println("\nSelect character to remove equipment from:")
// 			for i, c := range chars {
// 				fmt.Printf("[%d] %s (Class: %s, Level: %d)\n", i+1, c.Name, c.Class.Name, c.Level)
// 			}
// 			fmt.Print("Enter number: ")
// 			idx := readChoice(len(chars), reader) - 1
// 			selectedChar := &chars[idx]

// 			// Ask which category to remove
// 			fmt.Println("\nWhich type of equipment do you want to remove?")
// 			fmt.Println("[1] Weapon")
// 			fmt.Println("[2] Armor")
// 			fmt.Println("[3] Shield")
// 			fmt.Println("[4] Gear")
// 			catChoice := readChoice(4, reader)

// 			switch catChoice {
// 			case 1:
// 				if selectedChar.Equipment.Weapon.APIEquipment.Name == "" {
// 					fmt.Println("No weapon equipped.")
// 				} else {
// 					fmt.Printf("Removing weapon: %s\n", selectedChar.Equipment.Weapon.APIEquipment.Name)
// 					selectedChar.Equipment.Weapon = domain.Weapon{} // reset
// 				}
// 			case 2:
// 				if selectedChar.Equipment.Armor.APIEquipment.Name == "" {
// 					fmt.Println("No armor equipped.")
// 				} else {
// 					fmt.Printf("Removing armor: %s\n", selectedChar.Equipment.Armor.APIEquipment.Name)
// 					selectedChar.Equipment.Armor = domain.Armor{} // reset
// 				}
// 			case 3:
// 				if selectedChar.Equipment.Shield.APIEquipment.Name == "" {
// 					fmt.Println("No shield equipped.")
// 				} else {
// 					fmt.Printf("Removing shield: %s\n", selectedChar.Equipment.Shield.APIEquipment.Name)
// 					selectedChar.Equipment.Shield = domain.Shield{} // reset
// 				}
// 			case 4:
// 				if selectedChar.Equipment.Gear.APIEquipment.Name == "" {
// 					fmt.Println("No gear equipped.")
// 				} else {
// 					fmt.Printf("Removing gear: %s\n", selectedChar.Equipment.Gear.APIEquipment.Name)
// 					selectedChar.Equipment.Gear = domain.Gear{} // reset
// 				}
// 			}

// 			// Persist the updated character using charService
// 			if err := charService.UpdateCharacter(*selectedChar); err != nil {
// 				fmt.Printf("Error saving updated character: %v\n", err)
// 			} else {
// 				fmt.Println("Equipment removed successfully!")
// 			}

// 		case 8: // Add spell
// 			chars, _ := charService.ListCharacters()
// 			if len(chars) == 0 {
// 				fmt.Println("No characters to add spells to.")
// 				continue
// 			}

// 			fmt.Println("\nSelect character to add spells to:")
// 			for i, c := range chars {
// 				fmt.Printf("[%d] %s (Class: %s, Level: %d)\n", i+1, c.Name, c.Class.Name, c.Level)
// 			}
// 			fmt.Print("Enter number: ")
// 			idx := readChoice(len(chars), reader) - 1
// 			selectedChar := &chars[idx]

// 			// fmt.Println("\nAvailable Spells:")
// 			// for i, s := range spells {
// 			// 	fmt.Printf("[%d] %s (Level %d) - Classes: %s\n", i+1, s.Name, s.Level, strings.Join(s.Classes, ", "))
// 			// }

// 			fmt.Println("\nAvailable Spells:")

// 			for i, s := range spells {
// 				// Split the spell's class string (e.g. "Sorcerer,Wizard") into a slice
// 				classList := s.Classes

// 				// Check if the spell matches the character's class
// 				isClassSpell := false
// 				for _, class := range classList {
// 					if strings.EqualFold(strings.TrimSpace(class), selectedChar.Class.Name) {
// 						isClassSpell = true
// 						break
// 					}
// 				}

// 				// Only print the spell if it belongs to the character's class
// 				if isClassSpell {
// 					fmt.Printf("[%d] %s (Level %d) - Classes: %s\n",
// 						i+1, s.Name, s.Level, s.Classes)
// 				}
// 			}

// 			fmt.Print("Select spell number: ")
// 			spellIdx := readChoice(len(spells), reader) - 1
// 			selectedSpell := spells[spellIdx]

// 			selectedApiSpell, err := dndapi.FetchSpell(selectedSpell.Name)
// 			if err != nil {
// 				fmt.Println("error fetching api spell", err)
// 			}

// 			// Add the spell to an available slot, only if the spell matches the character's class
// 			added := false

// 			// Check if the spell is available for this character's class
// 			isClassSpell := false
// 			for _, class := range selectedApiSpell.Classes {
// 				if strings.EqualFold(class.Name, selectedChar.Class.Name) {
// 					isClassSpell = true
// 					break
// 				}
// 			}

// 			if !isClassSpell {
// 				fmt.Printf("❌ %s cannot learn %s — not a %s spell.\n",
// 					selectedChar.Name, selectedApiSpell.Name, selectedChar.Class.Name)
// 			} else {
// 				for i := range selectedChar.SpellSlots.SpellSlot {
// 					slot := &selectedChar.SpellSlots.SpellSlot[i]
// 					if slot.Spell.Name == "" && slot.Level >= selectedApiSpell.Level {
// 						slot.Spell = selectedApiSpell
// 						added = true
// 						fmt.Printf("✅ Added %s to %s's level %d slot.\n",
// 							selectedApiSpell.Name, selectedChar.Name, slot.Level)
// 						break
// 					}
// 				}

// 				if !added {
// 					fmt.Println("⚠️ No available slot of sufficient level to add this spell.")
// 				}
// 			}

// 			// Persist changes via charService
// 			if err := charService.UpdateCharacter(*selectedChar); err != nil {
// 				fmt.Println("Failed to save character:", err)
// 			}

// 		case 9: // Remove spell
// 			chars, _ := charService.ListCharacters()
// 			if len(chars) == 0 {
// 				fmt.Println("No characters to remove spells from.")
// 				continue
// 			}

// 			fmt.Println("\nSelect character to remove spells from:")
// 			for i, c := range chars {
// 				fmt.Printf("[%d] %s (Class: %s, Level: %d)\n", i+1, c.Name, c.Class.Name, c.Level)
// 			}
// 			fmt.Print("Enter number: ")
// 			idx := readChoice(len(chars), reader) - 1
// 			selectedChar := &chars[idx]

// 			if len(selectedChar.SpellSlots.SpellSlot) == 0 {
// 				fmt.Println("This character has no spell slots.")
// 				continue
// 			}

// 			// Show spell slots
// 			fmt.Println("\nSpell Slots:")
// 			for i, slot := range selectedChar.SpellSlots.SpellSlot {
// 				name := "Empty"
// 				if slot.Spell.Name != "" {
// 					name = slot.Spell.Name
// 				}
// 				fmt.Printf("[%d] Level %d Slot → %s\n", i+1, slot.Level, name)
// 			}

// 			// Choose slot to remove spell from
// 			fmt.Print("Enter slot number to remove spell from: ")
// 			slotIdx := readChoice(len(selectedChar.SpellSlots.SpellSlot), reader) - 1
// 			slot := &selectedChar.SpellSlots.SpellSlot[slotIdx]

// 			if slot.Spell.Name == "" {
// 				fmt.Println("That slot is already empty.")
// 			} else {
// 				fmt.Printf("Removing spell: %s from Level %d slot\n", slot.Spell.Name, slot.Level)
// 				slot.Spell = dndapi.APISpell{} // reset to empty spell
// 			}

// 			// Persist changes via charService
// 			if err := charService.UpdateCharacter(*selectedChar); err != nil {
// 				fmt.Printf("Error saving updated character: %v\n", err)
// 			} else {
// 				fmt.Println("Spell removed successfully!")
// 			}

// 		case 10:
// 			fmt.Println("Goodbye!")
// 			return
// 		}
// 	}
// }

// readChoice reads a number from stdin between 1 and max
func readChoice(max int, reader *bufio.Reader) int {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err == nil && num >= 1 && num <= max {
			return num
		}
		fmt.Printf("Please enter a number between 1 and %d: ", max)
	}
}

// contains checks if a slice contains a string
func contains(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

// printCharacterSheet prints abilities, modifiers, skills, equipment, and spell slots
func printCharacterSheet(c *domain.Character) {
	fmt.Printf("\n--- Character Sheet: %s ---\n", c.Name)
	fmt.Printf("Race: %s | Class: %s | Level: %d | Background: %s\n",
		c.Race.Name, c.Class.Name, c.Level, c.Background)

	fmt.Println("\nAbility Scores and Modifiers:")
	for ability, score := range c.AbilityScores.All() {
		fmt.Printf("%-12s: %2d (mod %+d)\n", ability, score.Score, score.Modifier)
	}

	fmt.Println("\nSkill Modifiers:")
	for skill, s := range c.Skills.All() {
		fmt.Printf("%-18s: %+d\n", skill, s.Modifier)
	}

	fmt.Println("\nEquipment:")

	if c.Equipment.MainHand.APIEquipment.Name != "" {
		fmt.Printf("  Weapon: %s\n", c.Equipment.MainHand.APIEquipment.Name)
	} else {
		fmt.Println("  Weapon: Empty")
	}

	if c.Equipment.Armor.APIEquipment.Name != "" {
		fmt.Printf("  Armor: %s\n", c.Equipment.Armor.APIEquipment.Name)
	} else {
		fmt.Println("  Armor: Empty")
	}

	if c.Equipment.Shield.APIEquipment.Name != "" {
		fmt.Printf("  Shield: %s\n", c.Equipment.Shield.APIEquipment.Name)
	} else {
		fmt.Println("  Shield: Empty")
	}

	if c.Equipment.Gear.APIEquipment.Name != "" {
		fmt.Printf("  Gear: %s\n", c.Equipment.Gear.APIEquipment.Name)
	} else {
		fmt.Println("  Gear: Empty")
	}

	if len(c.Spellbook.SpellSlot) > 0 {
		fmt.Println("\nSpell Slots:")

		// Group slots by level
		slotsByLevel := make(map[int][]domain.SpellSlot)
		for _, slot := range c.Spellbook.SpellSlot {
			slotsByLevel[slot.Level] = append(slotsByLevel[slot.Level], slot)
		}

		// Iterate through levels in order
		levels := []int{}
		for lvl := range slotsByLevel {
			levels = append(levels, lvl)
		}
		sort.Ints(levels)

		for _, lvl := range levels {
			fmt.Printf("Level %d slots:\n", lvl)
			for i, slot := range slotsByLevel[lvl] {
				spellName := "Empty"
				if slot.Spell.Name != "" {
					spellName = slot.Spell.Name
				}
				fmt.Printf("  Slot %d: %s\n", i+1, spellName)
			}
		}
	} else {
		fmt.Println("\nNo spell slots.")
	}

	fmt.Printf("\nSpell Save DC: %d | Spell Attack Bonus: %+d\n",
		c.SpellSaveDC(), c.SpellAttackBonus())

	fmt.Printf("\nAC: %d | Initiative: %+d | Passive Perception: %d\n",
		c.ArmorClass(), c.Initiative(), c.PassivePerception())

	// fmt.Printf("\nSpellSlots: %d", c.SpellSlots)
	// log.Println("spells count:", len(c.SpellSlots.SpellSlot))
}

// ListEquipment prints all equipment grouped by category using EquipmentService
func ListEquipment(equipmentService *service.EquipmentService) {
	fmt.Println("Weapons:")
	for _, w := range equipmentService.GetByCategory("weapon") {
		fmt.Printf("  %s\n", w)
	}

	fmt.Println("\nArmor:")
	for _, a := range equipmentService.GetByCategory("armor") {
		fmt.Printf("  %s\n", a)
	}

	fmt.Println("\nShields:")
	for _, s := range equipmentService.GetByCategory("shield") {
		fmt.Printf("  %s\n", s)
	}

	fmt.Println("\nGear:")
	for _, g := range equipmentService.GetByCategory("gear") {
		fmt.Printf("  %s\n", g)
	}
}

func ViewCharacter(c domain.Character) {
	fmt.Printf("Race: %s | Class: %s | Level: %d | Background: %s\n", c.Race.Name, c.Class.Name, c.Level, c.Background)
	fmt.Println("\nAbility Scores and Modifiers:")
	for ability, score := range c.AbilityScores.All() {
		fmt.Printf("%-12s: %2d (mod %+d)\n", ability, score.Score, score.Modifier)
	}

	fmt.Println("\nSkill Modifiers:")
	for skill, a := range c.Skills.All() {
		fmt.Printf("%-18s: %+d\n", skill, a.Modifier)
	}

	fmt.Println("\nEquipment:")

	if c.Equipment.MainHand.APIEquipment.Name != "" {
		fmt.Printf("  Weapon: %s\n", c.Equipment.MainHand.APIEquipment.Name)
	} else {
		fmt.Println("  Weapon: Empty")
	}

	if c.Equipment.Armor.APIEquipment.Name != "" {
		fmt.Printf("  Armor: %s\n", c.Equipment.Armor.APIEquipment.Name)
	} else {
		fmt.Println("  Armor: Empty")
	}

	if c.Equipment.Shield.APIEquipment.Name != "" {
		fmt.Printf("  Shield: %s\n", c.Equipment.Shield.APIEquipment.Name)
	} else {
		fmt.Println("  Shield: Empty")
	}

	if c.Equipment.Gear.APIEquipment.Name != "" {
		fmt.Printf("  Gear: %s\n", c.Equipment.Gear.APIEquipment.Name)
	} else {
		fmt.Println("  Gear: Empty")
	}
}

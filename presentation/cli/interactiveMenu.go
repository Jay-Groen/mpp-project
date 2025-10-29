package cli

// // readChoice reads a number from stdin between 1 and max
// func readChoice(max int, reader *bufio.Reader) int {
// 	for {
// 		input, _ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)
// 		num, err := strconv.Atoi(input)
// 		if err == nil && num >= 1 && num <= max {
// 			return num
// 		}
// 		fmt.Printf("Please enter a number between 1 and %d: ", max)
// 	}
// }

// // contains checks if a slice contains a string
// func contains(slice []string, val string) bool {
// 	for _, s := range slice {
// 		if s == val {
// 			return true
// 		}
// 	}
// 	return false
// }

// // printCharacterSheet prints abilities, modifiers, skills, equipment, and spell slots
// func printCharacterSheet(c *domain.Character) {
// 	fmt.Printf("\n--- Character Sheet: %s ---\n", c.Name)
// 	fmt.Printf("Race: %s | Class: %s | Level: %d | Background: %s\n",
// 		c.Race.Name, c.Class.Name, c.Level, c.Background)

// 	fmt.Println("\nAbility Scores and Modifiers:")
// 	for ability, score := range c.AbilityScores.All() {
// 		fmt.Printf("%-12s: %2d (mod %+d)\n", ability, score.Score, score.Modifier)
// 	}

// 	fmt.Println("\nSkill Modifiers:")
// 	for skill, s := range c.Skills.All() {
// 		fmt.Printf("%-18s: %+d\n", skill, s.Modifier)
// 	}

// 	fmt.Println("\nEquipment:")

// 	if c.Equipment.MainHand.EquipmentSpecific.Name != "" {
// 		fmt.Printf("  Weapon: %s\n", c.Equipment.MainHand.EquipmentSpecific.Name)
// 	} else {
// 		fmt.Println("  Weapon: Empty")
// 	}

// 	if c.Equipment.Armor.EquipmentSpecific.Name != "" {
// 		fmt.Printf("  Armor: %s\n", c.Equipment.Armor.EquipmentSpecific.Name)
// 	} else {
// 		fmt.Println("  Armor: Empty")
// 	}

// 	if c.Equipment.Shield.EquipmentSpecific.Name != "" {
// 		fmt.Printf("  Shield: %s\n", c.Equipment.Shield.EquipmentSpecific.Name)
// 	} else {
// 		fmt.Println("  Shield: Empty")
// 	}

// 	if c.Equipment.Gear.EquipmentSpecific.Name != "" {
// 		fmt.Printf("  Gear: %s\n", c.Equipment.Gear.EquipmentSpecific.Name)
// 	} else {
// 		fmt.Println("  Gear: Empty")
// 	}

// 	if len(c.Spellbook.SpellSlots) > 0 {
// 		fmt.Println("\nSpell Slots:")

// 		// Group slots by level
// 		slotsByLevel := make(map[int][]domain.SpellSlot)
// 		for _, slot := range c.Spellbook.SpellSlots {
// 			slotsByLevel[slot.Level] = append(slotsByLevel[slot.Level], slot)
// 		}

// 		// Iterate through levels in order
// 		levels := []int{}
// 		for lvl := range slotsByLevel {
// 			levels = append(levels, lvl)
// 		}
// 		sort.Ints(levels)

// 		for _, lvl := range levels {
// 			fmt.Printf("Level %d slots:\n", lvl)
// 			for i, slot := range slotsByLevel[lvl] {
// 				spellName := "Empty"
// 				if slot.Spell.Name != "" {
// 					spellName = slot.Spell.Name
// 				}
// 				fmt.Printf("  Slot %d: %s\n", i+1, spellName)
// 			}
// 		}
// 	} else {
// 		fmt.Println("\nNo spell slots.")
// 	}

// 	fmt.Printf("\nSpell Save DC: %d | Spell Attack Bonus: %+d\n",
// 		c.SpellSaveDC(), c.SpellAttackBonus())

// 	fmt.Printf("\nAC: %d | Initiative: %+d | Passive Perception: %d\n",
// 		c.ArmorClass(), c.Initiative(), c.PassivePerception())

// 	// fmt.Printf("\nSpellSlots: %d", c.SpellSlots)
// 	// log.Println("spells count:", len(c.SpellSlots.SpellSlot))
// }

// // ListEquipment prints all equipment grouped by category using EquipmentService
// func ListEquipment(equipmentService *service.EquipmentService) {
// 	fmt.Println("Weapons:")
// 	for _, w := range equipmentService.GetEquipmentByCategory("weapon") {
// 		fmt.Printf("  %s\n", w)
// 	}

// 	fmt.Println("\nArmor:")
// 	for _, a := range equipmentService.GetEquipmentByCategory("armor") {
// 		fmt.Printf("  %s\n", a)
// 	}

// 	fmt.Println("\nShields:")
// 	for _, s := range equipmentService.GetEquipmentByCategory("shield") {
// 		fmt.Printf("  %s\n", s)
// 	}

// 	fmt.Println("\nGear:")
// 	for _, g := range equipmentService.GetEquipmentByCategory("gear") {
// 		fmt.Printf("  %s\n", g)
// 	}
// }

// func ViewCharacter(c domain.Character) {
// 	fmt.Printf("Race: %s | Class: %s | Level: %d | Background: %s\n", c.Race.Name, c.Class.Name, c.Level, c.Background)
// 	fmt.Println("\nAbility Scores and Modifiers:")
// 	for ability, score := range c.AbilityScores.All() {
// 		fmt.Printf("%-12s: %2d (mod %+d)\n", ability, score.Score, score.Modifier)
// 	}

// 	fmt.Println("\nSkill Modifiers:")
// 	for skill, a := range c.Skills.All() {
// 		fmt.Printf("%-18s: %+d\n", skill, a.Modifier)
// 	}

// 	fmt.Println("\nEquipment:")

// 	if c.Equipment.MainHand.EquipmentSpecific.Name != "" {
// 		fmt.Printf("  Weapon: %s\n", c.Equipment.MainHand.EquipmentSpecific.Name)
// 	} else {
// 		fmt.Println("  Weapon: Empty")
// 	}

// 	if c.Equipment.Armor.EquipmentSpecific.Name != "" {
// 		fmt.Printf("  Armor: %s\n", c.Equipment.Armor.EquipmentSpecific.Name)
// 	} else {
// 		fmt.Println("  Armor: Empty")
// 	}

// 	if c.Equipment.Shield.EquipmentSpecific.Name != "" {
// 		fmt.Printf("  Shield: %s\n", c.Equipment.Shield.EquipmentSpecific.Name)
// 	} else {
// 		fmt.Println("  Shield: Empty")
// 	}

// 	if c.Equipment.Gear.EquipmentSpecific.Name != "" {
// 		fmt.Printf("  Gear: %s\n", c.Equipment.Gear.EquipmentSpecific.Name)
// 	} else {
// 		fmt.Println("  Gear: Empty")
// 	}
// }

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

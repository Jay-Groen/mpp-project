package cli

import (
	"flag"
	"fmt"
	"herkansing/onion/domain"
	"herkansing/onion/presentation"
	"os"
)

func HandleTestCommand(app *presentation.App, races []domain.Race, classes []domain.Class, spells []domain.Spell, equipment []domain.Equipment) {
	flagSet := flag.NewFlagSet("test", flag.ExitOnError)
	name := flagSet.String("name", "", "Character name (optional, for character-based tests)")
	flagSet.Parse(os.Args[2:])

	// Optional: If a name is given, load the character
	var char *domain.Character
	if *name != "" {
		c := presentation.FindCharacterByName(app, *name)
		char = &c
		fmt.Printf("🧙 Loaded character: %s (Level %d %s)\n", c.Name, c.Level, c.Class.Name)
	}

	// 🔬 Run your test logic here
	RunTestScenario(app, char, races, classes, spells, equipment)

	os.Exit(0)
}

func RunTestScenario(app *presentation.App, char *domain.Character, races []domain.Race, classes []domain.Class, spells []domain.Spell, equipment []domain.Equipment) {
	fmt.Println("🧪 Running test scenario...")

	c := presentation.FindCharacterByName(app, "funnybrie")

	fmt.Println(c.Name)
	fmt.Println(c.Class)
	fmt.Println(c.Class.HitDie)
	fmt.Println(c.MaxHP)

	// fmt.Println(characters)
	// fmt.Println(races)
	// fmt.Println(classes)
	// fmt.Println(spells)
	// fmt.Println(equipment)

	fmt.Println("✅ Test finished.")
}

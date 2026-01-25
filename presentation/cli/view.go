package cli

import (
	"flag"
	"fmt"
	"herkansing/onion/domain"
	"herkansing/onion/presentation"
	"os"
)

func HandleViewCommand(app *presentation.App, classes []domain.Class) {
	flagSet := flag.NewFlagSet("view", flag.ExitOnError)
	name := flagSet.String("name", "", "Character name")

	if err := flagSet.Parse(os.Args[2:]); err != nil {
		fmt.Println(err)
		return
	}

	presentation.ValidateRequired(map[string]*string{
		"name": name,
	})

	nameInput := *name
	c := presentation.FindCharacterByName(app, nameInput)

	// Get class reference from loaded classes
	selectedClass := presentation.FindClass(c.Class.Name, classes)

	presentation.PrintCharacterDetails(c)
	presentation.PrintEquipment(c)
	presentation.PrintSpellInfo(c, selectedClass)
	presentation.PrintCombatStats(c)
	os.Exit(0)
}

package cli

import (
	"flag"
	"fmt"
	"herkansing/onion/presentation"
	"os"
)

func HandleDeleteCommand(app *presentation.App) {
	flagSet := flag.NewFlagSet("delete", flag.ExitOnError)
	name := flagSet.String("name", "", "Character name")

	if err := flagSet.Parse(os.Args[2:]); err != nil {
		fmt.Println(err)
		return
	}

	presentation.ValidateRequired(map[string]*string{
		"name": name,
	})

	nameInput := *name

	// 🔍 Find the character by name
	selectedChar := presentation.FindCharacterByName(app, nameInput)

	// 🗑️ Delete character
	err := app.CharacterService.DeleteCharacter(selectedChar.Id)
	if err != nil {
		fmt.Printf("❌ Failed to delete '%s': %v\n", selectedChar.Name, err)
		os.Exit(1)
	}

	fmt.Printf("deleted %s\n", selectedChar.Name)
	os.Exit(0)
}

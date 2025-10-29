package cli

import (
	"flag"
	"fmt"
	"herkansing/onion/domain"
	"herkansing/onion/presentation"
	"log"
	"os"
)

func CreateHandler(app *presentation.App, classes []domain.Class, races []domain.Race) {
	flagSet := flag.NewFlagSet("create", flag.ExitOnError)

	name := flagSet.String("name", "", "Character name")
	race := flagSet.String("race", "", "Character race")
	class := flagSet.String("class", "", "Character class")
	level := flagSet.Int("level", 1, "Character level")
	str := flagSet.Int("str", 10, "Strength score")
	dex := flagSet.Int("dex", 10, "Dexterity score")
	con := flagSet.Int("con", 13, "Constitution score")
	intt := flagSet.Int("int", 8, "Intelligence score")
	wis := flagSet.Int("wis", 12, "Wisdom score")
	cha := flagSet.Int("cha", 14, "Charisma score")

	flagSet.Parse(os.Args[2:])

	presentation.ValidateRequired(map[string]*string{
		"name":  name,
		"class": class,
		"race":  race,
	})

	// ✅ Extracted helpers
	selectedClass := presentation.FindClass(*class, classes)
	selectedRace := presentation.FindRace(*race, races)
	chosenSkills := presentation.GetChosenSkillsForName(*name, selectedClass.SkillProficiencies)

	newChar, err := app.CharacterService.CreateCharacter(
		*name,
		selectedRace,
		selectedClass,
		"acolyte",
		[]string{},   // chosenAbilities
		chosenSkills, // chosenSkills
		*level, *str, *dex, *con, *intt, *wis, *cha,
	)
	if err != nil {
		log.Fatalf("Failed to create character: %v", err)
	}

	fmt.Printf("saved character %s\n", newChar.Name)
}

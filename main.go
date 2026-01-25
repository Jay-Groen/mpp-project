package main

import (
	"fmt"
	"herkansing/onion/presentation"
	"herkansing/onion/presentation/cli"
	"log"
	"os"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app, db := presentation.InitializeApp()

	classes, races, spells, equipment := presentation.LoadData(*app)

	defer db.Close()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <command> [flags]")
		fmt.Println("Available commands: create, view, delete, equip, prepare-spell, learn-spell, test, html")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "create":
		cli.CreateHandler(app, classes, races)

	case "view":
		cli.HandleViewCommand(app, classes)

	case "delete":
		cli.HandleDeleteCommand(app)

	case "equip":
		cli.HandleEquipCommand(app)

	case "prepare-spell":
		cli.HandlePrepareSpellCommand(app, classes, spells)

	case "learn-spell":
		cli.HandleLearnSpellCommand(app, classes, spells)

	case "test":
		cli.HandleTestCommand(app, races, classes, spells, equipment)

	case "html":
		log.Println("Starting web server...")
		cmd := exec.Command("go", "run", "./presentation/web/web.go")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to run web.go: %v", err)
		}

	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		fmt.Println("Available commands: create, view, delete, equip, prepare-spell, learn-spell, test, html")
		os.Exit(1)
	}

}

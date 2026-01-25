package cli

import (
	"flag"
	"fmt"
	"herkansing/onion/domain"
	"herkansing/onion/presentation"
	"log"
	"os"
	"strings"
)

func HandleEquipCommand(app *presentation.App) {
	flagSet := flag.NewFlagSet("equip", flag.ExitOnError)
	name := flagSet.String("name", "", "Character name (required)")
	weapon := flagSet.String("weapon", "", "Weapon to equip")
	slot := flagSet.String("slot", "", "Slot for weapon: 'main hand' or 'off hand'")
	armor := flagSet.String("armor", "", "Armor to equip")
	shield := flagSet.String("shield", "", "Shield to equip")

	if err := flagSet.Parse(os.Args[2:]); err != nil {
		fmt.Println(err)
		return
	}

	presentation.ValidateRequired(map[string]*string{
		"name": name,
	})

	char := presentation.FindCharacterByName(app, *name)

	// ⚔️ Equip items
	if *weapon != "" {
		if *slot == "" {
			fmt.Println("Error: -slot is required when equipping a weapon")
			os.Exit(1)
		}
		HandleWeaponEquip(app, &char, *weapon, *slot)
	}

	if *armor != "" {
		HandleArmorEquip(app, &char, *armor)
	}

	if *shield != "" {
		HandleShieldEquip(app, &char, *shield)
	}

	// 💾 Save updates
	if err := app.CharacterService.UpdateCharacter(char); err != nil {
		log.Fatalf("Failed to update equipment: %v", err)
	}
	os.Exit(0)
}

func HandleWeaponEquip(app *presentation.App, c *domain.Character, weaponName, slot string) {
	slotLower := strings.ToLower(slot)
	item := domain.EquipmentSpecific{Name: weaponName}

	switch slotLower {
	case "main hand":
		if c.Equipment.MainHand.EquipmentSpecific.Name != "" {
			fmt.Printf("main hand already occupied\n")
			os.Exit(1)
		}
		app.EquipmentService.AddEquipmentToCharacter(c, slotLower, item)
		fmt.Printf("Equipped weapon %s to main hand\n", weaponName)

	case "off hand":
		if c.Equipment.OffHand.EquipmentSpecific.Name != "" {
			fmt.Printf("off hand already occupied\n")
			os.Exit(1)
		}
		app.EquipmentService.AddEquipmentToCharacter(c, slotLower, item)
		fmt.Printf("Equipped weapon %s to off hand\n", weaponName)

	default:
		fmt.Println("Error: slot must be 'main hand' or 'off hand'")
		os.Exit(1)
	}
}

func HandleArmorEquip(app *presentation.App, c *domain.Character, armorName string) {
	if c.Equipment.Armor.EquipmentSpecific.Name != "" {
		fmt.Printf("Armor slot already occupied by %s\n", c.Equipment.Armor.EquipmentSpecific.Name)
		os.Exit(1)
	}
	item := domain.EquipmentSpecific{Name: armorName}
	app.EquipmentService.AddEquipmentToCharacter(c, "armor", item)
	fmt.Printf("Equipped armor %s\n", armorName)
}

func HandleShieldEquip(app *presentation.App, c *domain.Character, shieldName string) {
	if c.Equipment.Shield.EquipmentSpecific.Name != "" {
		fmt.Printf("Shield slot already occupied by %s\n", c.Equipment.Shield.EquipmentSpecific.Name)
		os.Exit(1)
	}
	item := domain.EquipmentSpecific{Name: shieldName}
	app.EquipmentService.AddEquipmentToCharacter(c, "shield", item)
	fmt.Printf("Equipped shield %s\n", shieldName)
}

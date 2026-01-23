package test

import (
    "herkansing/onion/domain"
    "testing"
)

func TestArmorClass_WithLightArmorAndShield(t *testing.T) {
    c := domain.Character{
        Class: domain.Class{Name: "Fighter"},
        AbilityScores: domain.AbilityScores{
            Dexterity: domain.AbilityScore{Modifier: 3},
            Constitution: domain.AbilityScore{Modifier: 1},
            Wisdom: domain.AbilityScore{Modifier: 0},
        },
        Equipment: domain.Equipment{
            Armor:  domain.Armor{EquipmentSpecific: domain.EquipmentSpecific{Name: "Leather Armor"}},
            Shield: domain.Shield{EquipmentSpecific: domain.EquipmentSpecific{Name: "Shield"}},
        },
    }

    // Leather: 11 + DEX (3) = 14, +2 shield = 16
    if got, want := c.ArmorClass(), 16; got != want {
        t.Fatalf("ArmorClass() = %d, want %d", got, want)
    }
}

func TestArmorClass_MediumArmorCapsDexBonus(t *testing.T) {
    c := domain.Character{
        Class: domain.Class{Name: "Ranger"},
        AbilityScores: domain.AbilityScores{
            Dexterity: domain.AbilityScore{Modifier: 4},
        },
        Equipment: domain.Equipment{
            Armor: domain.Armor{EquipmentSpecific: domain.EquipmentSpecific{Name: "Half Plate"}},
        },
    }

    // Half Plate: 15 + min(DEX,2) => 15 + 2 = 17
    if got, want := c.ArmorClass(), 17; got != want {
        t.Fatalf("ArmorClass() = %d, want %d", got, want)
    }
}

func TestArmorClass_UnarmoredBarbarianUsesDexPlusCon(t *testing.T) {
    c := domain.Character{
        Class: domain.Class{Name: "Barbarian"},
        AbilityScores: domain.AbilityScores{
            Dexterity: domain.AbilityScore{Modifier: 2},
            Constitution: domain.AbilityScore{Modifier: 3},
        },
        Equipment: domain.Equipment{}, // no armor
    }

    // Barbarian unarmored: 10 + DEX + CON = 15
    if got, want := c.ArmorClass(), 15; got != want {
        t.Fatalf("ArmorClass() = %d, want %d", got, want)
    }
}

func TestArmorClass_MonkDoesNotGetShieldBonus(t *testing.T) {
    c := domain.Character{
        Class: domain.Class{Name: "Monk"},
        AbilityScores: domain.AbilityScores{
            Dexterity: domain.AbilityScore{Modifier: 3},
            Wisdom: domain.AbilityScore{Modifier: 2},
        },
        Equipment: domain.Equipment{
            Shield: domain.Shield{EquipmentSpecific: domain.EquipmentSpecific{Name: "Shield"}},
        },
    }

    // Monk unarmored: 10 + DEX + WIS = 15; shield should NOT add +2 for monks
    if got, want := c.ArmorClass(), 15; got != want {
        t.Fatalf("ArmorClass() = %d, want %d", got, want)
    }
}

func TestPassivePerception_AddsProficiencyWhenProficient(t *testing.T) {
    c := domain.Character{
        ProficiencyBonus: 2,
        AbilityScores: domain.AbilityScores{
            Wisdom: domain.AbilityScore{Modifier: 1},
        },
        Skills: domain.Skills{
            Perception: domain.Skill{Proficient: true},
        },
    }

    // 10 + WIS(1) + PB(2) = 13
    if got, want := c.PassivePerception(), 13; got != want {
        t.Fatalf("PassivePerception() = %d, want %d", got, want)
    }
}

func TestInitiative_UsesDexModifier(t *testing.T) {
    c := domain.Character{
        AbilityScores: domain.AbilityScores{
            Dexterity: domain.AbilityScore{Modifier: 4},
        },
    }
    if got, want := c.Initiative(), 4; got != want {
        t.Fatalf("Initiative() = %d, want %d", got, want)
    }
}

package domain

import (
	"herkansing/onion/dndapi"
)

type Spell struct {
	Name    string   `json:"name"`
	Level   int      `json:"level"`
	Classes []string `json:"classes"`
}

type SpellSlot struct {
	Spell      dndapi.APISpell `json:"spell"`
	Level      int             `json:"level"`
}

type Spellbook struct {
	SpellSlot []SpellSlot `json:"spell_slot"`
}

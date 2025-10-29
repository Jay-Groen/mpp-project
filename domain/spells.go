package domain

type SpellSlot struct {
	Spell Spell `json:"spell"`
	Level int   `json:"level"`
}

type Spellbook struct {
	SpellSlots []SpellSlot `json:"spell_slots"`
}

type Spell struct {
	Name       string       `json:"name"`
	Level      int          `json:"level"`
	Desc       []string     `json:"desc,omitempty"`
	Range      string       `json:"range,omitempty"`
	Duration   string       `json:"duration,omitempty"`
	Classes    []string     `json:"classes"`
	Damage     SpellsDamage `json:"damage,omitempty"`
	DC         DC           `json:"dc,omitempty"`
	School     School       `json:"school,omitempty"`
	IsPrepared bool         `json:"is_prepared,omitempty"`
	IsKnown    bool         `json:"is_known,omitempty"`
}

type School struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type SpellsDamage struct {
	DamageType             DamageType        `json:"damage_type"`
	DamageAtCharacterLevel map[string]string `json:"damage_at_character_level"`
}

type DamageType struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type DC struct {
	DCType    DCType `json:"dc_type"`
	DCSuccess string `json:"dc_success"`
}

type DCType struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}
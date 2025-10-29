package domain

type SpellAPIFetcher interface {
    FetchSpell(name string) (Spell, error)
    FetchMultipleSpells(names []string) ([]Spell, error)
}

type SpellCSVLoader interface {
    LoadSpells() ([]Spell, error)
}
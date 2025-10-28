package domain

type AbilityScores struct {
	Strength     AbilityScore `json:"strength"`
	Dexterity    AbilityScore `json:"dexterity"`
	Constitution AbilityScore `json:"constitution"`
	Intelligence AbilityScore `json:"intelligence"`
	Wisdom       AbilityScore `json:"wisdom"`
	Charisma     AbilityScore `json:"charisma"`
}

type AbilityScore struct {
	Name     string `json:"name"`
	Score    int    `json:"score"`
	Modifier int    `json:"modifier"`
}

// Ability constants for reference
const (
	Strength     = "Strength"
	Dexterity    = "Dexterity"
	Constitution = "Constitution"
	Intelligence = "Intelligence"
	Wisdom       = "Wisdom"
	Charisma     = "Charisma"
)

// Modifier calculates the D&D ability modifier from a score
func Modifier(score int) int {
	mod := (score - 10) / 2
	if (score-10)%2 != 0 && score < 10 {
		mod-- // adjust for negative rounding
	}
	return mod
}

// All returns a map of all abilities and their scores for iteration
func (a AbilityScores) All() map[string]AbilityScore {
	return map[string]AbilityScore{
		Strength:     a.Strength,
		Dexterity:    a.Dexterity,
		Constitution: a.Constitution,
		Intelligence: a.Intelligence,
		Wisdom:       a.Wisdom,
		Charisma:     a.Charisma,
	}
}

func (a *AbilityScores) ApplyModifiers() {
	a.Strength.Modifier = Modifier(a.Strength.Score)
	a.Dexterity.Modifier = Modifier(a.Dexterity.Score)
	a.Constitution.Modifier = Modifier(a.Constitution.Score)
	a.Intelligence.Modifier = Modifier(a.Intelligence.Score)
	a.Wisdom.Modifier = Modifier(a.Wisdom.Score)
	a.Charisma.Modifier = Modifier(a.Charisma.Score)
}

// ApplyRacialBonuses applies a map of bonuses to the AbilityScores struct
func (a *AbilityScores) ApplyRacialBonuses(bonuses map[string]int) {
	if val, ok := bonuses[Strength]; ok {
		a.Strength.Score += val
	}
	if val, ok := bonuses[Dexterity]; ok {
		a.Dexterity.Score += val
	}
	if val, ok := bonuses[Constitution]; ok {
		a.Constitution.Score += val
	}
	if val, ok := bonuses[Intelligence]; ok {
		a.Intelligence.Score += val
	}
	if val, ok := bonuses[Wisdom]; ok {
		a.Wisdom.Score += val
	}
	if val, ok := bonuses[Charisma]; ok {
		a.Charisma.Score += val
	}
}

// Modifiers returns a map of all abilities and their modifiers
func (a AbilityScores) Modifiers() map[string]int {
	mods := make(map[string]int)
	for key, score := range a.All() {
		mods[key] = Modifier(score.Score)
	}
	return mods
}

package domain

type Race struct {
	Name                string         `json:"name"`
	AbilityScoreBonuses map[string]int `json:"ability_score_bonuses"` // e.g. {"Constitution": 2, "Wisdom": 1}
	Choice              bool           `json:"choice"`
	ChoiceAmount        int            `json:"choice_amount"`
	ChoiceAddAmount     int            `json:"choice_add_amount"`
}

// ApplyBonuses applies the racial bonuses to a character's AbilityScores
func (r Race) ApplyBonuses(a *AbilityScores) {
	a.ApplyRacialBonuses(r.AbilityScoreBonuses)
}

// HandleChoice applies additional chosen bonuses if the race has optional choices
// 'chosen' is a slice of ability names the player chooses to apply the extra points to
func (r Race) HandleChoice(a *AbilityScores, chosen []string) {
	if !r.Choice || len(chosen) == 0 {
		return
	}

	for _, ability := range chosen {
		switch ability {
		case Strength:
			a.Strength.Score += r.ChoiceAddAmount
		case Dexterity:
			a.Dexterity.Score += r.ChoiceAddAmount
		case Constitution:
			a.Constitution.Score += r.ChoiceAddAmount
		case Intelligence:
			a.Intelligence.Score += r.ChoiceAddAmount
		case Wisdom:
			a.Wisdom.Score += r.ChoiceAddAmount
		case Charisma:
			a.Charisma.Score += r.ChoiceAddAmount
		}
	}
}
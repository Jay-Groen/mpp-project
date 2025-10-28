package domain

type Skill struct {
	Name       string `json:"name"`
	Ability    string `json:"ability"`
	Proficient bool   `json:"proficient"`
	Modifier   int    `json:"modifier"`
}

type Skills struct {
	Acrobatics     Skill `json:"acrobatics"`      // DEX
	AnimalHandling Skill `json:"animal_handling"` // WIS
	Arcana         Skill `json:"arcana"`          // INT
	Athletics      Skill `json:"athletics"`       // STR
	Deception      Skill `json:"deception"`       // CHA
	History        Skill `json:"history"`         // INT
	Insight        Skill `json:"insight"`         // WIS
	Intimidation   Skill `json:"intimidation"`    // CHA
	Investigation  Skill `json:"investigation"`   // INT
	Medicine       Skill `json:"medicine"`        // WIS
	Nature         Skill `json:"nature"`          // INT
	Perception     Skill `json:"perception"`      // WIS
	Performance    Skill `json:"performance"`     // CHA
	Persuasion     Skill `json:"persuasion"`      // CHA
	Religion       Skill `json:"religion"`        // INT
	SleightOfHand  Skill `json:"sleight_of_hand"` // DEX
	Stealth        Skill `json:"stealth"`         // DEX
	Survival       Skill `json:"survival"`        // WIS
}

// NewSkills initializes all skills with Name and Ability
func NewSkills() Skills {
	return Skills{
		Acrobatics:     Skill{Name: "Acrobatics", Ability: "Dexterity"},
		AnimalHandling: Skill{Name: "Animal Handling", Ability: "Wisdom"},
		Arcana:         Skill{Name: "Arcana", Ability: "Intelligence"},
		Athletics:      Skill{Name: "Athletics", Ability: "Strength"},
		Deception:      Skill{Name: "Deception", Ability: "Charisma"},
		History:        Skill{Name: "History", Ability: "Intelligence"},
		Insight:        Skill{Name: "Insight", Ability: "Wisdom"},
		Intimidation:   Skill{Name: "Intimidation", Ability: "Charisma"},
		Investigation:  Skill{Name: "Investigation", Ability: "Intelligence"},
		Medicine:       Skill{Name: "Medicine", Ability: "Wisdom"},
		Nature:         Skill{Name: "Nature", Ability: "Intelligence"},
		Perception:     Skill{Name: "Perception", Ability: "Wisdom"},
		Performance:    Skill{Name: "Performance", Ability: "Charisma"},
		Persuasion:     Skill{Name: "Persuasion", Ability: "Charisma"},
		Religion:       Skill{Name: "Religion", Ability: "Intelligence"},
		SleightOfHand:  Skill{Name: "Sleight of Hand", Ability: "Dexterity"},
		Stealth:        Skill{Name: "Stealth", Ability: "Dexterity"},
		Survival:       Skill{Name: "Survival", Ability: "Wisdom"},
	}
}

// All returns a map of all skills for iteration
func (s Skills) All() map[string]Skill {
	return map[string]Skill{
		"Acrobatics":      s.Acrobatics,
		"Animal Handling": s.AnimalHandling,
		"Arcana":          s.Arcana,
		"Athletics":       s.Athletics,
		"Deception":       s.Deception,
		"History":         s.History,
		"Insight":         s.Insight,
		"Intimidation":    s.Intimidation,
		"Investigation":   s.Investigation,
		"Medicine":        s.Medicine,
		"Nature":          s.Nature,
		"Perception":      s.Perception,
		"Performance":     s.Performance,
		"Persuasion":      s.Persuasion,
		"Religion":        s.Religion,
		"Sleight Of Hand": s.SleightOfHand,
		"Stealth":         s.Stealth,
		"Survival":        s.Survival,
	}
}

// SkillModifier calculates the modifier for a single skill
// given the character's ability scores and proficiency bonus
func (s Skills) SkillModifier(skill Skill, abilities AbilityScores, profBonus int) int {
	// Get ability score associated with this skill
	var abilityScore int
	switch skill.Ability {
	case Strength:
		abilityScore = abilities.Strength.Score
	case Dexterity:
		abilityScore = abilities.Dexterity.Score
	case Constitution:
		abilityScore = abilities.Constitution.Score
	case Intelligence:
		abilityScore = abilities.Intelligence.Score
	case Wisdom:
		abilityScore = abilities.Wisdom.Score
	case Charisma:
		abilityScore = abilities.Charisma.Score
	default:
		abilityScore = 0
	}

	mod := Modifier(abilityScore)
	if skill.Proficient {
		mod += profBonus
	}
	return mod
}

// Modifiers calculates each skill's modifier based on ability and proficiency bonus
func (s *Skills) Modifiers(abilities AbilityScores, proficiency int) map[string]int {
	result := make(map[string]int)

	apply := func(skill *Skill, ability AbilityScore) {
		// mod := (ability.Score - 10) / 2
		mod := ability.Modifier
		if skill.Proficient {
			mod += proficiency
		}
		skill.Modifier = mod // update new field
		result[skill.Name] = mod
	}

	apply(&s.Acrobatics, abilities.Dexterity)
	apply(&s.AnimalHandling, abilities.Wisdom)
	apply(&s.Arcana, abilities.Intelligence)
	apply(&s.Athletics, abilities.Strength)
	apply(&s.Deception, abilities.Charisma)
	apply(&s.History, abilities.Intelligence)
	apply(&s.Insight, abilities.Wisdom)
	apply(&s.Intimidation, abilities.Charisma)
	apply(&s.Investigation, abilities.Intelligence)
	apply(&s.Medicine, abilities.Wisdom)
	apply(&s.Nature, abilities.Intelligence)
	apply(&s.Perception, abilities.Wisdom)
	apply(&s.Performance, abilities.Charisma)
	apply(&s.Persuasion, abilities.Charisma)
	apply(&s.Religion, abilities.Intelligence)
	apply(&s.SleightOfHand, abilities.Dexterity)
	apply(&s.Stealth, abilities.Dexterity)
	apply(&s.Survival, abilities.Wisdom)

	return result
}

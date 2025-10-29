package domain

func ApplyModifiers(s *Skills, modifiers map[string]int) {
	for skillName, mod := range modifiers {
		switch skillName {
		case "Acrobatics":
			s.Acrobatics.Modifier = mod
		case "Animal Handling":
			s.AnimalHandling.Modifier = mod
		case "Arcana":
			s.Arcana.Modifier = mod
		case "Athletics":
			s.Athletics.Modifier = mod
		case "Deception":
			s.Deception.Modifier = mod
		case "History":
			s.History.Modifier = mod
		case "Insight":
			s.Insight.Modifier = mod
		case "Intimidation":
			s.Intimidation.Modifier = mod
		case "Investigation":
			s.Investigation.Modifier = mod
		case "Medicine":
			s.Medicine.Modifier = mod
		case "Nature":
			s.Nature.Modifier = mod
		case "Perception":
			s.Perception.Modifier = mod
		case "Performance":
			s.Performance.Modifier = mod
		case "Persuasion":
			s.Persuasion.Modifier = mod
		case "Religion":
			s.Religion.Modifier = mod
		case "Sleight of Hand":
			s.SleightOfHand.Modifier = mod
		case "Stealth":
			s.Stealth.Modifier = mod
		case "Survival":
			s.Survival.Modifier = mod
		}
	}
}

// SkillModifiers returns a map of all skills and their calculated modifiers
func (c *Character) SkillModifiers() map[string]int {
	return c.Skills.Modifiers(c.AbilityScores, c.ProficiencyBonus)
}
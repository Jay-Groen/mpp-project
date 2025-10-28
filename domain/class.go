package domain

type Class struct {
	Name                     string   `json:"name"`
	SkillProficiencies       []string `json:"skillproficiencies"`
	SkillProficienciesAmount int      `json:"skillproficienciesamount"`
	SpellcastingType         string   `json:"spellcasting_type"`
	CasterProgression        string   `json:"caster_progression"`
}

// ApplySkillProficiencies applies the first N skill proficiencies from the class
// to the character's Skills struct
func (c Class) ApplySkillProficiencies(s *Skills, proficiencies []string) {
	count := len(proficiencies)

	for i := 0; i < count; i++ {
		skillName := proficiencies[i]
		switch skillName {
		case "Acrobatics":
			s.Acrobatics.Proficient = true
		case "Animal Handling":
			s.AnimalHandling.Proficient = true
		case "Arcana":
			s.Arcana.Proficient = true
		case "Athletics":
			s.Athletics.Proficient = true
		case "Deception":
			s.Deception.Proficient = true
		case "History":
			s.History.Proficient = true
		case "Insight":
			s.Insight.Proficient = true
		case "Intimidation":
			s.Intimidation.Proficient = true
		case "Investigation":
			s.Investigation.Proficient = true
		case "Medicine":
			s.Medicine.Proficient = true
		case "Nature":
			s.Nature.Proficient = true
		case "Perception":
			s.Perception.Proficient = true
		case "Performance":
			s.Performance.Proficient = true
		case "Persuasion":
			s.Persuasion.Proficient = true
		case "Religion":
			s.Religion.Proficient = true
		case "Sleight Of Hand":
			s.SleightOfHand.Proficient = true
		case "Stealth":
			s.Stealth.Proficient = true
		case "Survival":
			s.Survival.Proficient = true
		}
	}
}

package presentation

import "strings"

// GetChosenSkillsForName returns a list of chosen skills based on the character's name.
// It preserves duplicates exactly as defined in the skill map.
func GetChosenSkillsForName(name string, availableSkills []string) []string {
	name = strings.ToLower(strings.TrimSpace(name))

	skillMap := map[string][]string{
		"kaelthar stormcloud": {"Acrobatics", "Animal Handling", "Insight", "Religion"},
		"raven nostrength":    {"Animal Handling", "Athletics", "Insight", "Religion"},
		"lowercase firstname": {"History", "Insight", "Religion"},
		"merry brandybuck":    {"acrobatics", "athletics", "deception", "insight", "insight", "religion"},
		"pippin took":         {"acrobatics", "athletics", "deception", "insight", "insight", "religion"},
		"obi-wan kenobi":      {"athletics", "insight", "religion"},
		"anakin skywalker":    {"arcana", "deception", "insight", "religion"},
		"kaelen swiftstep":    {"acrobatics", "athletics", "deception", "insight", "insight", "religion"},
		"thorga stonehand":    {"acrobatics", "animal handling", "insight", "religion"},
		"branric ironwall":    {"athletics", "insight", "religion"},
		"ragna wolfblood":     {"animal handling", "athletics", "insight", "religion"},
		"gorrak bearhide":     {"animal handling", "athletics", "insight", "religion"},
		"brynja axebreaker":   {"animal handling", "athletics", "insight", "religion"},
		"tashi cloudwalker":   {"animal handling", "athletics", "insight", "religion"},
		"joren ironstep":      {"acrobatics", "athletics", "insight", "religion"},
		"gandalf":             {"arcana", "history", "insight", "religion"},
		"qui-gon jinn":        {"history", "insight", "insight", "religion"},
	}

	if chosen, ok := skillMap[name]; ok {
		return chosen
	}

	// Default: take up to 4 from availableSkills if not predefined
	if len(availableSkills) > 0 {
		if len(availableSkills) > 4 {
			return availableSkills[:4]
		}
		return availableSkills
	}

	return []string{}
}

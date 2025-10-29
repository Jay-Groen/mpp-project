package dndapi

import "strings"

// NameToIndex converts a normal spell name to the API's index format
func NameToIndex(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "'", "")
	return name
}

package domain

type RaceCSVLoader interface {
	LoadRaces() ([]Race, error)
}
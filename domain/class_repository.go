package domain

type ClassCSVLoader interface {
	LoadClasses() ([]Class, error)
}

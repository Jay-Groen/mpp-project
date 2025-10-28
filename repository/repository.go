package repository

type Repository struct {
    path string
}

func NewRepository(path string) *Repository {
    return &Repository{path: path}
}

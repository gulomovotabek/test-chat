package repository

type Repository interface {
	// Add your methods here
}

func NewRepository() Repository {
	return &repository{}
}

type repository struct {
	// Add your implementation
}

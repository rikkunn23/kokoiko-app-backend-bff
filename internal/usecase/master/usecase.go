package master

// Usecase
type Usecase struct {
	repo IRepository
}

// NewUsecase ...
func New(repo IRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

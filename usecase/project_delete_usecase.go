package usecase

type ProjectDeleteUseCase interface {
	Execute(id uint) error
}

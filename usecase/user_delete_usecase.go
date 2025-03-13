package usecase

type UserDeleteUseCase interface {
	Execute(id uint) error
}

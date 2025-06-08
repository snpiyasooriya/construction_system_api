package usecase

type DeleteProjectTypeUseCaseInterface interface {
	Execute(id uint) error
}

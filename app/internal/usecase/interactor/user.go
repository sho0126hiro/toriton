package interactor

import "github.com/sho0126hiro/toriton/app/internal/usecase/repository"

type UserInteractor struct {
	UserRepository repository.UserRepository
}

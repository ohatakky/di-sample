package di

import "di-sample/repository"

func InjectUserRepository() repository.UserRepository {
	return repository.NewUserRepository(
		InjectDB(),
	)
}

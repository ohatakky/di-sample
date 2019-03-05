package di

import (
	"di-sample/usecase"
)

func InjectMailer() usecase.Mailer {
	return usecase.NewMailer()
}

func InjectSignUpService() usecase.SignUpService {
	return usecase.NewSignUpService(
		InjectUserRepository(),
		InjectMailer(),
	)
}

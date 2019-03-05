package usecase

import (
	"di-sample/model"
	"di-sample/repository"
)

type SignUpService interface {
	SignUp(email, password string) error
}

type signUpService struct {
	repo   repository.UserRepository
	mailer Mailer
}

func (s signUpService) SignUp(email, password string) error {
	u := model.User{Email: email, Password: password}
	if err := s.repo.Save(u); err != nil {
		return err
	}
	return s.mailer.SendEmail(email, "done")
}

func NewSignUpService(repo repository.UserRepository, mailer Mailer) SignUpService {
	return signUpService{
		repo,
		mailer,
	}
}

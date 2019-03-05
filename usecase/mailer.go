package usecase

type Mailer interface {
	SendEmail(to, message string) error
}

func NewMailer() Mailer {
	return nil
}

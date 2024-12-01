package models

type VerifyEmail struct {
	Email string
	Code  string
}

type WelcomeMessage struct {
	Name    string
	Email   string
	Message string
}

type ForgotPassword struct {
	Name  string
	Email string
	Code  string
}

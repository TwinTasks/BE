package users

type Users struct {
	Username string
	Email    string
	Password string
}

type Handler interface {
}

type Query interface {
}

type Service interface {
}
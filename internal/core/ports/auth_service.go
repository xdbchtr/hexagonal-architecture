package ports

type AuthService interface {
	Register(name, email, password string) error
	Login(email, password string) (string, error)
}

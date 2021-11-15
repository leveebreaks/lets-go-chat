package repository

type AuthRepository interface {
	CreateUser(userName, password string) (string, error)
	CheckUser(userName, password string) bool
}

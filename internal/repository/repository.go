package repository

type Auth interface {
	CreateUser(userName, password string) (string, error)
	CheckUser(userName, password string) bool
}

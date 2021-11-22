package repository

type AuthRepository interface {
	CreateUser(userName, password string) (string, error)
	CheckUser(userName, password string) bool
}

var authRepo AuthRepository = NewInMemoryAuthRepo()

func InitAuthRepository(ar AuthRepository) {
	authRepo = ar
}

func CreateUser(userName, password string) (string, error) {
	return authRepo.CreateUser(userName, password)
}

func CheckUser(userName, password string) bool {
	return authRepo.CheckUser(userName, password)
}

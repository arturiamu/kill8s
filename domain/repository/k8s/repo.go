package k8s

var globalRepository Repository

type Repository interface {
}

func GetRepository() Repository {
	return globalRepository
}

func SetRepository(repo Repository) {
	globalRepository = repo
}

package storage

type Driver interface {
	FindByID(id string) (*Token, error)
	FindByHash(hash string) (*Token, error)
	RevokeToken(id string) error
	TouchLastUsed(id string) error
	StoreToken(t *Token) error
}

var driver Driver

func Setup(d Driver) {
	driver = d
}

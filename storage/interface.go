package storage

type Driver interface {
	FindByID(id int64) (*Token, error)
	FindByHash(hash string) (*Token, error)
	RevokeToken(id int64) error
	TouchLastUsed(id int64) error
	StoreToken(t *Token) error
}

var driver Driver

func Setup(d Driver) {
	driver = d
}

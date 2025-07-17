package storage

type Driver interface {
	FindByID(id int64) (*Token, error)
	FindByHash(hash string) (*Token, error)
	RevokeToken(hash string) error
	TouchLastUsed(hash string) error
	StoreToken(t *Token) error
}

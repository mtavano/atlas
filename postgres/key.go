package postgres

import (
	"errors"

	"github.com/mtavano/atlas"
)

// KeyStore errors
var (
	ErrMissingToken = errors.New("missing token, token must not be empty")
)

// KeyStore ...
type KeyStore struct {
	*Datastore
}

// FindKeyByToken ...
func (kst *KeyStore) FindKeyByToken(token string) (*atlas.Key, error) {
	if token == "" {
		return nil, ErrMissingToken
	}

	k := new(atlas.Key)

	query := `SELECT * FROM keys WHERE token = $1`
	if err := kst.SQL(query, token).QueryStruct(k); err != nil {
		return nil, err
	}
	return k, nil
}

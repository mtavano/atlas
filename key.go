package atlas

import "time"

// Key ...
type Key struct {
	ID     int    `db:"id"`
	GateID int    `db:"gate_id"`
	UserID int    `db:"user_id"`
	Name   string `db:"name"`
	Token  string `db:"token"`

	ExpiredAt *time.Time `db:"expired_at"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}

// KeyStore ...
type KeyStore interface {
	// Returns a key by given token
	FindKeyByToken(token string) (*Key, error)
}

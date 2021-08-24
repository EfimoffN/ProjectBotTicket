package sqlapi

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // here
)

type API struct {
	db *sqlx.DB
}

func NewAPI(db *sqlx.DB) *API {
	return &API{
		db: db,
	}
}

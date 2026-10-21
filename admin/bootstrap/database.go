package bootstrap

import (
	"database/sql"
	"os"
)

func OpenDatabase() (*sql.DB, error) {
	passoword, exsits := os.LookupEnv("LINKSPAN_DB_PASSWOR")
	if !exsits {
		return nil, fmt.
	}
}
package conexion

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var Session2 *sql.DB

func init() {
	Connect2()
}

func Connect2() {
	var err2 error

	uri := "postgres://postgres:jordan9595@@34.71.38.113:5432/turismo?sslmode=disable"

	Session2, err2 = sql.Open("postgres", uri)
	if err2 != nil {
		panic(err2)
	}

	if err2 != nil {
		time.Sleep(10000 * time.Millisecond)
		Connect()
	}
}

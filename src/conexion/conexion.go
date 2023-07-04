package conexion

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var Session *sql.DB

func init() {
	Connect()
}

func Connect() {
	var err error

	uri := "postgres://postgres:jordan9595@@34.71.38.113:5432/contactos?sslmode=disable"

	Session, err = sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}

	if err != nil {
		time.Sleep(10000 * time.Millisecond)
		Connect()
	}
}

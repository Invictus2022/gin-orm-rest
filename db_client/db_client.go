package dbclient

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypass"
	dbname   = "mydb"
)

var DBClient *sqlx.DB

func InitialiseDBConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname) // Creating the connection string

	db, err := sqlx.Connect("postgres", psqlInfo) // Opening a connection to our database
	if err != nil {
		panic(err.Error())
	}

	// It is vitally important that you call the Ping() method becuase the sql.Open() function call does not
	// ever create a connection to the database. Instead, it simply validates the arguments provided.

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db

	fmt.Println("Successfully connected!")
}

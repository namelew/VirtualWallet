package databases

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/namelew/VirtualWallet/internal/envoriment"
)

type Database struct {
	db *sql.DB
}

func New() *Database {
	return &Database{
		db: nil,
	}
}

func (d *Database) Connect() {
	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		envoriment.GetVar("DBHOST"),
		envoriment.GetVar("DBPORT"),
		envoriment.GetVar("DBUSER"),
		envoriment.GetVar("DBPASS"),
		envoriment.GetVar("DBNAME"),
		envoriment.GetVar("DBSSL"),
	)

	db, err := sql.Open("postgres", connectString)

	if err != nil {
		log.Fatal("unable to connect to database. ", err.Error())
	}

	d.db = db
}

func (d *Database) Disconnect() {
	if err := d.db.Close(); err != nil {
		log.Fatal("Unable to close connection with database. ", err.Error())
	}
}

func (d *Database) Migrate() {

}

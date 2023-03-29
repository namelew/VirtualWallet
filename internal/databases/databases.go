package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
	"github.com/namelew/VirtualWallet/internal/envoriment"
)

type Database struct {
	db *sql.DB
}

type Table interface {
	Get(db *sql.DB, id []uint64) error
	Add(db *sql.DB) error
	Update(db *sql.DB) error
	Remove(db *sql.DB) error
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
	data, err := os.ReadFile("./migrations/" + envoriment.GetVar("DBNAME") + ".up.sql")

	if err != nil {
		log.Panic("Unable to load migrate configs. ", err.Error())
	}

	sanitaze := func(s string) string {
		trash := []string{
			"\n", "\b", "\t", "\a", "\r", "\f", "\v",
		}

		for i := range trash {
			s = strings.ReplaceAll(s, trash[i], "")
		}

		return s
	}

	for _, command := range strings.Split(sanitaze(string(data)), ";") {
		_, err = d.db.Exec(command)

		if err != nil {
			log.Fatal("Unable to execute migration step ", command, ".", err.Error())
		}
	}
}

func (d *Database) Add(reg Table) error {
	return reg.Add(d.db)
}

func (d *Database) Update(reg Table) error {
	return reg.Update(d.db)
}

func (d *Database) Get(reg Table, id ...uint64) error {
	return reg.Get(d.db, id)
}

func (d *Database) Remove(reg Table) error {
	return reg.Remove(d.db)
}

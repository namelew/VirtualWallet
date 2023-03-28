package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
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
	driver, err := postgres.WithInstance(d.db, &postgres.Config{})

	if err != nil {
		log.Fatal("Unable to load migrate configs. ", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://./internal/databases/migrations", envoriment.GetVar("DBNAME"), driver)

	if err != nil {
		log.Fatal("Unable to migrate table changes. ", err.Error())
	}

	m.Up()
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

package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
	"github.com/namelew/VirtualWallet/internal/clients"
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

func (d *Database) AddClient(c clients.Client) {
	_, err := d.db.Exec("insert into clients(name,amount) values ($1, $2)", c.Name, c.Amount)

	if err != nil {
		log.Fatal("Unable to create new client. ", err.Error())
	}
}

func (d *Database) getClient(id uint64) *clients.Client {
	var client clients.Client

	err := d.db.QueryRow("select id,name,amount from clients where id = $1", id).Scan(&client.ID, &client.Name, &client.Amount)

	if err != nil {
		log.Fatal("Unable to bind client data. ", err.Error())
	}

	if client.ID == 0 {
		log.Fatalf("unable to find client %d on database", id)
	}

	return &client
}

func (d *Database) ClientTransfer(source uint64, target uint64, value float64) {
	sender := d.getClient(source)
	receiver := d.getClient(target)
}

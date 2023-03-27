package clients

import (
	"database/sql"
	"errors"
	"strconv"
)

type Client struct {
	ID     uint64
	Name   string
	Amount float64
	Lock   uint8
}

func (c *Client) TransferValidation(value float64) bool {
	return c.Amount >= value
}

func (c *Client) Add(d *sql.DB) error {
	_, err := d.Exec("insert into clients(name,amount,lock) values ($1, $2, $3)", c.Name, c.Amount, c.Lock)

	if err != nil {
		return errors.New("Unable to create new client. " + err.Error())
	}

	return nil
}

func (c *Client) Update(d *sql.DB) error {
	_, err := d.Exec("update clients set name=$2,amount=$3,lock=$4 where id=$1", c.ID, c.Name, c.Amount, c.Lock)

	if err != nil {
		return errors.New("unable to update client data. " + err.Error())
	}

	return nil
}

func (c *Client) Get(d *sql.DB, id uint64) error {
	err := d.QueryRow("select id,name,amount,lock from clients where id = $1", id).Scan(&c.ID, &c.Name, &c.Amount, &c.Lock)

	if err != nil {
		return errors.New("unable to bind client data. " + err.Error())
	}

	if c.ID == 0 {
		return errors.New("unable to find client " + strconv.FormatInt(int64(id), 10) + " on database")
	}

	return nil
}

func (c *Client) Remove(d *sql.DB) error {
	_, err := d.Exec("delete from clients where id=$1", c.ID)

	if err != nil {
		return errors.New("unable to delete client data. " + err.Error())
	}

	return nil
}

// func (c *Client) ClientTransfer(source uint64, target uint64, value float64) {
// 	sender := d.getClient(source)
// 	receiver := d.getClient(target)

// 	if sender.Lock != 0 || receiver.Lock != 0 {
// 		time.Sleep(time.Second)
// 		d.ClientTransfer(source, target, value)
// 		return
// 	}

// 	if !sender.TransferValidation(value) {
// 		log.Fatal("Can't execute tranfer, value granter then the client amount")
// 	}

// 	sender.Amount -= value
// 	receiver.Amount += value
// }

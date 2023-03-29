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
}

func (c *Client) AmountValidation(value float64) bool {
	return c.Amount >= value
}

func (c *Client) Add(d *sql.DB) error {
	_, err := d.Exec("insert into clients(name,amount) values ($1, $2)", c.Name, c.Amount)

	if err != nil {
		return errors.New("Unable to create new client. " + err.Error())
	}

	return nil
}

func (c *Client) Update(d *sql.DB) error {
	ret, err := d.Exec("update clients set name=$2,amount=$3 where id=$1", c.ID, c.Name, c.Amount)

	if err != nil {
		return errors.New("unable to update client data. " + err.Error())
	}

	rows, err := ret.RowsAffected()

	if err != nil || rows == 0 {
		return errors.New("unable to update client data. register not found")
	}

	return nil
}

func (c *Client) Get(d *sql.DB, id []uint64) error {
	err := d.QueryRow("select id,name,amount from clients where id = $1", id[0]).Scan(&c.ID, &c.Name, &c.Amount)

	if err != nil {
		return errors.New("unable to bind client data. " + err.Error())
	}

	if c.ID == 0 {
		return errors.New("unable to find client " + strconv.FormatInt(int64(id[0]), 10) + " on database")
	}

	return nil
}

func (c *Client) Remove(d *sql.DB) error {
	ret, err := d.Exec("delete from clients where id=$1", c.ID)

	if err != nil {
		return errors.New("unable to delete client data. " + err.Error())
	}

	rows, err := ret.RowsAffected()

	if err != nil || rows == 0 {
		return errors.New("unable to delete client data. register not found")
	}

	return nil
}

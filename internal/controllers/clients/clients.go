package clients

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/namelew/VirtualWallet/internal/clients"
	"github.com/namelew/VirtualWallet/internal/databases"
)

type Client struct {
	db *databases.Database
}

func New(d *databases.Database) *Client {
	return &Client{
		db: d,
	}
}

func (cl *Client) Get(c echo.Context) (interface{}, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return nil, echo.ErrBadRequest
	}
	var ret clients.Client

	err = cl.db.Get(&ret, uint64(id))

	if err != nil {
		return nil, echo.ErrNotFound
	}

	return ret.Amount, nil
}

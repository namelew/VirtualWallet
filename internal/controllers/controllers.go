package controllers

import (
	"github.com/labstack/echo"
	"github.com/namelew/VirtualWallet/internal/controllers/clients"
	"github.com/namelew/VirtualWallet/internal/controllers/transitions"
	"github.com/namelew/VirtualWallet/internal/databases"
)

type Controller struct {
	client     *clients.Client
	transition *transitions.Transition
}

func New(d *databases.Database) *Controller {
	return &Controller{
		client:     clients.New(d),
		transition: transitions.New(d),
	}
}

func (con *Controller) Get(c echo.Context) error {
	switch x := c.Request().RequestURI; {
	case x[:8] == "/clients":
		ret, err := con.client.Get(c)

		if err != nil {
			return err
		}

		return c.JSON(200, ret)
	}
	return echo.ErrBadRequest
}

func (con *Controller) Add(c echo.Context) error {
	switch x := c.Request().RequestURI; {
	case x[:12] == "/transitions":
		err := con.transition.Add(c)

		if err != nil {
			return err
		}

		return c.NoContent(200)
	}
	return echo.ErrBadRequest
}

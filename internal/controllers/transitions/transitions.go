package transitions

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/namelew/VirtualWallet/internal/clients"
	"github.com/namelew/VirtualWallet/internal/databases"
)

type Transition struct {
	db *databases.Database
}

func New(d *databases.Database) *Transition {
	return &Transition{
		db: d,
	}
}

func (t *Transition) Add(c echo.Context) error {
	source_id, err := strconv.Atoi(c.Param("source"))

	if err != nil {
		return echo.ErrBadRequest
	}

	target_id, err := strconv.Atoi(c.Param("target"))

	if err != nil {
		return echo.ErrBadRequest
	}

	value, err := strconv.ParseFloat(c.Param("val"), 64)

	if err != nil {
		return echo.ErrBadRequest
	}

	var sender, receiver clients.Client

	if err := t.db.Get(&sender, uint64(source_id)); err != nil {
		return echo.ErrNotFound
	}

	if err := t.db.Get(&receiver, uint64(target_id)); err != nil {
		return echo.ErrNotFound
	}

	if value > sender.Amount {
		return echo.ErrBadRequest
	}

	// precisa de uma função que checa se há uma operação não finalizada em execução com o mesmo target

	return nil
}

package transitions

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/namelew/VirtualWallet/internal/clients"
	"github.com/namelew/VirtualWallet/internal/databases"
	"github.com/namelew/VirtualWallet/internal/transations"
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

	scopy := sender

	new := transations.Transation{
		SenderID:   sender.ID,
		ReceiverID: receiver.ID,
		Amount:     value,
		Finished:   false,
		CreatedAt:  time.Now(),
	}

	ensure := func(d *databases.Database, nw databases.Table, n int) {
		err := d.Update(nw)
		for err != nil {
			seed := rand.NewSource(time.Now().UnixNano())
			random := rand.New(seed)
			time.Sleep(time.Duration(random.Intn(n)))
			err = d.Update(nw)
		}
	}

	if err := t.db.Add(&new); err != nil {
		return echo.ErrInternalServerError
	}

	sender.Amount -= value
	receiver.Amount += value

	if err := t.db.Update(&sender); err != nil {
		new.Finished = true
		go ensure(t.db, &new, 10)
		return echo.ErrInternalServerError
	}

	if err := t.db.Update(&receiver); err != nil {
		go ensure(t.db, &scopy, 10)
		new.Finished = true
		go ensure(t.db, &new, 10)
		return echo.ErrInternalServerError
	}

	new.Finished = true
	go ensure(t.db, &new, 10)

	return nil
}

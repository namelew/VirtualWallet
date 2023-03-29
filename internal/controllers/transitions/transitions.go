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

func (t *Transition) pack(s clients.Client, r clients.Client, tn *transations.Transation) {
	scp := s

	ensure := func(d *databases.Database, nw databases.Table, n int) {
		err := d.Update(nw)
		for err != nil {
			seed := rand.NewSource(time.Now().UnixNano())
			random := rand.New(seed)
			time.Sleep(time.Duration(random.Intn(n)) * time.Microsecond)
			err = d.Update(nw)
		}
	}

	if err := t.db.Add(tn); err != nil {
		seed := rand.NewSource(time.Now().UnixNano())
		random := rand.New(seed)
		time.Sleep(time.Duration(random.Intn(20)) * time.Microsecond)
		t.pack(s, r, tn)
		return
	}

	s.Amount -= tn.Amount
	r.Amount += tn.Amount

	if err := t.db.Update(&s); err != nil {
		tn.Finished = true
		go ensure(t.db, tn, 10)
		return
	}

	if err := t.db.Update(&r); err != nil {
		go ensure(t.db, &scp, 10)
		tn.Finished = true
		go ensure(t.db, tn, 10)
		return
	}

	tn.Finished = true
	go ensure(t.db, tn, 10)
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

	new := transations.Transation{
		SenderID:   sender.ID,
		ReceiverID: receiver.ID,
		Amount:     value,
		Finished:   false,
		CreatedAt:  time.Now(),
	}

	go t.pack(sender, receiver, &new)

	return nil
}

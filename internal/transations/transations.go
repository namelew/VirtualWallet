package transations

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type Transation struct {
	SenderID   uint64
	ReceiverID uint64
	Amount     float64
	CreatedAt  time.Time // layout: time.DateTime
	Finished   bool
}

func (t *Transation) Add(d *sql.DB) error {
	_, err := d.Exec("insert into transations(sender_id,receiver_id,amount,finished) values ($1, $2, $3, $4)",
		t.SenderID,
		t.ReceiverID,
		t.Amount,
		t.Finished,
	)

	if err != nil {
		return errors.New("Unable to create new transations. " + err.Error())
	}

	return nil
}

func (t *Transation) Update(d *sql.DB) error {
	_, err := d.Exec("update transations set amount=$4,finished=$5 where sender_id=$1 and receiver_id=$2 and created_at=$3",
		t.SenderID,
		t.ReceiverID,
		t.CreatedAt,
		t.Amount,
		t.Finished,
	)

	if err != nil {
		return errors.New("unable to update transation data. " + err.Error())
	}

	return nil
}

func (t *Transation) Get(d *sql.DB, id []uint64) error {
	err := d.QueryRow("select * from transations where sender_id=$1 and receiver_id=$2 and created_at=$3", id[0], id[1], t.CreatedAt).
		Scan(&t.SenderID, &t.ReceiverID, &t.Amount, &t.Finished, &t.CreatedAt)

	if err != nil {
		return errors.New("unable to bind transation data. " + err.Error())
	}

	if t.SenderID == 0 {
		return errors.New("unable to find transation (" + strconv.FormatInt(int64(id[0]), 10) + ", " + strconv.FormatInt(int64(id[2]), 10) + ", " + t.CreatedAt.GoString() + ") on database")
	}

	return nil
}

func (t *Transation) Remove(d *sql.DB) error {
	_, err := d.Exec("delete from transations where sender_id=$1 and receiver_id=$2 and created_at=$3", t.SenderID, t.ReceiverID, t.CreatedAt)

	if err != nil {
		return errors.New("unable to delete transation data. " + err.Error())
	}

	return nil
}

func Tranfer(c echo.Context) error {
	return nil
}

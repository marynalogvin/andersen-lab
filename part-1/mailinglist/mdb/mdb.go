package mdb

import (
	"database/sql"
	"log"
	"time"
)

type Subscriber struct {
	Id          int64
	Email       string
	ConfirmedAt time.Time
	OptOut      bool
}

func subscriberFromRow(rows *sql.Rows) (*Subscriber, error) {
	var (
		s           Subscriber
		confirmedAt int64
	)
	err := rows.Scan(&s.Id, &s.Email, &confirmedAt, &s.OptOut)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	s.ConfirmedAt = time.Unix(confirmedAt, 0)
	return &s, nil
}

func CreateSubscriber(db *sql.DB, email string) error {
	_, err := db.Exec(`
	INSERT INTO
	emails(email, confirmed_at, opt_out)
	VALUES(?, ?, false)`, email, time.Now().Unix())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetSubscriber(db *sql.DB, email string) (*Subscriber, error) {
	rows, err := db.Query(`
	SELECT id, email, confirmed_at, opt_out
	FROM emails
	WHERE email = ?`, email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		return subscriberFromRow(rows)
	}
	return nil, nil
}

func UpdateSubscriber(db *sql.DB, sub Subscriber) error {
	if sub.ConfirmedAt.IsZero() {
		sub.ConfirmedAt = time.Now()
	}
	_, err := db.Exec(`
	INSERT INTO emails(email, confirmed_at, opt_out)
	VALUES(?, ?, ?)
	ON CONFLICT(email) DO UPDATE SET
	confirmed_at=?,
	opt_out=?`, sub.Email, sub.ConfirmedAt, sub.OptOut, sub.ConfirmedAt.Unix(), sub.OptOut)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CancelSubscription(db *sql.DB, email string) error {
	_, err := db.Exec(`
	UPDATE emails
	SET opt_out=true
	WHERE email=?`, email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

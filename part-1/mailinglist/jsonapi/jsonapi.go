package jsonapi

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"mailinglist/mdb"
	"net/http"
)

func fromJson[T any](body io.Reader, target T) error {
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(body); err != nil {
		return err
	}
	return json.Unmarshal(buf.Bytes(), &target)
}

func returnJson[T any](w http.ResponseWriter, withData func() (T, error)) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	data, serverErr := withData()
	if serverErr != nil {
		w.WriteHeader(500)
		serverErrJson, err := json.Marshal(&serverErr)
		if err != nil {
			log.Print(err)
			return
		}
		w.Write(serverErrJson)
		return
	}
	dataJson, err := json.Marshal(&data)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}
	w.Write(dataJson)
}

func returnErr(w http.ResponseWriter, err error, code int) {
	returnJson(w, func() (interface{}, error) {
		w.WriteHeader(code)
		return err, nil
	})
}

func CreateSubscriber(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}
		var sub mdb.Subscriber
		if err := fromJson(r.Body, &sub); err != nil {
			log.Printf("Reading file error: %v", err)
			return
		}
		if err := mdb.CreateSubscriber(db, sub.Email); err != nil {
			returnErr(w, err, 400)
			return
		}
		returnJson(w, func() (interface{}, error) {
			log.Printf("JSON CreateSubscriber: %v\n", sub.Email)
			return mdb.GetSubscriber(db, sub.Email)
		})
	})
}

func GetSubscriber(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			return
		}
		var sub mdb.Subscriber
		if err := fromJson(r.Body, &sub); err != nil {
			log.Printf("Reading file error: %v", err)
			return
		}
		returnJson(w, func() (interface{}, error) {
			log.Printf("JSON GetSubscriber: %v\n", sub.Email)
			return mdb.GetSubscriber(db, sub.Email)
		})
	})
}

func UpdateSubscriber(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			return
		}
		var sub mdb.Subscriber
		if err := fromJson(r.Body, &sub); err != nil {
			log.Printf("Reading file error: %v", err)
			return
		}
		if err := mdb.UpdateSubscriber(db, sub); err != nil {
			returnErr(w, err, 400)
			return
		}
		returnJson(w, func() (interface{}, error) {
			log.Printf("JSON UpdateSubscriber: %v\n", sub.Email)
			return mdb.GetSubscriber(db, sub.Email)
		})
	})
}

func CancelSubscription(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}
		var sub mdb.Subscriber
		if err := fromJson(r.Body, &sub); err != nil {
			log.Printf("Reading file error: %v", err)
			return
		}
		if err := mdb.CancelSubscription(db, sub.Email); err != nil {
			returnErr(w, err, 400)
			return
		}
		returnJson(w, func() (interface{}, error) {
			log.Printf("JSON CancelSubscription: %v\n", sub.Email)
			return mdb.GetSubscriber(db, sub.Email)
		})
	})
}

func Serve(db *sql.DB, port string) {
	http.Handle("/email/create", CreateSubscriber(db))
	http.Handle("/email/get", GetSubscriber(db))
	http.Handle("/email/update", UpdateSubscriber(db))
	http.Handle("/email/cancel", CancelSubscription(db))
	log.Printf("JSON API serve listening on %v\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("JSON server error: %v", err)
	}
}

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64  `bun:"id,pk,autoincrement"`
	FristName     string `bun:"first_name"`
	LastName      string `bun:"last_name"`
	Address       string `bun:"address"`
	Age           int    `bun:"age"`
}

type UserReq struct {
	ID        int64  `json:"id"`
	FristName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Age       int    `json:"age"`
}

func main() {
	dsn := "postgresql://postgres:P123!@localhost:5443/data-liberation?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	r := mux.NewRouter()
	r.HandleFunc("/users", (func(w http.ResponseWriter, r *http.Request) {
		var req UserReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err := db.NewInsert().Model(&User{
			FristName: req.FristName,
			LastName:  req.LastName,
			Address:   req.Address,
			Age:       req.Age,
		}).Exec(r.Context())
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})).Methods("POST")

	r.HandleFunc("/users", (func(w http.ResponseWriter, r *http.Request) {
		var req UserReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err := db.NewUpdate().Model(&User{
			ID:        req.ID,
			FristName: req.FristName,
			LastName:  req.LastName,
			Address:   req.Address,
			Age:       req.Age,
		}).Exec(r.Context())
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})).Methods("PUT")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server started at:8000")
	log.Fatal(srv.ListenAndServe())
}

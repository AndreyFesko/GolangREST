package config

import (
	"database/sql"
	"fmt"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "employee"
	password = "employee"
	dbname   = "dbank"
)

var DB *sql.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		glog.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		glog.Fatal(err)
	}
}

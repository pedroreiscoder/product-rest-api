package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

type connection struct {
	Host     string
	Port     uint
	User     string
	Password string
	DBName   string
}

func (c *connection) ToString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)
}

func Init() {
	connInfo := connection{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "P3dr0@123",
		DBName:   "store",
	}

	var err error
	db, err = sql.Open("postgres", connInfo.ToString())

	if err != nil {
		log.Fatal(err)
	}
}

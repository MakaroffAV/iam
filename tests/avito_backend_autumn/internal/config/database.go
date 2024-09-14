package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type DB struct {
	host string
	port string
	user string
	pass string
	name string
}

func NewDB() DB {
	return DB{
		host: os.Getenv("POSTGRES_HOST"),
		port: os.Getenv("POSTGRES_PORT"),
		user: os.Getenv("POSTGRES_USERNAME"),
		pass: os.Getenv("POSTGRES_PASSWORD"),
		name: os.Getenv("POSTGRES_DATABASE"),
	}
}

func (d DB) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		d.host,
		d.port,
		d.user,
		d.pass,
		d.name,
	)
}

func (d DB) MustConn() *sql.DB {
	c, err := sql.Open("postgres", d.dsn())
	if err != nil {
		log.Fatalln(
			err, "invalid db configuration",
		)
	}
	if err := c.Ping(); err != nil {
		log.Fatalln(
			err, "invalid pq db connection",
		)
	}
	return c
}

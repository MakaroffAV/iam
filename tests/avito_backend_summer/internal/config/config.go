package config

import (
	"fmt"
	"os"
)

var DbTplString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Moscow"

var DbParams = map[string]string{
	"host": os.Getenv("PDB_HOST"),
	"port": os.Getenv("PDB_PORT"),
	"user": os.Getenv("PDB_USER"),
	"pass": os.Getenv("PDB_PASS"),
	"name": os.Getenv("PDB_NAME"),
}

var DbConnString = fmt.Sprintf(
	DbTplString,
	DbParams["host"],
	DbParams["port"],
	DbParams["user"],
	DbParams["pass"],
	DbParams["name"],
)

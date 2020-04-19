package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/malikov0216/createDataBase/config"
	"github.com/malikov0216/createDataBase/migration"
)

var db *sql.DB

func main() {
	var err error
	dbInfo := fmt.Sprintf("host=%s port=%s search_path=%s dbname=%s user=%s password=%s sslmode=disable", config.DbHost, config.DbPort, config.DbSchema, config.DbName, config.DbUser, config.DbPassword)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	_, err = db.Query(migration.CreateInteraction)
	if err != nil {
		log.Println("failed to run migration", err.Error())
		return
	}
}

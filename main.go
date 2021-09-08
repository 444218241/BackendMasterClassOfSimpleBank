package main

import (
	"database/sql"
	"fmt"
	"log"
	"techschool/samplebank/api"
	db "techschool/samplebank/db/sqlc"
	"techschool/samplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	// config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("Config err", err)
	}
	fmt.Println(config.DBDriver)

	// conn type is *DB
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("Connect err", err)
	}

	// store --- *db.Store
	store := db.NewStore(conn)

	// server -- struct : *Server{}
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start the server", err)
	}
}

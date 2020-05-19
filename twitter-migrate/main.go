package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"

	"github.com/brunopita/go-common/commonsys"
	"github.com/brunopita/twitter-example/twitter-pg/spg"
)

var db *sql.DB
var env *commonsys.Environment

func init() {
	var err error

	env = commonsys.GetEnvironment()
	host := env.GetOrDefault("POSTGRES_HOST", "192.168.0.153")
	port := env.GetOrDefault("POSTGRES_PORT", "15432")
	user := env.GetOrDefault("POSTGRES_USER", "twitter")
	password := env.GetOrDefault("POSTGRES_PASSWORD", "teste@123")
	dbname := env.GetOrDefault("POSTGRES_DATABASE", "twitter")

	db, err = spg.GetConnection(host, port, user, password, dbname)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	defer db.Close()

	file, err := ioutil.ReadFile("./scripts/twitter.sql")
	if err != nil {
		log.Fatalln(err)
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

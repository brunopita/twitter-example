package spg

import (
	"database/sql"
	"fmt"

	"github.com/brunopita/go-common/commonsys"
)

var env *commonsys.Environment
var host, port, user, password, dbname string

func init() {
	env = commonsys.GetEnvironment()
	host = env.GetOrDefault("PG_HOST", "locahost")
	port = env.GetOrDefault("PG_PORT", "5432")
	user = env.GetOrDefault("PG_USERNAME", "twitter")
	password = env.GetOrDefault("PG_PASSWORD", "teste@123")
	dbname = env.GetOrDefault("PG_DATABASE", "twitter")
}

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

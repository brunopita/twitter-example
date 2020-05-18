package spg

import (
	"database/sql"
	"fmt"

	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

func GetConnection(host, port, user, password, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := apmsql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

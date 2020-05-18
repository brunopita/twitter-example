package controller

import (
	"database/sql"

	"github.com/brunopita/go-common/commonsys"
	"github.com/brunopita/twitter-example/twitter-consumer/slog"
	"github.com/brunopita/twitter-example/twitter-pg/spg"
	"github.com/brunopita/twitter-example/twitter-pg/tdao"
	"github.com/sirupsen/logrus"
)

var db *sql.DB
var env *commonsys.Environment
var log *logrus.Entry

func init() {
	var err error

	env = commonsys.GetEnvironment()
	log = slog.Logger("twiitter-api", "controller")
	host := env.GetOrDefault("POSTGRES_HOST", "192.168.0.153")
	port := env.GetOrDefault("POSTGRES_PORT", "15432")
	user := env.GetOrDefault("POSTGRES_USER", "twitter")
	password := env.GetOrDefault("POSTGRES_PASSWORD", "teste@123")
	dbname := env.GetOrDefault("POSTGRES_DATABASE", "twitter")

	db, err = spg.GetConnection(host,port,user,password,dbname)
	if err != nil {
		log.Error(err)
	}
}


func TopFiveFollowController() []tdao.User {
	result, err := tdao.GetTopFiveUserFollowers(db)
	if err != nil {
		log.Error(err)
	}
	return result
}
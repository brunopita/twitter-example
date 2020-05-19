package controller

import (
	"context"
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

	db, err = spg.GetConnection(host, port, user, password, dbname)
	if err != nil {
		log.Error(err)
	}
}

func TopFiveFollowController(ctx context.Context) ([]tdao.User, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	result, err := tdao.GetTopFiveUserFollowers(tx)
	if err != nil {
		log.Error(err)
		return result, err
	}
	log.Info(result)
	return result, nil
}

func PostsForHourController(ctx context.Context) ([]tdao.QttyHourHashtag, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	result, err := tdao.GetQttyForHourByHashtag(tx)
	if err != nil {
		log.Error(err)
		return result, err
	}
	log.Info(result)
	return result, nil
}

func TotalPostHashtagByLocate(ctx context.Context) ([]tdao.QttyHashtagLocate, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	result, err := tdao.GetQttyPostForHashtagByLocate(tx)
	if err != nil {
		log.Error(err)
		return result, err
	}
	log.Info(result)
	return result, nil
}

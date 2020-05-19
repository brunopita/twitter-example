package main

import (
	"context"

	"github.com/brunopita/go-common/commonsys"
	"github.com/brunopita/twitter-example/twitter-consumer/client"
	"github.com/brunopita/twitter-example/twitter-consumer/search"
	"github.com/brunopita/twitter-example/twitter-consumer/slog"
	"github.com/brunopita/twitter-example/twitter-pg/spg"
	"github.com/brunopita/twitter-example/twitter-pg/tdao"
	"github.com/sirupsen/logrus"
)

var env *commonsys.Environment
var log *logrus.Entry
var host, port, user, password, dbname string

func init() {
	env = commonsys.GetEnvironment()
	log = slog.Logger("twiitter-consumer", "main")
	host = env.GetOrDefault("POSTGRES_HOST", "192.168.0.153")
	port = env.GetOrDefault("POSTGRES_PORT", "15432")
	user = env.GetOrDefault("POSTGRES_USER", "twitter")
	password = env.GetOrDefault("POSTGRES_PASSWORD", "teste@123")
	dbname = env.GetOrDefault("POSTGRES_DATABASE", "twitter")
}

func main() {

	client := client.GetTwitterClient()
	db, err := spg.GetConnection(host, port, user, password, dbname)
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	ctx := context.Background()

	for _, hashtag := range []string{"#openbanking", "#remediation", "#devops", "#sre", "#microservices", "#observability", "#oauth", "#metrics", "#logmonitoring", "#opentracing"} {
		search, resp, err := search.SearchByHashtag(hashtag, client)
		db.Begin()
		if resp.StatusCode != 200 {
			log.Errorf(resp.Status)
		}
		if err != nil {
			log.Errorln(err)
		}
		for _, val := range search.Statuses {
			var twitteUser tdao.User
			var tweet tdao.Tweet

			tx, err := db.BeginTx(ctx, nil)
			if err != nil {
				log.Error(err)
				continue
			}

			twitteUser = tdao.BuildUser(val.User)
			err = tdao.InsertUser(twitteUser, tx)
			if err != nil {
				log.Error("Rollback", err)
				tx.Rollback()
				continue
			}
			log.Info("Insert user: ", twitteUser)

			tweet = tdao.BuildTweet(&val, hashtag)
			err = tdao.InsertTweet(tweet, tx)
			if err != nil {
				log.Error("Rollback", err)
				tx.Rollback()
				continue
			}
			log.Info("Insert tweet: ", tweet)

			tx.Commit()
		}
	}

}

func InsertUser() {

}

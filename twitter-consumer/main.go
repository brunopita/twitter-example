package main

import (
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
	port = env.GetOrDefault("POSTGRES_PORT", "5432")
	user = env.GetOrDefault("POSTGRES_USER", "twitter")
	password = env.GetOrDefault("POSTGRES_PASSWORD", "twitter@123")
	dbname = env.GetOrDefault("POSTGRES_DATABASE", "twitter")
}

func main() {

	client := client.GetTwitterClient()

	for _, hashtag := range []string{"#openbanking", "#remediation", "#devops", "#sre", "#microservices", "#observability", "#oauth", "#metrics", "#logmonitoring", "#opentracing"} {
		search, resp, err := search.SearchByHashtag(hashtag, client)
		if resp.StatusCode != 200 {
			log.Errorf(resp.Status)
		}
		if err != nil {
			log.Errorln(err)
		}
		for _, val := range search.Statuses {
			var twitteUser tdao.User
			var tweet tdao.Tweet

			db, err := spg.GetConnection(host, port, user, password, dbname)
			if err != nil {
				log.Error(err)
			}
			defer db.Close()

			twitteUser = tdao.BuildUser(val.User)
			err = tdao.InsertUser(twitteUser, db)
			if err != nil {
				log.Error(err)
			}

			tweet = tdao.BuilTweet(val, hashtag)
			err = tdao.InsertTweet(tweet, db)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

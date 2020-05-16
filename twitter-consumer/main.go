package main

import (
	"github.com/brunopita/go-common/commonsys"
	"github.com/brunopita/twitter-consumer/client"
	"github.com/brunopita/twitter-consumer/search"
	"github.com/brunopita/twitter-consumer/slog"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/sirupsen/logrus"
)

var env *commonsys.Environment
var log *logrus.Entry

func init() {
	env = commonsys.GetEnvironment()
	log = slog.Logger("twiitter-consumer", "main")
}

func main() {

	client := client.GetTwitterClient()

	for _, val := range []string{"#openbanking", "#remediation", "#devops", "#sre", "#microservices", "#observability", "#oauth", "#metrics", "#logmonitoring", "#opentracing"} {
		search, resp, err := search.SearchByHashtag(val, client)
		if resp.StatusCode != 200 {
			log.Errorf(resp.Status)
		}
		if err != nil {
			log.Errorln(err)
		}
		for _, val := range search.Statuses {
			var twitte twitter.Tweet
			twitte = val
			// log.Info(twitte.User.FollowersCount)
			log.Info(twitte.CreatedAt)
			log.Info(twitte.Entities.Hashtags)
			// log.Info(twitte.ID)
			// log.Info(twitte.User.ID)
		}
	}
}

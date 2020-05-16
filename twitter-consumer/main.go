package main

import (
	"github.com/brunopita/go-common/commonsys"
	"github.com/brunopita/twitter-consumer/client"
	"github.com/brunopita/twitter-consumer/search"
	"github.com/brunopita/twitter-consumer/slog"
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

	search, resp, err := search.SearchByHashtag("java", client)
	if resp.StatusCode != 200 {
		log.Errorf(resp.Status)
	}
	if err != nil {
		log.Errorln(err)
	}
	if search != nil {
		log.Info(search)
	}
}

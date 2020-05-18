package main

import (
	"net/http"

	"github.com/brunopita/go-common/commonsys"
	"github.com/brunopita/twitter-example/twitter-api/routes"
	"github.com/brunopita/twitter-example/twitter-consumer/slog"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmhttp"
)
var log *logrus.Entry
var env *commonsys.Environment


func init() {
	log = slog.Logger("twiitter-api", "main")
	env = commonsys.GetEnvironment()
}

func main() {
	var err error
	router := http.NewServeMux()

	router.HandleFunc("/top-five-follow", routes.TopFiveFollow)
	// router.HandleFunc("/posts-for-hour")
	// router.HandleFunc("/total-post-by-hashtag-locate")

	err = http.ListenAndServe(":3000",apmhttp.Wrap(router))
	if err != nil {
		log.Error("Fish program, error: ", err )
	}
}
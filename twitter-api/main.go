package main

import (
	"net/http"

	"github.com/brunopita/go-common/commonsys"
	"github.com/brunopita/twitter-example/twitter-api/middleware"
	"github.com/brunopita/twitter-example/twitter-api/routes"
	"github.com/brunopita/twitter-example/twitter-consumer/slog"
	"github.com/go-chi/render"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgorilla"
	"go.elastic.co/apm/module/apmhttp"
)

var log *logrus.Entry
var env *commonsys.Environment

func init() {
	log = slog.Logger("twiitter-api", "main")
	env = commonsys.GetEnvironment()
}

func main() {
	router := mux.NewRouter()
	apmgorilla.Instrument(router)

	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(middleware.ValidContentType)

	router.HandleFunc("/top-five-follow", routes.TopFiveFollow)
	router.HandleFunc("/posts-for-hour", routes.PostsForHour)
	router.HandleFunc("/total-post-hashtag-by-locate", routes.TotalPostHashtagByLocate)
	err := http.ListenAndServe(":3000", apmhttp.Wrap(router))
	if err != nil {
		log.Error("Fish program, error: ", err)
	}
}

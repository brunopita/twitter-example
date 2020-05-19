module github.com/brunopita/twitter-example/twitter-api

go 1.14

require (
	github.com/brunopita/go-common v0.0.0-20181128165638-9bc7e4da8eb3
	github.com/brunopita/twitter-example/twitter-consumer v0.0.0-20200517204602-051e6806e388
	github.com/brunopita/twitter-example/twitter-pg v0.0.0-20200519024706-9a4889fe2430
	github.com/go-chi/render v1.0.1
	github.com/gorilla/mux v1.7.4
	github.com/sirupsen/logrus v1.6.0
	go.elastic.co/apm/module/apmgorilla v1.8.0
	go.elastic.co/apm/module/apmhttp v1.8.0
	go.elastic.co/apm/module/apmsql v1.8.0
)

package slog

import (
	"fmt"
	"net"
	"time"

	"github.com/brunopita/go-common/commonsys"
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

var env *commonsys.Environment
var err error
var conn net.Conn

func init() {
	env = commonsys.GetEnvironment()
}

func Logger(application string, class string) *logrus.Entry {
	var hook *logrustash.Hook
	url := env.GetOrDefault("LOGSTASH_URL", "192.168.15.21:5000")
	log := logrus.New()

	for {

		hook, err = logrustash.NewHook("tcp", url, application)
		if err != nil {
			fmt.Println("Erro ao buscar conexão com logstash, tentando novamente em 5 s")
			time.Sleep(5 * time.Second)
			continue
		} else {
			break
		}
	}

	log.AddHook(hook)
	ctx := log.WithFields(logrus.Fields{
		"method": class,
	})
	return ctx
}

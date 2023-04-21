package server

import (
	"testing"

	"github.com/galaxy-toolkit/server/config"
	"github.com/galaxy-toolkit/server/log"
)

func TestRouter(t *testing.T) {
	var conf config.Config
	err := config.Load("", &conf)
	if err != nil {
		t.Fatal(err)
	}

	router := NewRouter()
	router.Use(
		NewLimiterHandler(conf.Server),
		NewLoggerHandler(conf.Server, log.Writer(conf.Log)),
	)

	if err := router.Listen(conf.Server.Host + ":" + conf.Server.Port); err != nil {
		t.Fatal(err)
	}
}

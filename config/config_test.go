package config

import (
	"testing"
	"time"
)

func TestLoadConfig(t *testing.T) {
	var config Config

	err := Load("", &config)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", config)
}

func TestLoadAndWatchConfig(t *testing.T) {
	var config Config

	err := LoadAndWatch("", &config)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", config)

	time.Sleep(time.Second * 20)
}

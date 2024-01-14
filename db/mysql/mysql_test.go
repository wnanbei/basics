package mysql

import (
	"testing"

	"github.com/galaxy-toolkit/server/config"
)

func TestMySQLGenerator(t *testing.T) {
	var conf config.Config
	err := config.Load("", &conf)
	if err != nil {
		t.Fatal(err)
	}

	generator, err := NewGenerator(conf.Database.MySQL, GeneratorConfig{ModelPath: "./model"})
	if err != nil {
		t.Fatal(err)
	}

	generator.GenerateModelAs("ip", "IP")
	generator.Execute()
}

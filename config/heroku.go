// +build heroku

package config

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gorelic"
	"os"
)

const APPNAME = "GOIGNEOUS"

func IsHeroku() bool {
	return true
}

func PostgresArgs() string {
	return os.Getenv("DATABASE_URL")
}

func Url() string {
	return os.Getenv("URL")
}

func Initialize(m *martini.ClassicMartini) {
	fmt.Println("Running in production environment")

	NEW_RELIC_LICENSE_KEY := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if len(NEW_RELIC_LICENSE_KEY) > 0 {
		gorelic.InitNewrelicAgent(NEW_RELIC_LICENSE_KEY, APPNAME, true)
		m.Use(gorelic.Handler)
	}
}

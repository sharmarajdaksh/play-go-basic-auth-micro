package config

import (
	"github.com/kkyr/fig"
)

type config struct {
	Global struct {
		ListenPort string `fig:"listenPort" default:"9091"`
	}
	Security struct {
		SaltSize int `fig:"saltSize" default:"16"`
	}
	Database struct {
		Postgres struct {
			PostgresHost           string `fig:"postgresHost" default:"127.0.0.1"`
			PostgresConnectionPort string `fig:"postgresConnectionPort" default:"9920"`
			PostgresUsername       string `fig:"postgresUsername" default:"postgres"`
			PostgresPassword       string `fig:"postgresPassword" default:"postgres"`
			PostgresDatabaseName   string `fig:"postgresDatabaseName" default:"postgres"`
		}
	}
}

// C represents a global config object
var C config

// LoadConfig loads up the global config struct from file on startup
func LoadConfig() error {
	return fig.Load(&C,
		fig.File("./config/config.yaml"),
	)
}

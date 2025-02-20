package config

import (
	"fmt"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Logger struct {
		Formatter logrus.Formatter
	}
	Application struct {
		Port           string
		Name           string
		AllowedOrigins []string
	}
	PostgreSQL struct {
		Driver             string
		Host               string
		Port               string
		Username           string
		Password           string
		Database           string
		DSN                string
		MaxOpenConnections int
		MaxIdleConnections int
	}
}

func (cfg *Config) postgreSQL() {
	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	username := os.Getenv("POSTGRESQL_USERNAME")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	database := os.Getenv("POSTGRESQL_DATABASE")
	sslmode := os.Getenv("POSTGRESQL_SSLMODE")
	maxOpenConnections, _ := strconv.Atoi(os.Getenv("POSTGRESQL_MAX_OPEN_CONNECTIONS"))
	maxIdleConnections, _ := strconv.Atoi(os.Getenv("POSTGRESQL_MAX_IDLE_CONNECTIONS"))

	connVal := url.Values{}
	connVal.Add("sslmode", sslmode)
	connVal.Add("TimeZone", "Asia/Jakarta")

	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
	dsn := fmt.Sprintf("%s?%s", dataSource, connVal.Encode())

	cfg.PostgreSQL.Driver = "postgres"
	cfg.PostgreSQL.Host = host
	cfg.PostgreSQL.Port = port
	cfg.PostgreSQL.Username = username
	cfg.PostgreSQL.DSN = dsn
	cfg.PostgreSQL.Password = password
	cfg.PostgreSQL.Database = database
	cfg.PostgreSQL.MaxOpenConnections = maxOpenConnections
	cfg.PostgreSQL.MaxIdleConnections = maxIdleConnections
}

func (cfg *Config) logFormatter() {
	formatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			filename := fmt.Sprintf("%s:%d", f.File, f.Line)
			return funcname, filename
		},
	}

	cfg.Logger.Formatter = formatter
}

func (cfg *Config) app() {
	appName := os.Getenv("APP_NAME")
	port := os.Getenv("PORT")

	rawAllowedOrigins := strings.Trim(os.Getenv("ALLOWED_ORIGINS"), " ")

	allowedOrigins := make([]string, 0)
	if rawAllowedOrigins == "" {
		allowedOrigins = append(allowedOrigins, "*")
	} else {
		allowedOrigins = strings.Split(rawAllowedOrigins, ",")
	}

	cfg.Application.Port = port
	cfg.Application.Name = appName
	cfg.Application.AllowedOrigins = allowedOrigins
}

func Load() *Config {
	cfg := new(Config)
	cfg.app()
	cfg.logFormatter()
	cfg.postgreSQL()
	return cfg
}

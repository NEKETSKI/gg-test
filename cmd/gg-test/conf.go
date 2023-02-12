package main

import (
	"time"

	flag "github.com/spf13/pflag"
)

const (
	defaultLogTimeKey    = "time"
	defaultLogLevel      = "debug"
	defaultPgHost        = "localhost"
	defaultPgPort        = 5432
	defaultPgUsername    = "postgres"
	defaultPgPassword    = "postgres"
	defaultPgDatabase    = "postgres"
	defaultCheckInterval = time.Minute
)

var (
	_ = flag.String("log.time.key", defaultLogTimeKey, "Word to determine timestamp in log message")
	_ = flag.String("log.level", defaultLogLevel, "Logging level (info or debug)")
	_ = flag.String("pg.host", defaultPgHost, "PostgreSQL host address")
	_ = flag.Int("pg.port", defaultPgPort, "PostgreSQL port")
	_ = flag.String("pg.username", defaultPgUsername, "PostgreSQL username")
	_ = flag.String("pg.password", defaultPgPassword, "PostgreSQL password")
	_ = flag.String("pg.database", defaultPgDatabase, "PostgreSQL database name")
	_ = flag.Duration("check.interval", defaultCheckInterval,
		"The interval with which the request for endpoints will be executed")
)

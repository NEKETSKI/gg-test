package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/NEKETSKY/gg-test/internal/app"
	"github.com/NEKETSKY/gg-test/internal/repository/postgres"
	"github.com/NEKETSKY/gg-test/pkg/logger/zap"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// The version value can be changed at build time using the -ldflags Go build flag.
var version = "dev"

func init() {
	fmt.Println("Service version: ", version)
	flag.Parse()
	if err := viper.BindPFlags(flag.CommandLine); err != nil {
		log.Fatalf("failed to bind command line flags: %v", err)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func main() {
	loggerConfig, err := zap.NewConfig(viper.GetString("log.level"), viper.GetString("log.time.key"))
	if err != nil {
		log.Fatalf("failed to create logger config: %v", err)
	}
	logger, err := zap.InitLogger(&loggerConfig)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	repo, err := postgres.Init(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		viper.GetString("pg.username"), viper.GetString("pg.password"), viper.GetString("pg.host"),
		viper.GetInt("pg.port"), viper.GetString("pg.database")))
	if err != nil {
		logger.Fatalf("failed to initialize repository connection: %v", err)
	}
	defer func() {
		if err := repo.Close(); err != nil {
			logger.Fatalf("failed to close database connection: %v", err)
		}
	}()

	ctx, cancelFunc := context.WithCancel(context.Background())
	sigtermChan := make(chan os.Signal, 1)
	signal.Notify(sigtermChan, syscall.SIGTERM, syscall.SIGINT)

	service := app.New(logger, repo)
	go service.Start(ctx, viper.GetDuration("check.interval"))

	logger.Infof("Received signal: %v", <-sigtermChan)
	cancelFunc()
}

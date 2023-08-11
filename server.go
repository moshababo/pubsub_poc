package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"pubsub_poc/common"
	"pubsub_poc/config"
	"pubsub_poc/service"
)

func runServer() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cfg := config.DefaultConfig()
	logger := common.NewLogger(common.LogLevelInfo, "SERVER")
	logger.Info("starting")

	srv := service.New(cfg, logger)
	if err := srv.Start(ctx); err != nil {
		return fmt.Errorf("failed to start: %w", err)
	}

	return nil
}

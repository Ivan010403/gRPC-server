package main

import (
	"gRPCserver/internal/app"
	"gRPCserver/internal/config"
	"log/slog"
	"os"
)

func main() {
	logger := createLogger()
	if logger == nil {
		panic("logger failed to create")
	}

	logger.Info("logger has been created successfully")

	cfg, err := config.ReadConfig()
	if err != nil {
		logger.Error("failed to read config")
		return
	}

	logger.Info("config has been read successfully")

	application := app.NewApp(logger, cfg.GRPC_server.Port, cfg.GRPC_server.MaxReadWriteConn, cfg.GRPC_server.MaxCheckConn)

	application.GRPCsrv.MustRun()
}

// TODO: add switch with env and different levels of logging
func createLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return logger
}
package main

import (
	"context"
	"music_info/internal/config"
	"music_info/internal/server"
)

// @title Music
// @version 1.0
// @description This is a sample Music server.

// @host localhost:8080
// @BasePath /api/v1
func main() {
	ctx := context.Background()

	// config
	cfg := config.New()

	// envorinments [db and etc...]
	env := server.NewEnv(ctx, cfg)
	defer env.Stop()

	// server
	srv := server.NewServer(env, cfg)
	// server start
	srv.Run(ctx, cfg)
	defer srv.Stop(ctx)
}

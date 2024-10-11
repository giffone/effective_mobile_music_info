package main

import (
	"context"
	"music_info/internal/config"
	"music_info/internal/server"
)

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

package server

import (
	"context"
	"log"
	"music_info/internal/config"
	"music_info/internal/repo/postgres"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	pool *pgxpool.Pool
}

func NewEnv(ctx context.Context, cfg *config.Config) *Env {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return &Env{
		pool: postgres.New(ctx, cfg),
	}
}

func (e *Env) Stop() {
	e.pool.Close()
	log.Println("envorinments stopped")
}

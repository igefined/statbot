package main

import (
	"context"
	"fmt"
	"log"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/igilgyrg/statbot/api/bots/telegram"
	"github.com/igilgyrg/statbot/api/clients/coin"
	"github.com/igilgyrg/statbot/internal/currency"
	"github.com/igilgyrg/statbot/internal/storage/supabase"
	"github.com/igilgyrg/statbot/schema"
)

type DBConfig struct {
	Port     int    `config:"DB_PORT"`
	Host     string `config:"DB_HOST,required"`
	Username string `config:"DB_USER,required"`
	Password string `config:"DB_PASSWORD,required"`
	DBName   string `config:"DB_NAME,required"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := DBConfig{
		Port: 5432,
	}

	if err := confita.NewLoader(env.NewBackend()).Load(ctx, &cfg); err != nil {
		log.Fatalf("sendsay: failed to load config: %s", err)
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	schema.Migrate(&schema.DB, url)
	builder := schema.New(ctx, url)

	client := coin.NewClient()

	repository := supabase.NewRepository(builder)

	service := currency.NewService(repository, client)

	telegramClient := telegram.NewBot(service)

	telegramClient.Run(ctx)
}

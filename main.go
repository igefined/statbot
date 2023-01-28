package main

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type DBConfig struct {
	Port     int    `config:"DB_PORT"`
	Host     string `config:"DB_HOST,required"`
	Username string `config:"DB_USER,required"`
	Password string `config:"DB_PASSWORD,required"`
	DBName   string `config:"DB_NAME,required"`
}

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

func NewMux(lc fx.Lifecycle, logger *zap.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Sync()
			return server.Shutdown(ctx)
		},
	})

	return mux
}

func Register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/5fb2054478353fd8d514056d1745b3a9eef066deadda4b90967af7ca65ce6505", h)
}

var Module = fx.Provide(
	NewHandler,
)

var _commonComponents = fx.Options(
	Module,
)

func NewHandler() (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dat, err := os.ReadFile("./download.jpeg")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.Write(dat)
	}), nil
}

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	application := fx.New(
		fx.Provide(
			NewLogger,
			NewMux,
		),
		_commonComponents,
		fx.Invoke(Register),
	)
	application.Run()

	//cfg := DBConfig{
	//	Port: 5432,
	//}
	//
	//if err := confita.NewLoader(env.NewBackend()).Load(ctx, &cfg); err != nil {
	//	log.Fatalf("sendsay: failed to load config: %s", err)
	//}
	//
	//url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	//schema.Migrate(&schema.DB, url)
	//builder := schema.New(ctx, url)
	//
	//client := coin.NewClient()
	//
	//repository := supabase.NewRepository(builder)
	//
	//service := currency.NewService(repository, client)
	//
	//telegramClient := telegram.NewBot(service)
	//
	//telegramClient.Run(ctx)
}

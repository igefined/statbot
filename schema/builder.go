package schema

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type (
	QBuilder struct {
		querier *pgx.Conn
	}
)

func New(ctx context.Context, url string) *QBuilder {
	target, err := newInstance(ctx, url)
	if err != nil {
		log.Fatal(err)

		return nil
	}

	return target
}

func newInstance(ctx context.Context, URL string) (*QBuilder, error) {
	var err error

	conn, err := pgx.Connect(ctx, URL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &QBuilder{querier: conn}, nil
}

func (q QBuilder) Querier() *pgx.Conn {
	return q.querier
}

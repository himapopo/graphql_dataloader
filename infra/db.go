package infra

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/qustavo/sqlhooks/v2"
)

var DB *sqlx.DB

type Hooks struct{}

// Before hook will print the query with it's args and return the context with the timestamp
func (h *Hooks) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	fmt.Printf("> %s %q\n", query, args)
	return context.WithValue(ctx, "begin", time.Now()), nil
}

// After hook will get the timestamp registered on the Before hook and print the elapsed time
func (h *Hooks) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	begin := ctx.Value("begin").(time.Time)
	fmt.Printf(". took: %s\n", time.Since(begin))
	return ctx, nil
}

func Init() {
	var err error
	sql.Register("psqlWithHooks", sqlhooks.Wrap(pq.Driver{}, &Hooks{}))

	// Connect to the registered wrapped driver
	DB, err = sqlx.Connect("psqlWithHooks", "user=yokokawataiki dbname=sample_db password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

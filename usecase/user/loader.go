package user

import (
	"context"
	model2 "graphql_dataloader/domain/model"
	"graphql_dataloader/graph/model"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/jmoiron/sqlx"
)

type ctxKey string

const (
	LoadersKey = ctxKey("dataloaders")
)

type UserReader interface {
	FetchUsers(context.Context, []int64) []*dataloader.Result[*model.User]
}

type userReader struct {
	db *sqlx.DB
}

func NewUserReader(db *sqlx.DB) *userReader {
	return &userReader{
		db: db,
	}
}

type UserLoader struct {
	FetchUsers *dataloader.Loader[int64, *model.User]
}

func NewUserLoader(reader UserReader) UserLoader {
	return UserLoader{
		FetchUsers: dataloader.NewBatchedLoader(reader.FetchUsers),
	}
}

func (u *userReader) FetchUsers(ctx context.Context, keys []int64) []*dataloader.Result[*model.User] {
	var err error
	res := make([]*dataloader.Result[*model.User], len(keys))
	sql := `SELECT * FROM users WHERE id IN (?)`
	sql, params, err := sqlx.In(sql, keys)
	if err != nil {
		res = append(res, &dataloader.Result[*model.User]{
			Error: err,
		})
		return res
	}
	sql = sqlx.Rebind(sqlx.DOLLAR, sql)
	users := []*model2.User{}
	if err := u.db.Select(&users, sql, params...); err != nil {
		res = append(res, &dataloader.Result[*model.User]{
			Error: err,
		})
		return res
	}
	userById := map[int64]*model2.User{}
	for _, u := range users {
		userById[u.ID] = u
	}

	for i, k := range keys {
		u := userById[k]
		res[i] = &dataloader.Result[*model.User]{
			Data: &model.User{
				ID:   u.ID,
				Name: u.Name,
			},
		}
	}

	return res
}

func For(ctx context.Context) UserLoader {
	return ctx.Value(LoadersKey).(UserLoader)
}

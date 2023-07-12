package repo

import (
	"context"
	"pomfret.cn/project-project/internal/database"
)

type ProjectAuthNodeRepo interface {
	FindNodeStringList(ctx context.Context, authId int64) ([]string, error)
	DeleteByAuthId(background context.Context, conn database.DbConn, id int64) error
	Save(background context.Context, conn database.DbConn, authId int64, nodes []string) error
}

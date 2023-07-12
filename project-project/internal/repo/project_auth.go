package repo

import (
	"context"
	"pomfret.cn/project-project/internal/data"
)

type ProjectAuthRepo interface {
	FindAuthList(ctx context.Context, code int64) (list []*data.ProjectAuth, err error)
	FindAuthListPage(ctx context.Context, code int64, page int64, size int64) (list []*data.ProjectAuth, total int64, err error)
}

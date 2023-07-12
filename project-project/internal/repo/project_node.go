package repo

import (
	"context"
	"pomfret.cn/project-project/internal/data"
)

type ProjectNodeRepo interface {
	FindAll(ctx context.Context) (list []*data.ProjectNode, err error)
}

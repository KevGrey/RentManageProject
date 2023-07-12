package repo

import (
	"context"
	"pomfret.cn/project-project/internal/data"
)

type MenuRepo interface {
	FindMenus(ctx context.Context) ([]*data.ProjectMenu, error)
}

package domain

import (
	"golang.org/x/net/context"
	"pomfret.cn/project-project/internal/dao"
	"pomfret.cn/project-project/internal/data"
	"pomfret.cn/project-project/internal/repo"
	"pomfret.cn/project-project/pkg/model"
	"pomfret.cn/project_common/errs"
	"time"
)

type MenuDomain struct {
	menuRepo repo.MenuRepo
}

func (d *MenuDomain) MenuTreeList() ([]*data.ProjectMenuChild, *errs.BError) {
	c, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	menus, err := d.menuRepo.FindMenus(c)
	if err != nil {
		return nil, model.DBError
	}
	menuChildren := data.CovertChild(menus)
	return menuChildren, nil

}

func NewMenuDomain() *MenuDomain {
	return &MenuDomain{
		menuRepo: dao.NewMenuDao(),
	}
}

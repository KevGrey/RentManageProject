package menu_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	menu "pomfret.cn/project-grpc/menu"
	"pomfret.cn/project-project/internal/dao"
	"pomfret.cn/project-project/internal/database/tran"
	"pomfret.cn/project-project/internal/domain"
	"pomfret.cn/project-project/internal/repo"
	"pomfret.cn/project_common/errs"
)

type MenuService struct {
	menu.UnimplementedMenuServiceServer
	cache       repo.Cache
	transaction tran.Transaction
	menuDomain  *domain.MenuDomain
}

func New() *MenuService {
	return &MenuService{
		cache:       dao.Rc,
		transaction: dao.NewTransaction(),
		menuDomain:  domain.NewMenuDomain(),
	}
}

func (m *MenuService) MenuList(context.Context, *menu.MenuReqMessage) (*menu.MenuResponseMessage, error) {
	treeList, err := m.menuDomain.MenuTreeList()
	if err != nil {
		zap.L().Error("MenuList error", zap.Error(err))
		return nil, errs.GrpcError(err)
	}
	var list []*menu.MenuMessage
	copier.Copy(&list, treeList)
	return &menu.MenuResponseMessage{List: list}, nil
}

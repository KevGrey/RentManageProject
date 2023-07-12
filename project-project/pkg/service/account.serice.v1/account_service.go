package account_serice_v1

import (
	"context"
	"github.com/jinzhu/copier"
	account "pomfret.cn/project-grpc/account"
	"pomfret.cn/project-project/internal/dao"
	"pomfret.cn/project-project/internal/database/tran"
	"pomfret.cn/project-project/internal/domain"
	"pomfret.cn/project-project/internal/repo"
	"pomfret.cn/project_common/encrypts"
	"pomfret.cn/project_common/errs"
)

type AccountService struct {
	account.UnimplementedAccountServiceServer
	cache             repo.Cache
	transaction       tran.Transaction
	accountDomain     *domain.AccountDomain
	projectAuthDomain *domain.ProjectAuthDomain
}

func New() *AccountService {
	return &AccountService{
		cache:             dao.Rc,
		transaction:       dao.NewTransaction(),
		accountDomain:     domain.NewAccountDomain(),
		projectAuthDomain: domain.NewProjectAuthDomain(),
	}
}

func (a *AccountService) Account(ctx context.Context, msg *account.AccountReqMessage) (*account.AccountResponse, error) {
	//1。去account表查询account
	//2。去auth表查询authList
	accountList, total, err := a.accountDomain.AccountList(
		msg.OrganizationCode,
		msg.MemberId,
		msg.Page,
		msg.PageSize,
		msg.DepartmentCode,
		msg.SearchType)
	if err != nil {
		return &account.AccountResponse{}, errs.GrpcError(err)
	}
	authList, err := a.projectAuthDomain.AuthList(encrypts.DecryptNoErr(msg.OrganizationCode))
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var maList []*account.MemberAccount
	copier.Copy(&maList, accountList)
	var prList []*account.ProjectAuth
	copier.Copy(&prList, authList)
	return &account.AccountResponse{
		AccountList: maList,
		AuthList:    prList,
		Total:       total,
	}, nil
}
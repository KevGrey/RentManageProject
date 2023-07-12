package department_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	department "pomfret.cn/project-grpc/department"
	"pomfret.cn/project-project/internal/dao"
	"pomfret.cn/project-project/internal/database/tran"
	"pomfret.cn/project-project/internal/domain"
	"pomfret.cn/project-project/internal/repo"
	"pomfret.cn/project_common/encrypts"
	"pomfret.cn/project_common/errs"
)

type DepartmentService struct {
	department.UnimplementedDepartmentServiceServer
	cache            repo.Cache
	transaction      tran.Transaction
	departmentDomain *domain.DepartmentDomain
}

func New() *DepartmentService {
	return &DepartmentService{
		cache:            dao.Rc,
		transaction:      dao.NewTransaction(),
		departmentDomain: domain.NewDepartmentDomain(),
	}
}
func (d *DepartmentService) List(ctx context.Context, msg *department.DepartmentReqMessage) (*department.ListDepartmentMessage, error) {
	organizationCode := encrypts.DecryptNoErr(msg.OrganizationCode)
	var parentDepartmentCode int64
	if msg.ParentDepartmentCode != "" {
		parentDepartmentCode = encrypts.DecryptNoErr(msg.ParentDepartmentCode)
	}
	dps, total, err := d.departmentDomain.List(
		organizationCode,
		parentDepartmentCode,
		msg.Page,
		msg.PageSize)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var list []*department.DepartmentMessage
	//config.SendLog(kk.Info("list", "DepartmentService.list", kk.FieldMap{
	//	"organizationCode":     organizationCode,
	//	"parentDepartmentCode": parentDepartmentCode,
	//	"page":                 msg.Page,
	//}))
	copier.Copy(&list, dps)
	return &department.ListDepartmentMessage{List: list, Total: total}, nil
}
func (d *DepartmentService) Save(ctx context.Context, msg *department.DepartmentReqMessage) (*department.DepartmentMessage, error) {
	organizationCode := encrypts.DecryptNoErr(msg.OrganizationCode)
	var departmentCode int64
	if msg.DepartmentCode != "" {
		departmentCode = encrypts.DecryptNoErr(msg.DepartmentCode)
	}
	var parentDepartmentCode int64
	if msg.ParentDepartmentCode != "" {
		parentDepartmentCode = encrypts.DecryptNoErr(msg.ParentDepartmentCode)
	}
	dp, err := d.departmentDomain.Save(
		organizationCode,
		departmentCode,
		parentDepartmentCode,
		msg.Name)
	if err != nil {
		return &department.DepartmentMessage{}, errs.GrpcError(err)
	}
	var res = &department.DepartmentMessage{}
	copier.Copy(res, dp)
	return res, nil
}

func (d *DepartmentService) Read(ctx context.Context, msg *department.DepartmentReqMessage) (*department.DepartmentMessage, error) {
	//organizationCode := encrypts.DecryptNoErr(msg.OrganizationCode)
	departmentCode := encrypts.DecryptNoErr(msg.DepartmentCode)
	dp, err := d.departmentDomain.FindDepartmentById(departmentCode)
	if err != nil {
		return &department.DepartmentMessage{}, errs.GrpcError(err)
	}
	var res = &department.DepartmentMessage{}
	copier.Copy(res, dp.ToDisplay())
	return res, nil
}

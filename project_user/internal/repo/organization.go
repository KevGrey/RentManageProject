package repo

import (
	"context"
	"pomfret.cn/project-user/internal/data/organization"
	"pomfret.cn/project-user/internal/database"
)

type OrganizationRepo interface {
	SaveOrganization(conn database.DbConn, ctx context.Context, org *organization.Organization) error
	FindOrganizationByMemId(ctx context.Context, memId int64) ([]*organization.Organization, error)
}

package repo

import (
	"context"
	"pomfret.cn/project-user/internal/data/member"
	"pomfret.cn/project-user/internal/database"
)

type MemberRepo interface {
	GetMemberByEmail(ctx context.Context, email string) (bool, error)
	GetMemberByAccount(ctx context.Context, account string) (bool, error)
	GetMemberByMobile(ctx context.Context, mobile string) (bool, error)
	SaveMember(conn database.DbConn, ctx context.Context, mem *member.Member) error
	FindMember(ctx context.Context, account string, password string) (mem *member.Member, err error)
	FindMemberById(background context.Context, id int64) (mem *member.Member, err error)
	FindMemberByIds(background context.Context, ids []int64) (list []*member.Member, err error)
}

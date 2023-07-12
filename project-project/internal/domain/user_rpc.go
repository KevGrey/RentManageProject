package domain

import (
	"context"
	login "pomfret.cn/project-grpc/user/login"
	"pomfret.cn/project-project/internal/rpc"
	"time"
)

type UserRpcDomain struct {
	lc login.LoginServiceClient
}

func NewUserRpcDomain() *UserRpcDomain {
	return &UserRpcDomain{
		lc: rpc.LoginServiceClient,
	}

}

func (d *UserRpcDomain) MemberList(mIdList []int64) ([]*login.MemberMessage, map[int64]*login.MemberMessage, error) {
	c, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	messageList, err := d.lc.FindMemInfoByIds(c, &login.UserMessage{MIds: mIdList})
	mMap := make(map[int64]*login.MemberMessage)
	for _, v := range messageList.List {
		mMap[v.Id] = v
	}
	return messageList.List, mMap, err
}

func (d *UserRpcDomain) MemberInfo(ctx context.Context, memberCode int64) (*login.MemberMessage, error) {
	c, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	memberMessage, err := d.lc.FindMemInfoById(c, &login.UserMessage{
		MemId: memberCode,
	})
	return memberMessage, err
}

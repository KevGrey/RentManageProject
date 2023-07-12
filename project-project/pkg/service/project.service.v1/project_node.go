package project_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	project "pomfret.cn/project-grpc/project"
	"pomfret.cn/project_common/errs"
)

func (p *ProjectService) NodeList(c context.Context, msg *project.ProjectRpcMessage) (*project.ProjectNodeResponseMessage, error) {
	list, err := p.nodeDomain.TreeList()
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var nodes []*project.ProjectNodeMessage
	copier.Copy(&nodes, list)
	return &project.ProjectNodeResponseMessage{Nodes: nodes}, nil
}

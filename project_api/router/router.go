package router

import (
	"github.com/gin-gonic/gin"
	"pomfret.cn/project_api/api/project"
	"pomfret.cn/project_api/api/user"
)

// Router 接口
type Router interface {
	Route(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (*RegisterRouter) Route(router Router, r *gin.Engine) {
	router.Route(r)
}

//var routers []Router

func InitRouter(r *gin.Engine) {
	router := New()

	////以后的模块路由在这进行注册
	router.Route(&user.RouterUser{}, r)

	router.Route(&project.RouterProject{}, r)
	//for _, ro := range routers {
	//	ro.Route(r)
	//}
}

//func Register(ro ...Router) {
//	routers = append(routers, ro...)
//}
//type gRPCConfig struct {
//	Addr         string
//	RegisterFunc func(*grpc.Server)
//}

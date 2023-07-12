package router

import (
	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	account "pomfret.cn/project-grpc/account"
	auth "pomfret.cn/project-grpc/auth"
	department "pomfret.cn/project-grpc/department"
	menu "pomfret.cn/project-grpc/menu"
	task "pomfret.cn/project-grpc/task"
	"pomfret.cn/project-project/internal/rpc"
	account_serice_v1 "pomfret.cn/project-project/pkg/service/account.serice.v1"
	auth_service_v1 "pomfret.cn/project-project/pkg/service/auth.service.v1"
	department_service_v1 "pomfret.cn/project-project/pkg/service/department.service.v1"
	menu_service "pomfret.cn/project-project/pkg/service/menu.service.v1"
	task_service "pomfret.cn/project-project/pkg/service/task.service"

	"pomfret.cn/project_api/api/project"

	project1 "pomfret.cn/project-grpc/project"
	"pomfret.cn/project-project/config"
	project_service "pomfret.cn/project-project/pkg/service/project.service.v1"
	"pomfret.cn/project_common/discovery"
	"pomfret.cn/project_common/logs"
)

// 接口
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
	//以后的模块路由在这进行注册
	//router.Route(&project.RouterProject{}, r)
	//router.Route(&user.RouterUser{}, r)
	router.Route(&project.RouterProject{}, r)
	//for _, ro := range routers {
	//	ro.Route(r)
	//}
}

// 注册并启动grpc服务
type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.C.GC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			project1.RegisterProjectServiceServer(g, project_service.New())
			task.RegisterTaskServiceServer(g, task_service.New())
			account.RegisterAccountServiceServer(g, account_serice_v1.New())
			department.RegisterDepartmentServiceServer(g, department_service_v1.New())
			auth.RegisterAuthServiceServer(g, auth_service_v1.New())
			menu.RegisterMenuServiceServer(g, menu_service.New())

		}}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			otelgrpc.UnaryServerInterceptor(),
			//interceptor.New().CacheInterceptor(),
		)),
	)
	c.RegisterFunc(s)

	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	go func() {
		log.Printf("grpc server started as: %s \n", c.Addr)
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}
func RegisterEtcdServer() {
	etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)

	info := discovery.Server{
		Name:    config.C.GC.Name,
		Addr:    config.C.GC.Addr,
		Version: config.C.GC.Version,
		Weight:  config.C.GC.Weight,
	}
	r := discovery.NewRegister(config.C.EtcdConfig.Addrs, logs.LG)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}
}
func InitUserRpc() {
	rpc.InitRpcUserClient()
}

package router

import (
	"github.com/gin-gonic/gin"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	login "pomfret.cn/project-grpc/user/login"
	"pomfret.cn/project-user/config"
	loginServiceV1 "pomfret.cn/project-user/pkg/service/login.service.v1"
	"pomfret.cn/project_api/api/project"
	"pomfret.cn/project_api/api/user"
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
	router.Route(&user.RouterUser{}, r)
	router.Route(&project.RouterProject{}, r)
	//for _, ro := range routers {
	//	ro.Route(r)
	//}
}

//func Register(ro ...Router) {
//	routers = append(routers, ro...)
//
//}

// 注册并启动grpc服务
type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.C.GC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			login.RegisterLoginServiceServer(g, loginServiceV1.New())
		},
	}
	//grpc拦截器
	//cacheInterceptor := interceptor.New()
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		otelgrpc.UnaryServerInterceptor(),
		//interceptor.New().CacheInterceptor(),
	)))
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

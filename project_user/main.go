package main

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"pomfret.cn/project-user/config"
	"pomfret.cn/project-user/router"
	"pomfret.cn/project-user/tracing"
	srv "pomfret.cn/project_common"
)

func main() {
	r := gin.Default()

	tp, tpErr := tracing.JaegerTraceProvider()
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	//路由注册
	router.InitRouter(r)
	//grpc服务注册
	gc := router.RegisterGrpc()
	//grpc服务注册到etcd中
	router.RegisterEtcdServer()

	stop := func() {
		gc.Stop()
	}
	//srv.Run(r, "project_user", ":80")
	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, stop)
}

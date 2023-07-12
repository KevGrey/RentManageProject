package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"net/http"
	_ "pomfret.cn/project_api/api"
	"pomfret.cn/project_api/api/midd"
	"pomfret.cn/project_api/config"
	"pomfret.cn/project_api/router"
	"pomfret.cn/project_api/tracing"
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
	r.Use(midd.RequestLog())
	r.Use(otelgin.Middleware("project_api"))
	r.StaticFS("/upload", http.Dir("upload"))
	//路由
	router.InitRouter(r)
	//开启pprof 默认的访问路径是/debug/pprof
	pprof.Register(r)

	srv.Run(r, config.C.SC.Name, config.C.SC.Addr, nil)
}

package project

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	common "pomfret.cn/project_common"
	"pomfret.cn/project_common/errs"
	"strings"
)

var ignores = []string{
	"project/login/register",
	"project/login",
	"project/login/getCaptcha",
	"project/organization",
	"project/auth/apply"}

func Auth() func(*gin.Context) {
	return func(c *gin.Context) {
		zap.L().Info("开始做授权认证")
		//判断次uri是否在用户的授权列表中
		//根据请求的uri路径 进行匹配
		result := &common.Result{}
		uri := c.Request.RequestURI
		a := NewAuth()
		nodes, err := a.GetAuthNodes(c)
		for _, v := range ignores {
			if strings.Contains(uri, v) {
				c.Next()
				return
			}
		}
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}

		for _, v := range nodes {
			if strings.Contains(uri, v) {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusOK, result.Fail(403, "无操作权限"))
		c.Abort()
	}
}

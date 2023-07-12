package model

import (
	"pomfret.cn/project_common/errs"
)

var (
	RedisError            = errs.NewError(999, "redis错误")
	DBError               = errs.NewError(998, "db错误")
	ParamsError           = errs.NewError(401, "参数错误")
	NoLegalMobile         = errs.NewError(10102001, "手机号不合法")     //手机号不合法
	CaptchaError          = errs.NewError(10102002, "验证码错误")      //手机号不合法
	CaptchaNoExist        = errs.NewError(10102003, "验证码不存在或已过期") //手机号不合法
	EmailError            = errs.NewError(10102004, "邮箱已经存在了")    //手机号不合法
	AccountError          = errs.NewError(10102005, "账号已经存在了")    //手机号不合法
	MobileError           = errs.NewError(10102006, "手机号已经存在了")   //手机号不合法
	AccountAndPwdError    = errs.NewError(10102007, "账号密码不正确")    //手机号不合法
	TaskNameNotNull       = errs.NewError(20102001, "任务标题不能为空")
	TaskStagesNotNull     = errs.NewError(20102002, "任务步骤不存在")
	ProjectAlreadyDeleted = errs.NewError(20102003, "项目已经删除了")
)

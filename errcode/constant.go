package errcode

// This file is generated. Do not edit.

// Error code should start from 10 000 000 to ensure equal length.
// Format: AA BBB CCC
// AA: Platform (10 is for system error. Business related errors should start at 20)
// BBB: Module
// CCC: Specific error
const (
	SUCCESS = 0

	// Platform: "common 通用"
	// Module: "service system 服务系统", Code format: 10 000 xxx
	ServiceInternalError         = 10000000 // Internal error 通用服务内部错误
	ServiceUnavailable           = 10000001 // Service unavailable 服务不可用
	ServiceTimeout               = 10000002 // Service timeout 服务自身任务超时
	ServiceBusy                  = 10000003 // Service is busy 服务器繁忙
	ServiceDegradation           = 10000004 // Service degradation 服务降级停用
	ServiceObsolete              = 10000005 // Service is obsolete 服务已过期
	ServiceDependencyUnavailable = 10000006 // Service dependency unavailable 服务依赖不可达
	// Module: "request error 请求错误", Code format: 10 001 xxx
	ReqFail               = 10001000 // Common request fail 通用请求错误
	ReqSvcNotFound        = 10001001 // Service not found 通用请求错误
	ReqTooFrequent        = 10001002 // Request too frequent 请求太频繁
	ReqDuplicateOperation = 10001003 // Duplicate Operation 重复操作
	ReqOperationNotAllow  = 10001004 // Operation is not allow 操作不允许
	ReqParamIllegal       = 10001005 // Parameter illegal 参数不合法
	ReqLackRequiredParam  = 10001006 // Lack required parameters 缺少关键参数
	ReqParseParamFail     = 10001007 // Fail to parse parameters 参数无法解析
	ReqParamLengthExceed  = 10001008 // Parameter length exceed limit 参数长度超过限制
	ReqPreconditionFailed = 10001009 // precondition fail 不满足前提条件
	ReqNoPrivilege        = 10001010 // No privilege 无权限
	ReqUnauthorized       = 10001011 // Unauthorized 未授权
	ReqForbidden          = 10001012 // Forbidden 禁止访问

)

// error code default description
var codeDefaultDesc = map[int32]string{
	SUCCESS: "success",

	// Platform: "common 通用"
	// Module: "service system 服务系统"
	ServiceInternalError:         "Internal error 通用服务内部错误",
	ServiceUnavailable:           "Service unavailable 服务不可用",
	ServiceTimeout:               "Service timeout 服务自身任务超时",
	ServiceBusy:                  "Service is busy 服务器繁忙",
	ServiceDegradation:           "Service degradation 服务降级停用",
	ServiceObsolete:              "Service is obsolete 服务已过期",
	ServiceDependencyUnavailable: "Service dependency unavailable 服务依赖不可达",
	// Module: "request error 请求错误"
	ReqFail:               "Common request fail 通用请求错误",
	ReqSvcNotFound:        "Service not found 通用请求错误",
	ReqTooFrequent:        "Request too frequent 请求太频繁",
	ReqDuplicateOperation: "Duplicate Operation 重复操作",
	ReqOperationNotAllow:  "Operation is not allow 操作不允许",
	ReqParamIllegal:       "Parameter illegal 参数不合法",
	ReqLackRequiredParam:  "Lack required parameters 缺少关键参数",
	ReqParseParamFail:     "Fail to parse parameters 参数无法解析",
	ReqParamLengthExceed:  "Parameter length exceed limit 参数长度超过限制",
	ReqPreconditionFailed: "precondition fail 不满足前提条件",
	ReqNoPrivilege:        "No privilege 无权限",
	ReqUnauthorized:       "Unauthorized 未授权",
	ReqForbidden:          "Forbidden 禁止访问",
}

// CodeDefaultDesc returns the default description for the given error code.
func CodeDefaultDesc(code int32) string {
	if desc, ok := codeDefaultDesc[code]; ok {
		return desc
	} else {
		return ""
	}
}

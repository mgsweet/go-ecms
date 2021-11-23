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
    CommonServiceInternalError = 10000000 // 通用服务内部错误 internal error
    CommonServiceUnavailable = 10000001 // 服务不可用 service unavailable
    CommonServiceTimeout = 10000002 // 服务自身任务超时 service timeout
    CommonServiceBusy = 10000003 // 服务器繁忙 service is busy
    CommonServiceDegradation = 10000004 // 服务降级停用 service degradation
    CommonServiceObsolete = 10000005 // 服务已过期 service is obsolete
    CommonServiceDependencyUnavailable = 10000006 // 服务依赖不可达 service dependency unavailable
    
)

// error code default description
var codeDefaultDesc = map[int32]string{
    SUCCESS: "success",

    // Platform: "common 通用"
    // Module: "service system 服务系统"
    CommonServiceInternalError: "通用服务内部错误 internal error",
    CommonServiceUnavailable: "服务不可用 service unavailable",
    CommonServiceTimeout: "服务自身任务超时 service timeout",
    CommonServiceBusy: "服务器繁忙 service is busy",
    CommonServiceDegradation: "服务降级停用 service degradation",
    CommonServiceObsolete: "服务已过期 service is obsolete",
    CommonServiceDependencyUnavailable: "服务依赖不可达 service dependency unavailable",
    
}

// CodeDefaultDesc returns the default description for the given error code.
func CodeDefaultDesc(code int32) string {
    if desc, ok := codeDefaultDesc[code]; ok {
        return desc
    } else {
        return ""
    }
}

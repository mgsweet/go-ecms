---
title: "错误总览"
date: 2021-11-22T17:39:10+08:00
---
## common 通用 (10)
### service system 服务系统 (000)
| Name                       | Code     | Description                     |
|----------------------------|----------|---------------------------------|
| ServiceInternalError | 10000000 | Internal error 通用服务内部错误 |
| ServiceUnavailable | 10000001 | Service unavailable 服务不可用 |
| ServiceTimeout | 10000002 | Service timeout 服务自身任务超时 |
| ServiceBusy | 10000003 | Service is busy 服务器繁忙 |
| ServiceDegradation | 10000004 | Service degradation 服务降级停用 |
| ServiceObsolete | 10000005 | Service is obsolete 服务已过期 |
| ServiceDependencyUnavailable | 10000006 | Service dependency unavailable 服务依赖不可达 |
### request error 请求错误 (001)
| Name                       | Code     | Description                     |
|----------------------------|----------|---------------------------------|
| ReqFail | 10001000 | Common request fail 通用请求错误 |
| ReqSvcNotFound | 10001001 | Service not found 通用请求错误 |
| ReqTooFrequent | 10001002 | Request too frequent 请求太频繁 |
| ReqDuplicateOperation | 10001003 | Duplicate Operation 重复操作 |
| ReqOperationNotAllow | 10001004 | Operation is not allow 操作不允许 |
| ReqParamIllegal | 10001005 | Parameter illegal 参数不合法 |
| ReqLackRequiredParam | 10001006 | Lack required parameters 缺少关键参数 |
| ReqParseParamFail | 10001007 | Fail to parse parameters 参数无法解析 |
| ReqParamLengthExceed | 10001008 | Parameter length exceed limit 参数长度超过限制 |
| ReqPreconditionFailed | 10001009 | precondition fail 不满足前提条件 |
| ReqNoPrivilege | 10001010 | No privilege 无权限 |
| ReqUnauthorized | 10001011 | Unauthorized 未授权 |
| ReqForbidden | 10001012 | Forbidden 禁止访问 |

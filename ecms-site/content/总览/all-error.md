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
| ReqSvcNotFound | 10001001 | Service not found 通用请求错误  |
| ReqTooFrequent | 10001002 | Request too frequent 请求太频繁 |
| ReqDuplicateOperation | 10001003 | Duplicate Operation 重复操作 |

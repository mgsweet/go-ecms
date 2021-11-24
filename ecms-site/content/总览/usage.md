---
title: "使用方法"
date: 2021-11-22T17:39:10+08:00
---
## 前提条件
If you want to build site, install Hugo. For mac with homebrew:
```shell
brew install hugo
```
For other platforms, see: https://gohugo.io/getting-started/installing/

## 如何生成代码
To generate both go code and site code, run:
```shell
./build.sh
```

To build go code only, run:
```shell
./build_go.sh
```

To build site only, run:
```shell
./build_site.sh
```

## 如何新增平台
以新增 *组织中心 (organization)* 平台为例
1. 在 `platforms/` 文件夹下新增一个文件夹，文件夹名称为平台标识，例如 `organization/`
2. 在 `platforms/config.yaml` 填写相应配置，例如：
```yaml
- name: "organization 组织中心" # 这个名字只会影响生成注释和网页文档
  code: "20" # 平台标识，用于生成错误码，业务相关的错误码从 20 开始, 需要注意不能和其他平台的前缀重复
  prefix: "Org" # 该平台下生成的错误码名字前缀，需要注意不能和其他平台的前缀重复
  dir: "organization" # 对应刚刚创建的文件夹名字
```
3. 新建文件 `platforms/organization/config.yaml` 里面需要至少填写一个空的 `modules` yaml 数组对象：
```yaml
modules:
```
4. 检查 `platforms/config.yaml` 中的 `code` 和 `prefix` 是否唯一。

## 如何新增模块
以新增 *组织中心 (organization)* 的 *权限服务为例 (permission)* 模块为例
1. 在 `platforms/organization/` 创建 `permission.yaml` 文件, 最少保留一个空的 `specific_errors:` yaml 数组对象，建议填写一个 *000* 通用模块错误：
```yaml
specific_errors:
- suffix: "Fail" # 错误码后缀，用于生成错误码名字，需要在当前文件唯一，首字母需大写
  code: "000" # 具体错误码，对应 CCC, 需要在当前文件唯一
  desc: "Organization permission service fail 组织中心权限服务通用错误" # 只会影响注释和网页文档
```
2. 在 `platforms/organization/config.yaml` 里面填写相应配置:
```yaml
- name: "permission error 权限服务模块错误" # 只会影响注释和网页文档
  code: "001" # 模块标识，对应 BBB, 需在当前文件内唯一
  prefix: "Permission" # 模块错误码前缀，需要注意在当前文件内唯一
  file: "permission.yaml" # 对应刚刚创建的 yaml 文件名
```

## 如何新增具体错误码
以新增 *组织中心 (organization)* 的 *权限服务为例 (permission)* 模块下的 无授权错误 (no authentication) 为例:
1. 在 `platforms/organization/permission.yaml` 中填写相应配置：
```yaml
- suffix: "NoAuth" # 错误码后缀，最终生成: OrgPermissionNoAuth，需要在当前文件唯一，首字母需大写
  code: "001" # 具体错误码，最终生成：21 001 001, 需要在当前文件唯一
  desc: "No authentication from organization permission control center 无组织中心权限系统授权" # 只会影响注释和网页文档
```
2. 运行命令生成代码

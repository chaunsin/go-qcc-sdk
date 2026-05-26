# go-qcc-sdk

[GoDoc](https://godoc.org/github.com/chaunsin/go-qcc-sdk) [Go Report Card](https://goreportcard.com/report/github.com/chaunsin/go-qcc-sdk)

企查查go sdk

## 注意

此项目为非官方项目，接口定义等相关内容以企查查官网为准: [https://openapi.qcc.com/dataApi](https://openapi.qcc.com/dataApi)

目前项目处于开发阶段可能会有目录、方法等重大调整,需要注意！

欢迎贡献代码！！！

## 功能

目前开发分两步走：

- [X] 实现企查查 OpenAPI 接口的基础查询能力。SDK 会按官方接口文档持续补充和修正接口定义，已实现接口以仓库中的 `{ApiCode}.go` 文件和代码导出的 Go 方法为准。
- [ ] 增加数据缓存能力。由于部分接口数据更新频率较低，多次请求结果可能相同，且接口调用通常会产生费用，后续计划支持对接口结果进行可配置缓存。

企查查接口会随官方文档调整而变化，需要查看当前支持范围时，可直接浏览根目录下的 `{ApiCode}.go` 文件，或使用仓库内置的 `qcc-check` 技能对本地实现与官方文档进行覆盖率审计。

## 快速开始

> go: 1.23.0 

```go
package example

import (
	"context"
	"fmt"

	api "github.com/chaunsin/go-qcc-sdk"
)

func Example() {
	cli := api.New(&api.Config{
		Key:       "<your key>",
		SecretKey: "<your secret key>",
		Debug:     true,
	})
	resp, err := cli.FuzzySearchGetList(context.Background(), &api.FuzzySearchGetListReq{
		SearchKey: "企查查科技股份有限公司",
		PageSize:  10,
		PageIndex: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp: %+v\n", resp)
}

```

## AI Agent 快速开发

本仓库在 `skills` 目录内置了面向 AI agent 的开发技能，可用于快速审计、创建或更新企查查 SDK 接口。建议先审计接口覆盖情况，再按需创建或修复指定 ApiCode。

### 使用方式

- 审计官方文档与本地 SDK 覆盖情况：`/qcc:check` 或 `/qcc-check`
- 创建或更新指定 ApiCode：`/qcc:create 886` 或 `/qcc-create https://openapi.qcc.com/dataApi/886`

### 技能说明


| 技能           | 目录                  | 用途                                                  |
| ------------ | ------------------- | --------------------------------------------------- |
| `qcc-check`  | `skills/qcc-check`  | 从官方文档统计 ApiCode 与接口数量，并与本地 `{ApiCode}.go` 实现做覆盖率对比。 |
| `qcc-create` | `skills/qcc-create` | 根据官方接口文档创建、审计或修复对应的 Go SDK 请求/响应结构与方法。              |


`skills` 是唯一维护源，`.agents/skills` 与 `.claude/skills` 通过符号链接指向该目录，方便不同 AI agent 复用同一套技能说明。官方接口信息以 [https://openapi.qcc.com/dataApi](https://openapi.qcc.com/dataApi) 为准；使用技能时不要保存真实 QCC key、secret、cookie 或付费接口响应。

## 问题

由于企查查得接口可能不定时得更新,接口出入参数可能发生变化，如有问题请在提issue时提供接口名称、ApiCode，方便我及时更新。
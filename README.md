# go-qcc-sdk

[![GoDoc](https://godoc.org/github.com/chaunsin/go-qcc-sdk?status.svg)](https://godoc.org/github.com/chaunsin/go-qcc-sdk) [![Go Report Card](https://goreportcard.com/badge/github.com/chaunsin/go-qcc-sdk)](https://goreportcard.com/report/github.com/chaunsin/go-qcc-sdk)


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
		PageIndex: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp: %+v\n", resp)
}

```

## AI Agent 技能

本仓库在 `skills` 目录内置了面向 AI agent 的技能，适用于 Codex 和 Claude, 技能分为应用接入和 SDK 维护两类。普通业务项目接入本 SDK 时优先使用 `go-qcc-sdk`；只有维护本仓库接口实现时才使用 `qcc-check` 或 `qcc-create`。

### 技能说明

| 技能           | 目录                    | 用途                                                  |
| ------------ | --------------------- | --------------------------------------------------- |
| `go-qcc-sdk` | `skills/go-qcc-sdk`   | 面向 SDK 使用者，帮助在 Go 应用中安装依赖、加载凭证、配置客户端、调用已有方法、处理 `Response[T]`/`Result`/`Paging`、编写业务封装和本地 mock/`httptest` 测试；不用于修改 SDK 源码。 |
| `qcc-check`  | `skills/qcc-check`    | 面向 SDK 维护者，从官方文档统计 ApiCode 与接口数量，并与本地 `{ApiCode}.go` 实现做覆盖率对比。 |
| `qcc-create` | `skills/qcc-create`   | 面向 SDK 维护者，根据官方接口文档创建、审计或修复对应的 Go SDK 请求/响应结构与方法。 |

### go-qcc-sdk 使用方式

安装后，可以在支持技能的 AI agent 中直接提出接入需求，例如：

```text
使用 go-qcc-sdk 帮我在 Go 服务中接入企查查 SDK，凭证从环境变量读取，并为调用逻辑加 mock 测试。
```

如果你的 agent 支持显式技能调用，也可以使用：

```text
$go-qcc-sdk 帮我封装一个企业模糊搜索服务，并说明错误处理方式。
```

`go-qcc-sdk` 只处理 SDK 消费侧问题，例如配置、调用已有方法、错误处理、业务封装和本地测试；不要用它来新增接口、修复 SDK 结构体或做官方文档覆盖率审计。

### 安装 go-qcc-sdk 技能

技能命令以 [vercel-labs/skills](https://github.com/vercel-labs/skills) 的 `skills` CLI 为准。以下命令建议在你的业务项目根目录执行，用于把本仓库的 `go-qcc-sdk` 技能安装到当前项目或用户全局环境。

```bash
# 查看本仓库可安装的技能
npx skills add chaunsin/go-qcc-sdk --list

# 安装到当前项目作用域（默认）：适合随项目共享给团队
npx skills add chaunsin/go-qcc-sdk --skill go-qcc-sdk

# 安装到全局作用域：适合所有项目复用
npx skills add chaunsin/go-qcc-sdk --skill go-qcc-sdk --global

# 只安装给 Codex
npx skills add chaunsin/go-qcc-sdk --skill go-qcc-sdk --agent codex

# 从本地 clone 的仓库安装（替换为实际 clone 目录）
npx skills add /path/to/go-qcc-sdk --skill go-qcc-sdk
```

常用作用域和安装选项：

| 目标         | 命令选项                         | 说明 |
| ---------- | ------------------------------ | ---- |
| 当前项目       | 默认，无需额外参数                    | 安装到当前项目的 agent skills 目录，适合提交到业务项目。 |
| 用户全局       | `--global` 或 `-g`              | 安装到用户目录，适合跨项目复用。 |
| 指定 agent   | `--agent codex` 或 `-a codex`  | 只安装到指定 agent；也可多次传入 `-a`。 |
| 复制文件       | `--copy`                       | 不使用符号链接，改为复制技能文件。 |
| 非交互确认      | `--yes` 或 `-y`                | 跳过确认提示，适合脚本或 CI。 |

### 更新和卸载技能

```bash
# 查看已安装技能
npx skills list

# 更新 go-qcc-sdk（交互选择作用域）
npx skills update go-qcc-sdk

# 只更新当前项目作用域
npx skills update go-qcc-sdk --project

# 只更新全局作用域
npx skills update go-qcc-sdk --global

# 卸载 go-qcc-sdk（交互选择安装位置）
npx skills remove go-qcc-sdk

# 从全局作用域卸载
npx skills remove --global go-qcc-sdk

# 只从 Codex 卸载
npx skills remove --agent codex --skill go-qcc-sdk
```

### SDK 维护技能使用方式

- 审计官方文档与本地 SDK 覆盖情况：`/qcc:check` 或 `/qcc-check`
- 创建或更新指定 ApiCode：`/qcc:create 886` 或 `/qcc-create https://openapi.qcc.com/dataApi/886`

`skills` 是本仓库技能的唯一维护源，`.agents/skills` 与 `.claude/skills` 通过符号链接指向该目录，方便不同 AI agent 复用同一套技能说明。官方接口信息以 [https://openapi.qcc.com/dataApi](https://openapi.qcc.com/dataApi) 为准；使用技能时不要保存真实 QCC key、secret、cookie 或付费接口响应。

## 问题

由于企查查得接口可能不定时得更新,接口出入参数可能发生变化，如有问题请在提issue时提供接口名称、ApiCode，方便我及时更新。

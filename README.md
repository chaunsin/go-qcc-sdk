# go-qcc-sdk

企查查go sdk

## 注意

此项目为个人项目，接口定义等相关内容以企查查官网为准。https://openapi.qcc.com/dataApi

## 功能

分两步走

1. 实现接口得基本查询功能(**进行中**)

2. 增加数据缓存。由于有些接口部分数据更新频率较低，多次请求结果是一样得，且接口调用收费，因此本sdk支持缓存接口数据功能。

| ApiCode |    接口名称    | 状态  |
|:-------:|:----------:|:---:|
|   213   |   企业年报信息   | 已实现 |
|   271   |   税号开票信息   | 已实现 |
|   410   |   企业工商照面   | 已实现 |
|   643   |   实际控制人    | 已实现 |
|   669   | 企业人员董监高信息  | 已实现 |
|   731   | 股东信息(工商登记) | 已实现 |
|   732   |    主要人员    | 已实现 |
|   734   |    变更记录    | 已实现 |
|   735   |   企业工商详情   | 已实现 |
|   736   |   企业风险扫描   | 已实现 |
|   763   |  董监高风险扫描   | 已实现 |
|   855   |  企业二要素核验   | 已实现 |
|   865   |   行政处罚核查   | 已实现 |
|   884   |   行政处罚核查   | 已实现 |
|   886   |   企业高级搜索   | 已实现 |
|   962   |    企业图谱    | 已实现 |
|  2001   |   企业信息核验   | 已实现 |
|  2003   |   客户身份识别   | 已实现 |

## 快速开始

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

## 问题

由于企查查得接口可能不定时得更新,接口出入参数可能发生变化，如有问题请在提issue时提供接口名称、ApiCode，方便我及时更新。
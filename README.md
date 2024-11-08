# go-qcc-sdk

[![GoDoc](https://godoc.org/github.com/chaunsin/go-qcc-sdk?status.svg)](https://godoc.org/github.com/chaunsin/go-qcc-sdk) [![Go Report Card](https://goreportcard.com/badge/github.com/chaunsin/go-qcc-sdk)](https://goreportcard.com/report/github.com/chaunsin/go-qcc-sdk)

企查查go sdk

## 注意

此项目为个人项目，接口定义等相关内容以企查查官网为准: https://openapi.qcc.com/dataApi

目前项目处于开发阶段可能会有目录、方法等重大调整,需要注意！

## 功能

目前开发分两步走

1. 实现接口得基本查询功能(**进行中**)

2. 增加数据缓存。由于有些接口部分数据更新频率较低，多次请求结果是一样得，且接口调用收费，因此本sdk支持缓存接口数据功能。

| ApiCode |           接口名称            | 状态  |
|:-------:|:-------------------------:|:---:|
|   213   |          企业年报信息           | 已实现 |
|   271   |          税号开票信息           | 已实现 |
|   410   |          企业工商照面           | 已实现 |
|   628   |        企业受益股东穿透识别         | 已实现 |
|   642   |         股权穿透(四层)          | 已实现 |
|   643   |           实际控制人           | 已实现 |
|   663   |       企业对外投资穿透(十层)        | 已实现 |
|   669   |         企业人员董监高信息         | 已实现 |
|   699   | 上市企业(企业简介&主要指标&十大股东&企业高管) | 已实现 |
|   731   |        股东信息(工商登记)         | 已实现 |
|   732   |           主要人员            | 已实现 |
|   733   |           分支机构            | 已实现 |
|   734   |           变更记录            | 已实现 |
|   735   |          企业工商详情           | 已实现 |
|   736   |          企业风险扫描           | 已实现 |
|   740   |           失信核查            | 已实现 |
|   741   |          被执行人核查           | 已实现 |
|   742   |          限制高消费核查          | 已实现 |
|   744   |          司法拍卖核查           | 已实现 |
|   752   |          股权冻结核查           | 已实现 |
|   761   |          破产重整核查           | 已实现 |
|   763   |          董监高风险扫描          | 已实现 |
|   855   |          企业二要素核验          | 已实现 |
|   856   |          企业三要素核验          | 已实现 |
|   865   |          行政处罚核查           | 已实现 |
|   882   |          法院公告核查           | 已实现 |
|   884   |         企业对外投资核查          | 已实现 |
|   886   |          企业高级搜索           | 已实现 |
|   887   |          裁判文书核查           | 已实现 |
|   888   |          开庭公告核查           | 已实现 |
|   889   |          立案信息核查           | 已实现 |
|   914   |          上市企业搜索           | 已实现 |
|   915   |          上市公告搜索           | 已实现 |
|   942   |          控制企业核查           | 已实现 |
|   943   |         董监高控制企业核查         | 已实现 |
|   962   |           企业图谱            | 已实现 |
|  1002   |          企业关联关系           | 已实现 |
|  1003   |           受益所有人           | 已实现 |
|  1026   |      股东信息(最新公示&工商登记)      | 已实现 |
|  2001   |          企业信息核验           | 已实现 |
|  2003   |          客户身份识别           | 已实现 |

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
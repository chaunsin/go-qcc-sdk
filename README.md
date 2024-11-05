# go-qcc-sdk

企查查go sdk

# 注意

此项目为个人项目，接口定义等相关内容以企查查官网为准。https://openapi.qcc.com/dataApi

### 目前处于开发中,部分接口可以使用。

# 快速开始

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

# 问题

由于企查查得接口可能不定时得更新,接口出入参数可能发生变化，如有问题请在提issue时提供接口名称、ApiCode，方便我及时更新。
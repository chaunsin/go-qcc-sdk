// MIT License
//
// Copyright (c) 2024 chaunsin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package example

import (
	"context"
	"errors"
	"fmt"

	api "github.com/chaunsin/go-qcc-sdk"
)

func Example() {
	cli := api.New(&api.Config{
		Key:       "<你的key>",
		SecretKey: "<你的SecretKey>",
		Debug:     true,
	})
	resp, err := cli.FuzzySearchGetList(context.Background(), &api.FuzzySearchGetListReq{
		SearchKey: "企查查科技股份有限公司",
		PageIndex: 1,
	})
	if err != nil {
		var apiErr *api.APIError
		if errors.As(err, &apiErr) {
			switch apiErr.Status {
			case api.StatusKeyArrears: // KEY 已欠费
				// ... business handling
			case api.StatusProcessing: // 205 等待处理中，可稍后轮询重试
				// ... poll later
			default:
				fmt.Printf("qcc error: %s\n", apiErr.Message)
			}
			return
		}
		// 其他错误，网络错误，异常错误等。
		panic(err)
	}

	if resp == nil {
		fmt.Println("查询无结果")
	} else {
		fmt.Printf("resp: %+v\n", resp)
	}
}

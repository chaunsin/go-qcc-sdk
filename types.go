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

package api

type Paging struct {
	PageSize     int64 `json:"PageSize,omitempty"`
	PageIndex    int64 `json:"PageIndex,omitempty"`
	TotalRecords int64 `json:"TotalRecords,omitempty"`
}

type Response[T any] struct {
	Status      Status `json:"Status"`
	Message     string `json:"Message"`
	OrderNumber string `json:"OrderNumber"`
	Paging      Paging `json:"Paging,omitempty"`
	Result      T      `json:"Result"`
}

// err 返回响应信封携带的业务错误。
// StatusSuccess(200) 返回 nil；StatusNoResult(201) 也返回 nil（各 endpoint 会在其之前
// 提前 return nil, nil 处理，此处仅作兜底）；其余状态码（含 205 等待处理中）返回 *APIError。
// 私有方法，暂不对外暴露；通过方法提升，嵌入 Response[T] 的 XxxResp 可直接调用 resp.err()。
func (r *Response[T]) err() error {
	switch r.Status {
	case StatusSuccess, StatusNoResult:
		return nil
	default:
		return &APIError{Status: r.Status, Message: r.Message, OrderNumber: r.OrderNumber}
	}
}

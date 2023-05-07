package server

import (
	"github.com/galaxy-toolkit/server/server/code"
	"github.com/gofiber/fiber/v2"
)

// BasicResponse 基础请求响应
type BasicResponse struct {
	Code code.Code `json:"code"` // 响应状态码
	Msg  string    `json:"msg"`  // 响应信息
}

// DataResponse 带响应数据响应
type DataResponse[DATA any] struct {
	BasicResponse
	Data DATA `json:"data"` // 响应数据
}

// PageResponse 分页查询内层响应
type PageResponse[DATA any] struct {
	Page     int   `json:"page"`      // 当前页
	PageSize int   `json:"page_size"` // 每页大小
	Total    int64 `json:"total"`     // 总数
	Data     DATA  `json:"data"`      // 数据
}

// SendOk 请求响应成功
func SendOk(ctx *fiber.Ctx) error {
	return ctx.JSON(BasicResponse{
		Code: code.SUCCESS,
		Msg:  code.GetMsg(code.SUCCESS),
	})
}

// SendFailed 请求失败
func SendFailed(ctx *fiber.Ctx) error {
	return ctx.JSON(BasicResponse{
		Code: code.Failed,
		Msg:  code.GetMsg(code.Failed),
	})
}

// SendCode 返回任意 code
func SendCode(ctx *fiber.Ctx, c code.Code) error {
	return ctx.JSON(BasicResponse{
		Code: c,
		Msg:  code.GetMsg(c),
	})
}

// SendDataOk 请求响应成功，并返回数据
func SendDataOk[DATA any](ctx *fiber.Ctx, data DATA) error {
	return ctx.JSON(DataResponse[DATA]{
		BasicResponse: BasicResponse{
			Code: code.SUCCESS,
			Msg:  code.GetMsg(code.SUCCESS),
		},
		Data: data,
	})
}

// SendDataCode 并返回数据，并指定 code
func SendDataCode[DATA any](ctx *fiber.Ctx, data DATA, c code.Code) error {
	return ctx.JSON(DataResponse[DATA]{
		BasicResponse: BasicResponse{
			Code: c,
			Msg:  code.GetMsg(c),
		},
		Data: data,
	})
}

// SendPageDataOk 请求响应成功，并返回分页数据
func SendPageDataOk[DATA any](ctx *fiber.Ctx, data DATA, page, pageSize int, total int64) error {
	return ctx.JSON(DataResponse[PageResponse[DATA]]{
		BasicResponse: BasicResponse{
			Code: code.SUCCESS,
			Msg:  code.GetMsg(code.SUCCESS),
		},
		Data: PageResponse[DATA]{
			Page:     page,
			PageSize: pageSize,
			Total:    total,
			Data:     data,
		},
	})
}

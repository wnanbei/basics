package server

import (
	"github.com/galaxy-toolkit/server/constant/code"
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

// ParamsParseFailedResponse 参数解析失败响应
type ParamsParseFailedResponse struct {
	BasicResponse
	FailedFields []*FailedField `json:"failed_fields"` // 异常字段
}

// SendJson 返回 json 响应。用于收拢响应返回
func SendJson(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(data)
}

// SendOk 请求响应成功
func SendOk(ctx *fiber.Ctx) error {
	return SendJson(ctx, BasicResponse{
		Code: code.SUCCESS,
		Msg:  code.GetMsg(code.SUCCESS),
	})
}

// SendFailed 请求失败
func SendFailed(ctx *fiber.Ctx) error {
	return SendJson(ctx, BasicResponse{
		Code: code.Failed,
		Msg:  code.GetMsg(code.Failed),
	})
}

// SendCode 返回任意 code
func SendCode(ctx *fiber.Ctx, c code.Code) error {
	return SendJson(ctx, BasicResponse{
		Code: c,
		Msg:  code.GetMsg(c),
	})
}

// SendDataOk 请求响应成功，并返回数据
func SendDataOk[DATA any](ctx *fiber.Ctx, data DATA) error {
	return SendJson(ctx, DataResponse[DATA]{
		BasicResponse: BasicResponse{
			Code: code.SUCCESS,
			Msg:  code.GetMsg(code.SUCCESS),
		},
		Data: data,
	})
}

// SendDataCode 并返回数据，并指定 code
func SendDataCode[DATA any](ctx *fiber.Ctx, data DATA, c code.Code) error {
	return SendJson(ctx, DataResponse[DATA]{
		BasicResponse: BasicResponse{
			Code: c,
			Msg:  code.GetMsg(c),
		},
		Data: data,
	})
}

// SendPageDataOk 请求响应成功，并返回分页数据
func SendPageDataOk[DATA any](ctx *fiber.Ctx, data DATA, page, pageSize int, total int64) error {
	return SendJson(ctx, DataResponse[PageResponse[DATA]]{
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

// SendParamsParseFailed 返回解析失败的字段信息
func SendParamsParseFailed(ctx *fiber.Ctx, failedFields []*FailedField) error {
	return SendJson(ctx, ParamsParseFailedResponse{
		BasicResponse: BasicResponse{
			Code: code.ParamsParseFailed,
			Msg:  code.GetMsg(code.ParamsParseFailed),
		},
		FailedFields: failedFields,
	})
}

// SendError 根据 err 确定返回内容
// 如果 err 为 nil,则返回 Success
// 如果 err 为 *code.Error,则返回对应的 code 和 msg
// 如果 err 不为 nil,则返回 Failed
func SendError(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return SendJson(ctx, BasicResponse{
			Code: code.SUCCESS,
			Msg:  code.GetMsg(code.SUCCESS),
		})
	}

	if e, ok := err.(*code.Error); ok {
		return SendJson(ctx, BasicResponse{
			Code: e.Code,
			Msg:  code.GetMsg(e.Code),
		})
	}

	return SendJson(ctx, BasicResponse{
		Code: code.Failed,
		Msg:  code.GetMsg(code.Failed),
	})
}

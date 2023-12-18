package model

import (
	"github.com/kataras/iris/v12"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 400
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, ctx iris.Context) {
	ctx.JSON(Response{
		code,
		data,
		msg,
	})
}

func Ok(ctx iris.Context) {
	Result(SUCCESS, map[string]interface{}{}, "ok", ctx)
}

func OkWithMessage(message string, ctx iris.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, ctx)
}

func OkWithData(data interface{}, ctx iris.Context) {
	Result(SUCCESS, data, "success", ctx)
}

func OkWithDetailed(data interface{}, message string, ctx iris.Context) {
	Result(SUCCESS, data, message, ctx)
}

func Fail(ctx iris.Context) {
	Result(ERROR, map[string]interface{}{}, "failed", ctx)
}

func FailWithMessage(message string, ctx iris.Context) {
	Result(ERROR, map[string]interface{}{}, message, ctx)
}

func FailWithDetailed(data interface{}, message string, ctx iris.Context) {
	Result(ERROR, data, message, ctx)
}

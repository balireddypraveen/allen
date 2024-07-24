package response

import (
	"github.com/balireddypraveen/allen/internal/common/constants"
	"github.com/balireddypraveen/allen/internal/pkg/context"
	"github.com/balireddypraveen/allen/internal/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Response gin.H

func FormatResponse(data interface{}, status bool, err error, message string) *Response {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	res := &Response{
		"success": status,
		"data":    data,
		"code":    errMsg,
		"msg":     message,
	}
	return res
}

// Success data argument should be gin.H{} or []gin.H{}
func Success(ctx *gin.Context, reqCtx context.ReqCtx, statusCode int, data interface{}) {
	if ctx.Writer.Written() {
		reqCtx.Log.Info("response body was already written! will not overwrite")
		return
	}
	res := FormatResponse(data, true, nil, constants.EmptyString)

	ctx.Writer.Header().Set("X-Request-Id", ctx.GetString("X-Request-Id"))
	ctx.Writer.Header().Set("X-Span-Request-Id", ctx.GetString("X-Span-Request-Id"))
	ctx.Writer.Header().Set("X-Amzn-Trace-Id", ctx.GetString("X-Amzn-Trace-Id"))

	ctx.JSON(statusCode, res)
}

func Fail(ctx *gin.Context, statusCode int, errors []gin.H, msg string) {
	log := logger.GetLogger()
	if ctx.Writer.Written() {
		log.Warn("response body was already written! will not overwrite")
		return
	}
	res := FormatResponse(Response{
		"errors": errors,
	}, false, nil, msg)

	ctx.JSON(statusCode, res)
}

func Error(ctx *gin.Context, reqCtx context.ReqCtx, statusCode int, errCode ErrorCode, msg string) {
	if ctx.Writer.Written() {
		reqCtx.Log.Info("response body was already written! will not overwrite")
		return
	}
	res := Response{
		"success": false,
		"data":    nil,
		"code":    errCode,
		"msg":     msg,
	}

	ctx.Writer.Header().Set("X-Request-Id", ctx.GetString("X-Request-Id"))
	ctx.Writer.Header().Set("X-Span-Request-Id", ctx.GetString("X-Span-Request-Id"))
	ctx.Writer.Header().Set("X-Amzn-Trace-Id", ctx.GetString("X-Amzn-Trace-Id"))
	ctx.JSON(statusCode, res)
}

package xresult

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zrpcErr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc/status"
	"my_chat/pkg/xerr"
	"net/http"
)

// 返回的数据结构
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any) *Response {
	return &Response{
		Code: 200,
		Msg:  "Success",
		Data: data,
	}
}

func Fail(code int, err string) *Response {
	return &Response{
		Code: code,
		Msg:  err,
		Data: nil,
	}
}

func OkHandler(_ context.Context, v any) any {
	return Success(v)
}

func ErrHandler(name string) func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		// 设置默认错误码和错误信息
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := xerr.ErrMsg(errcode)

		causeErr := errors.Cause(err)

		// 判断是不是业务错误
		if e, ok := causeErr.(*zrpcErr.CodeMsg); ok {
			errcode = e.Code
			errmsg = e.Msg
		} else {
			// 判断是不是 gRPC 错误
			if grpcStatus, ok := status.FromError(causeErr); ok {
				grrpcCode := int(grpcStatus.Code())
				errcode = grrpcCode
				errmsg = grpcStatus.Message()
			}
		}

		// 记录错误日志
		logx.WithContext(ctx).Errorf("[serverName = %s], error : %v", name, err)
		return http.StatusBadRequest, Fail(errcode, errmsg)
	}
}

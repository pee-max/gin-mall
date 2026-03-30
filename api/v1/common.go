package v1

import (
	"encoding/json"
	"errors"
	"gin_mall/serializer"
	"net/http"
)

func ErrorResponse(err error) serializer.Response {
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    "JSON 格式不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusBadRequest,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}

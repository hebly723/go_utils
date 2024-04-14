package utils

import (
	"encoding/json"
	"strconv"
)

const (
	SUCCESS_CODE = 0
)

type FormatResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Code    int                     `json:"code"`
	Data    *map[string]interface{} `json:"data"`
}

func NewFormatResponseWithMap(message error, i *map[string]interface{}) string {
	var f FormatResponse
	f.Success = true
	f.Code = SUCCESS_CODE

	if message != nil {
		f.Success = false
		f.Message = message.Error()
		customErr, ok := message.(*CustomError)
		if !ok {
			panic(message)
		}
		f.Code = customErr.code
	}
	f.Data = i
	bs, err := json.Marshal(f)
	if err != nil {
		return "{\"success\":false,\"code\":\"" + strconv.Itoa(ErrJsonMarshalerFailed.code) +
			"\", \"msg\":\"" + ErrJsonMarshalerFailed.msg + "\"}"
	}
	return string(bs)
}

func NewFormatResponseWithStruct(message error, i interface{}) string {
	bjs, e := json.Marshal(i)

	if e != nil {
		return "{\"success\":false,\"code\":\"" + strconv.Itoa(ErrJsonMarshalerFailed.code) +
			"\", \"msg\":\"" + ErrJsonMarshalerFailed.msg + "\"}"
	}

	m := new(map[string]interface{})

	if err := json.Unmarshal(bjs, m); err != nil {
		return "{\"success\":false,\"code\":\"" + strconv.Itoa(ErrJsonMarshalerFailed.code) +
			"\", \"msg\":\"" + ErrJsonMarshalerFailed.msg + "\"}"
	}

	return NewFormatResponseWithMap(message, m)
}

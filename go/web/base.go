package web

import (
	"base"
	"bytes"
	"errors"
	"io"
	"net/http"
)

type BaseResponse struct {
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Value interface{} `json:"value"`
}

var (
	OK_RESULT        = "{\"msg\": \"\",\"code\":\"0\"}"
	NOT_LOGIN_RESULT = "{\"msg\": \"未检测到登录用户，请先登录\",\"code\":100}"
	ERROR_RESULT     = "{\"msg\": \"服务异常\",\"code\":\"-1\"}"
)

func getPostBody(r *http.Request) (data []byte, err error) {
	if r.Method != "POST" {
		return nil, errors.New("METHOD_ERROR")
	}
	requestBody := r.Body
	buf := new(bytes.Buffer)
	buf.ReadFrom(requestBody)
	data = buf.Bytes()
	defer requestBody.Close()
	if err != nil {
		return nil, err
	}
	return data, err
}

func ToBean(r *http.Request, req interface{}) (err error) {
	var body []byte
	body, err = getPostBody(r)
	if err != nil {
		return
	}
	err = base.JSON.Unmarshal(body, req)
	return
}

func outJSON(w http.ResponseWriter, value interface{}, err error) {
	out := BaseResponse{}
	if err != nil {
		out.Code = "-1"
		out.Msg = err.Error()
	} else {
		out.Code = "0"
		out.Value = value
	}
	w.Header().Set("Content-Type", "application/json")
	by, err := base.JSON.Marshal(out)
	if err != nil {
		io.WriteString(w, ERROR_RESULT)
		return
	}
	w.Write(by)
}

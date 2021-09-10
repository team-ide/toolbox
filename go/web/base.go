package web

import (
	"base"
	"bytes"
	"errors"
	"io"
	"net/http"
)

type BaseResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

var (
	OK_RESULT    = "{\"msg\": \"\",\"code\":\"0\"}"
	ERROR_RESULT = "{\"msg\": \"服务异常\",\"code\":\"-1\"}"
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

func outJSON(w http.ResponseWriter, res interface{}, err error) {
	if err != nil {
		res = BaseResponse{
			Code: "-1",
			Msg:  err.Error(),
		}
	}
	if res == nil {
		res = BaseResponse{
			Code: "0",
			Msg:  "OK",
		}
	}
	w.Header().Set("Content-Type", "application/json")
	by, err := base.JSON.Marshal(res)
	if err != nil {
		io.WriteString(w, ERROR_RESULT)
		return
	}
	w.Write(by)
}

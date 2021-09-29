package web

import (
	"base"
	"config"
	"errors"
	"io"
	"net/http"
	"worker"

	"github.com/gorilla/sessions"
)

func handleOpen(rw http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	workers := []map[string]interface{}{}

	workerCache := worker.WorkerCache
	for _, oneWorker := range workerCache {
		workMap := oneWorker.WorkMap
		workerData := make(map[string]interface{})
		workerData["name"] = oneWorker.Name
		workerData["icon"] = oneWorker.Icon
		workerData["comment"] = oneWorker.Comment
		workerData["text"] = oneWorker.Text
		workerData["configs"] = oneWorker.Configs

		workerWorks := []map[string]interface{}{}
		for workName := range workMap {
			workData := make(map[string]interface{})

			workData["name"] = workName

			workerWorks = append(workerWorks, workData)
		}
		workerData["works"] = workerWorks
		workers = append(workers, workerData)

	}
	data["workers"] = workers

	enum_map := map[string]interface{}{}

	sqlConditionalOperations := []base.Option{}
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "等于", Value: "="})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不等于", Value: "<>"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "大于", Value: ">"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "大于或等于", Value: ">="})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "小于", Value: "<"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "小于或等于", Value: "<="})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "包含", Value: "like"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不包含", Value: "not like"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "开始以", Value: "like start"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "开始不是以", Value: "not like start"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "结束以", Value: "like end"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "结束不是以", Value: "not like end"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "是null", Value: "is null"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不是null", Value: "is not null"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "是空", Value: "is empty"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不是空", Value: "is not empty"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "介于", Value: "between"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不介于", Value: "not between"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "在列表", Value: "in"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "不在列表", Value: "not in"})
	sqlConditionalOperations = append(sqlConditionalOperations, base.Option{Text: "自定义", Value: "custom"})

	enum_map["sqlConditionalOperations"] = sqlConditionalOperations

	data["enum_map"] = enum_map

	outJSON(rw, data, nil)
}

var sessionStore = sessions.NewCookieStore([]byte("TEAMIDE-TOOLBOX"))

func init() {
	sessionStore.MaxAge(60 * 60 * 2)
}

func SetSessionUser(w http.ResponseWriter, r *http.Request, user *base.LoginUser) error {
	param := GetSessionParam(r)
	param.User = user
	return SetSessionParam(w, r, param)
}

func SetSessionParam(w http.ResponseWriter, r *http.Request, sessionParam *base.SessionParam) error {
	var err error
	session, _ := sessionStore.Get(r, "teamide-toolbox")
	value := ""
	if sessionParam != nil {
		var by []byte
		by, err = base.JSON.Marshal(sessionParam)
		if err != nil {
			return err
		}
		value = string(by)
	}
	session.Values["session"] = value
	err = session.Save(r, w)
	return err
}
func GetSessionParam(r *http.Request) *base.SessionParam {
	session, _ := sessionStore.Get(r, "teamide-toolbox")
	value, ok := session.Values["session"]
	sessionParam := &base.SessionParam{}
	if ok && value != "" {
		base.JSON.Unmarshal([]byte(value.(string)), sessionParam)
	} else {
		sessionParam = &base.SessionParam{}
	}
	return sessionParam
}

func handleSession(rw http.ResponseWriter, r *http.Request) {
	sessionParam := GetSessionParam(r)
	outJSON(rw, sessionParam, nil)
}

func GetLoginUser(r *http.Request) *base.LoginUser {
	sessionParam := GetSessionParam(r)
	if sessionParam != nil && sessionParam.User != nil {
		return sessionParam.User
	}
	return nil
}
func ValidateLogin(w http.ResponseWriter, r *http.Request) bool {
	res := GetLoginUser(r) != nil
	if !res {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, NOT_LOGIN_RESULT)
	}
	return res
}

func handleLogin(rw http.ResponseWriter, r *http.Request) {
	request := &base.UserLoginReqeust{}
	response := make(map[string]interface{})
	err := ToBean(r, request)
	if err != nil {
		outJSON(rw, response, err)
		return
	}
	var loginUser *base.LoginUser

	for _, user := range config.Config.User {
		if user.Name == request.Name && user.Auth == request.Auth {
			loginUser = &base.LoginUser{
				Name: user.Name,
				Role: user.Role,
			}
		}
	}

	if loginUser == nil {
		outJSON(rw, response, errors.New("用户名或密码错误！"))
		return
	}

	SetSessionUser(rw, r, loginUser)

	outJSON(rw, response, nil)

}

func handleLogout(rw http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	SetSessionUser(rw, r, nil)

	outJSON(rw, response, nil)

}

package zookeeper

import (
	"base"
)

type getRequest struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

type getResponse struct {
	Data string `json:"data"`
}

func getWork(req interface{}) (res interface{}, err error) {
	request := &getRequest{}
	response := &getResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var zkService *ZKService
	zkService, err = getService(request.Url)
	if err != nil {
		return
	}
	var isEx bool
	isEx, err = zkService.Exists(request.Path)
	if err != nil {
		return
	}
	if isEx {
		var data []byte
		data, err = zkService.Get(request.Path)
		if err != nil {
			return
		}
		response.Data = string(data)
	}
	res = response
	return
}

type getChildrenRequest struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

type getChildrenResponse struct {
	Children []string `json:"children"`
}

func getChildrenWork(req interface{}) (res interface{}, err error) {
	request := &getChildrenRequest{}
	response := &getChildrenResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var zkService *ZKService
	zkService, err = getService(request.Url)
	if err != nil {
		return
	}
	var isEx bool
	isEx, err = zkService.Exists(request.Path)
	if err != nil {
		return
	}
	if isEx {
		var children []string
		children, err = zkService.GetChildren(request.Path)
		if err != nil {
			return
		}
		response.Children = children
	}
	res = response
	return
}

type saveRequest struct {
	Url  string `json:"url"`
	Path string `json:"path"`
	Data string `json:"data"`
}

type saveResponse struct {
}

func saveWork(req interface{}) (res interface{}, err error) {
	request := &saveRequest{}
	response := &saveResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var zkService *ZKService
	zkService, err = getService(request.Url)
	if err != nil {
		return
	}
	var isEx bool
	isEx, err = zkService.Exists(request.Path)
	if err != nil {
		return
	}
	if isEx {
		err = zkService.SetData(request.Path, []byte(request.Data))
	} else {
		err = zkService.CreateIfNotExists(request.Path, []byte(request.Data))
	}
	res = response
	return
}

type deleteRequest struct {
	Url  string `json:"url"`
	Path string `json:"path"`
}

type deleteResponse struct {
}

func deleteWork(req interface{}) (res interface{}, err error) {
	request := &deleteRequest{}
	response := &deleteResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var zkService *ZKService
	zkService, err = getService(request.Url)
	if err != nil {
		return
	}
	var isEx bool
	isEx, err = zkService.Exists(request.Path)
	if err != nil {
		return
	}
	if isEx {
		err = zkService.Delete(request.Path)
		if err != nil {
			return
		}
	}
	res = response
	return
}

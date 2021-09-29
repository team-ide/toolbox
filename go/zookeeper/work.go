package zookeeper

import (
	"base"
)

type getRequest struct {
	Address string `json:"address"`
	Path    string `json:"path"`
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
	var service *ZKService
	service, err = getService(request.Address)
	if err != nil {
		return
	}
	var isEx bool
	isEx, err = service.Exists(request.Path)
	if err != nil {
		return
	}
	if isEx {
		var data []byte
		data, err = service.Get(request.Path)
		if err != nil {
			return
		}
		response.Data = string(data)
	}
	res = response
	return
}

type getChildrenRequest struct {
	Address string `json:"address"`
	Path    string `json:"path"`
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
	var service *ZKService
	service, err = getService(request.Address)
	if err != nil {
		return
	}
	var isEx bool
	isEx, err = service.Exists(request.Path)
	if err != nil {
		return
	}
	if isEx {
		var children []string
		children, err = service.GetChildren(request.Path)
		if err != nil {
			return
		}
		response.Children = children
	}
	res = response
	return
}

type saveRequest struct {
	Address string `json:"address"`
	Path    string `json:"path"`
	Data    string `json:"data"`
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
	var service *ZKService
	service, err = getService(request.Address)
	if err != nil {
		return
	}
	var isEx bool
	isEx, err = service.Exists(request.Path)
	if err != nil {
		return
	}
	if isEx {
		err = service.SetData(request.Path, []byte(request.Data))
	} else {
		err = service.CreateIfNotExists(request.Path, []byte(request.Data))
	}
	res = response
	return
}

type deleteRequest struct {
	Address string `json:"address"`
	Path    string `json:"path"`
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
	var service *ZKService
	service, err = getService(request.Address)
	if err != nil {
		return
	}
	var isEx bool
	isEx, err = service.Exists(request.Path)
	if err != nil {
		return
	}
	if isEx {
		err = service.Delete(request.Path)
		if err != nil {
			return
		}
	}
	res = response
	return
}

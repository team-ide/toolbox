package redis

import (
	"base"
)

type getRequest struct {
	Address string `json:"address"`
	Auth    string `json:"auth"`
	Cluster bool   `json:"cluster"`
	Key     string `json:"key"`
}

type getResponse struct {
	Value string `json:"value"`
}

func getWork(req interface{}) (res interface{}, err error) {
	request := &getRequest{}
	response := &getResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service Service
	service, err = getService(request.Address, request.Auth, request.Cluster)
	if err != nil {
		return
	}
	var value string
	value, err = service.Get(request.Key)
	if err != nil {
		return
	}
	response.Value = value
	res = response
	return
}

type keysRequest struct {
	Address string `json:"address"`
	Auth    string `json:"auth"`
	Cluster bool   `json:"cluster"`
	Pattern string `json:"pattern"`
	Size    int    `json:"size"`
}

type keysResponse struct {
	Keys  []string `json:"keys"`
	Count int      `json:"count"`
}

func keysWork(req interface{}) (res interface{}, err error) {
	request := &keysRequest{}
	response := &keysResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service Service
	service, err = getService(request.Address, request.Auth, request.Cluster)
	if err != nil {
		return
	}
	var count int
	var keys []string
	count, keys, err = service.Keys(request.Pattern, request.Size)
	if err != nil {
		return
	}
	response.Keys = keys
	response.Count = count
	res = response
	return
}

type saveRequest struct {
	Address string `json:"address"`
	Auth    string `json:"auth"`
	Cluster bool   `json:"cluster"`
	Key     string `json:"key"`
	Value   string `json:"value"`
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
	var service Service
	service, err = getService(request.Address, request.Auth, request.Cluster)
	if err != nil {
		return
	}
	err = service.Set(request.Key, request.Value)
	if err != nil {
		return
	}
	res = response
	return
}

type deleteRequest struct {
	Address string `json:"address"`
	Auth    string `json:"auth"`
	Cluster bool   `json:"cluster"`
	Key     string `json:"key"`
}

type deleteResponse struct {
	Count int `json:"count"`
}

func deleteWork(req interface{}) (res interface{}, err error) {
	request := &deleteRequest{}
	response := &deleteResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service Service
	service, err = getService(request.Address, request.Auth, request.Cluster)
	if err != nil {
		return
	}
	var count int
	count, err = service.Del(request.Key)
	response.Count = count
	res = response
	if err != nil {
		return
	}
	return
}

type deletePatternRequest struct {
	Address string `json:"address"`
	Auth    string `json:"auth"`
	Cluster bool   `json:"cluster"`
	Pattern string `json:"pattern"`
}

type deletePatternResponse struct {
	Count int `json:"count"`
}

func deletePatternWork(req interface{}) (res interface{}, err error) {
	request := &deletePatternRequest{}
	response := &deletePatternResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var service Service
	service, err = getService(request.Address, request.Auth, request.Cluster)
	if err != nil {
		return
	}
	var count int
	count, err = service.DelPattern(request.Pattern)
	response.Count = count
	res = response
	if err != nil {
		return
	}
	return
}

package zookeeper

import "base"

type getRequest struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

type getResponse struct {
	Value string `json:"value"`
}

func getWork(req interface{}) (res interface{}, err error) {
	request := &getRequest{}
	err = base.ToBean(res.([]byte), request)
	if err != nil {

	}
	return res, err
}

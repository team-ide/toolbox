package web

import (
	"base"
	"io"
	"net/http"
	"net/url"
	"os"
	"resource"
)

func downloadResource(path string, rw http.ResponseWriter, r *http.Request) {

	absPath, err := resource.GetFileAbsPath(path)
	if err != nil {
		outJSON(rw, nil, err)
		return
	}
	exists, err := base.PathExists(absPath)

	if err != nil {
		outJSON(rw, nil, err)
		return
	}
	if !exists {
		rw.WriteHeader(404)
		return
	}

	fi, err := os.Open(absPath)
	if err != nil {
		outJSON(rw, nil, err)
		return
	}
	defer fi.Close()
	stat, _ := os.Stat(absPath)
	fileName := stat.Name()
	fileName = url.QueryEscape(fileName) // 防止中文乱码
	rw.Header().Set("Content-Type", "application/octet-stream")
	rw.Header().Set("content-disposition", "attachment; filename=\""+fileName+"\"")
	_, err = io.Copy(rw, fi)
	if err != nil {
		outJSON(rw, nil, err)
		return
	}
	return
}

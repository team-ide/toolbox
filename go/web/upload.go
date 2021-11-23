package web

import (
	"base"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"resource"
)

func uploadResource(path string, rw http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)
	//接收客户端传来的文件 file 与客户端保持一致
	file, handler, err := r.FormFile("file")

	if err != nil {
		outJSON(rw, nil, err)
		return
	}
	filePath := path + string(os.PathSeparator) + handler.Filename
	fileAbsPath, err := resource.GetFileAbsPath(filePath)
	exists, err := base.PathExists(fileAbsPath)

	if err != nil {
		outJSON(rw, nil, err)
		return
	}

	if exists {
		outJSON(rw, nil, errors.New(fmt.Sprint("文件[", handler.Filename, "]已存在！")))
		return
	}
	fileAbsDir := filepath.Dir(fileAbsPath)
	exists, _ = base.PathExists(fileAbsPath)
	if !exists {
		os.MkdirAll(fileAbsDir, os.ModeDir)
	}

	f, err := os.OpenFile(fileAbsPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		outJSON(rw, nil, err)
		return
	}
	defer f.Close()

	io.Copy(f, file)

	fileRelPath, err := resource.GetFileRelPath(filePath)
	outJSON(rw, fileRelPath, nil)
	return
}

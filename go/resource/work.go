package resource

import (
	"base"
	"config"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Name    string `json:"name"`
	RelPath string `json:"relPath"`
	Size    int64  `json:"size"`
	IsDir   bool   `json:"isDir"`
}

// 获取文件绝对路径
func GetFileAbsPath(path string) (filePath string, err error) {
	ResourceDirPath := config.ResourceDirPath
	dirPath := filepath.Dir(ResourceDirPath + path + string(os.PathSeparator) + "/ss")
	if strings.Index(dirPath+string(os.PathSeparator), ResourceDirPath) == 0 {
		filePath = dirPath
	} else {
		err = errors.New(fmt.Sprint("路径[", path, "]", "不在资源文件目录中！"))
		return
	}
	return
}

// 获取文件相对路径
func GetFileRelPath(path string) (fileRelPath string, err error) {
	fileRelPath, err = GetFileAbsPath(path)
	if err != nil {
		return
	}
	ResourceDirPath := config.ResourceDirPath
	if fileRelPath == ResourceDirPath || fileRelPath+string(os.PathSeparator) == ResourceDirPath {
		fileRelPath = ""
		return
	}
	fileRelPath = strings.ReplaceAll(fileRelPath, config.ResourceDirPath, "")
	return
}

func validatePath(path string) (fileAbsPath string, err error) {
	fileAbsPath, err = GetFileAbsPath(path)
	if err != nil {
		return
	}
	var exists bool
	exists, err = base.PathExists(fileAbsPath)
	if err != nil {
		return
	}
	if !exists {
		err = errors.New(fmt.Sprint("路径[", path, "]", "不存在！"))
		return
	}
	return
}

type listRequest struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type listResponse struct {
	Files []FileInfo `json:"files"`
}

func listWork(req interface{}) (res interface{}, err error) {
	request := &listRequest{}
	response := &listResponse{}
	err = base.ToBean(req.([]byte), request)
	if err != nil {
		return
	}
	var absPath string
	absPath, err = validatePath(request.Path)
	if err != nil {
		return
	}
	var dirs []os.DirEntry
	dirs, err = os.ReadDir(absPath)
	if err != nil {
		return
	}
	fileInfos := []FileInfo{}
	for _, dir := range dirs {
		fileInfo := FileInfo{
			Name:  dir.Name(),
			IsDir: dir.IsDir(),
		}
		fileAbsPath := request.Path + "/" + fileInfo.Name
		fileAbsPath, err = GetFileAbsPath(fileAbsPath)
		if err != nil {
			return
		}
		fileInfo.RelPath, err = GetFileRelPath(request.Path + "/" + fileInfo.Name)
		if err != nil {
			return
		}
		if !fileInfo.IsDir {
			var file fs.FileInfo
			file, err = os.Stat(fileAbsPath)
			if err != nil {
				return
			}
			fileInfo.Size = file.Size()
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	response.Files = fileInfos
	res = response
	return
}

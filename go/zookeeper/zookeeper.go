package zookeeper

import (
	"errors"
	"strings"
	"time"

	"github.com/go-zookeeper/zk"
)

//注册处理器在线信息等
type ZKService struct {
	conn      *zk.Conn        //zk连接
	connEvent <-chan zk.Event // zk事件通知管道
}

func CreateZKService(server string) (*ZKService, error) {
	service := &ZKService{}
	err := service.init(server)
	return service, err
}

func (service *ZKService) init(server string) error {
	var err error
	service.conn, service.connEvent, err = zk.Connect(strings.Split(server, ","), time.Second*3)
	return err
}

//创建节点
func (service *ZKService) Create(path string, data []byte, mode int32) (err error) {
	isExist, err := service.Exists(path)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("node:" + path + " already exists")
	}
	if strings.LastIndex(path, "/") > 0 {
		parentPath := path[0:strings.LastIndex(path, "/")]
		err = service.CreateIfNotExists(parentPath)
		if err != nil {
			return err
		}
	}
	if _, err = service.conn.Create(path, data, mode, zk.WorldACL(zk.PermAll)); err != nil {
		if err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

func (service *ZKService) SetData(path string, data []byte) (err error) {
	isExist, state, err := service.conn.Exists(path)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New("node:" + path + " not exists")
	}
	if _, err = service.conn.Set(path, data, state.Aversion); err != nil {
		if err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

//一层层检查，如果不存在则创建父节点
func (service *ZKService) CreateIfNotExists(path string) (err error) {
	isExist, err := service.Exists(path)
	if err != nil {
		return err
	}
	if isExist {
		return nil
	}
	if strings.LastIndex(path, "/") > 0 {
		parentPath := path[0:strings.LastIndex(path, "/")]
		err = service.CreateIfNotExists(parentPath)
		if err != nil {
			return err
		}
	}
	if _, err = service.conn.Create(path, []byte{}, 0, zk.WorldACL(zk.PermAll)); err != nil {
		if err != zk.ErrNodeExists {
			return err
		}
	}
	return nil
}

//判断节点是否存在
func (service *ZKService) Exists(path string) (isExist bool, err error) {
	isExist, _, err = service.conn.Exists(path)
	return isExist, err
}

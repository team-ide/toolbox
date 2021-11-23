package config

import (
	"base"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Server    *server     `toml:"server"`
	User      []user      `toml:"user"`
	Redis     []redis     `toml:"redis"`
	Zookeeper []zookeeper `toml:"zookeeper"`
	Kafka     []kafka     `toml:"kafka"`
	Database  []database  `toml:"database"`
	Resource  resource    `toml:"resource"`
}

type server struct {
	Host    string
	Port    int
	Context string
	Data    string
}

type user struct {
	Name string
	Auth string
	Role int
}

type redis struct {
	Address string `json:"address"`
	Auth    string `json:"auth"`
}

type zookeeper struct {
	Address string `json:"address"`
}

type kafka struct {
	Address string `json:"address"`
}

type database struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type resource struct {
	Dir string `json:"dir"`
}

var (
	Config          *TomlConfig
	ResourceDirPath string
)

func init() {

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := path + "/./config.toml"

	exists, err := base.PathExists(filePath)
	if exists && err == nil {
		// fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &Config); err != nil {
			panic(err)
		}
	}
	Config = formatConfig(Config)
}

//格式化配置，填充默认值
func formatConfig(config *TomlConfig) *TomlConfig {
	if config == nil {
		config = &TomlConfig{}
	}
	if config.Server == nil {
		config.Server = &server{}
	}
	if config.Server.Host == "" {
		config.Server.Host = "0.0.0.0"
	}
	if config.Server.Port == 0 {
		config.Server.Port = 18000
	}
	if config.Server.Context == "" {
		config.Server.Context = "/toolbox"
	}
	if config.Resource.Dir == "" {
		config.Resource.Dir = "./data/resource/dir"
	}
	if len(config.User) == 0 {
		panic("未配置登录用户")
	}

	resourceDirPath := config.Resource.Dir
	if strings.Index(resourceDirPath, "/") != 0 {
		rootPath, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		resourceDirPath = rootPath + "/" + resourceDirPath
	}

	exists, err := base.PathExists(resourceDirPath)
	if err != nil {
		panic(err)
	}
	if !exists {
		err = os.MkdirAll(resourceDirPath, os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	ResourceDirPath = filepath.Dir(resourceDirPath+string(os.PathSeparator)+"xx") + string(os.PathSeparator)
	return config
}

func GetFromSystem(key string) string {
	return os.Getenv(key)
}

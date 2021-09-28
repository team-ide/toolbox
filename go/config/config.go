package config

import (
	"base"
	"os"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Server    *server     `toml:"server"`
	User      []user      `toml:"user"`
	Redis     []redis     `toml:"redis"`
	Zookeeper []zookeeper `toml:"zookeeper"`
	Kafka     []kafka     `toml:"kafka"`
	Database  []database  `toml:"database"`
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

var (
	Config *TomlConfig
)

func init() {
	filePath := "./config.toml"
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
	if len(config.User) == 0 {
		panic("未配置登录用户")
	}
	return config
}

func GetFromSystem(key string) string {
	return os.Getenv(key)
}

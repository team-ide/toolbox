package config

import (
	"base"
	"os"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Server *server `toml:"server"`
}

type server struct {
	Host    string
	Port    int
	Context string
	Data    string
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
	return config
}

func GetFromSystem(key string) string {
	return os.Getenv(key)
}

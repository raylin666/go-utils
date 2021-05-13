package config

import (
	"go-server/config/autoload"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configs = new(Config)

type Config struct {
	Environment string                       `yaml:"Environment"`
	App         autoload.App                 `yaml:"App"`
	Http        autoload.Http                `yaml:"Http"`
	Database    map[string]autoload.Database `yaml:"Database"`
	Redis       map[string]autoload.Redis	 `yaml:"Redis"`
	Logs 		autoload.Logs				 `yaml:"Logs"`
}

// 初始化加载配置文件
func InitAutoloadConfig(envFileName string)  {
	cYaml, err := ioutil.ReadFile(envFileName)
	if err != nil {
		panic(err)
	}

	_ = yaml.Unmarshal(cYaml, &configs)
}

// 获取配置项
func Get() *Config {
	return configs
}

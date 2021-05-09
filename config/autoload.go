package config

import (
	"github.com/raylin666/go-gin-api/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	configs = new(Config)
)

type Config struct {
	*config.Config
	Grpc struct{
		System struct{
			Network string `yaml:"Network"`
			Host 	string `yaml:"Host"`
			Port    uint16 `yaml:"Port"`
		} `yaml:"System"`
	} `yaml:"Grpc"`
}

func InitConfig(YmlEnvFileName string) {
	cYaml, err := ioutil.ReadFile(YmlEnvFileName)
	if err != nil {
		panic(err)
	}

	configs.Config = config.Get()
	_ = yaml.Unmarshal(cYaml, &configs)
}

// 获取配置项
func Get() *Config {
	return configs
}

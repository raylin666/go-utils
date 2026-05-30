// Package config 提供 YAML 配置文件加载功能
// 功能特性：
//   - 从文件路径加载 YAML 配置
//   - 从字符串加载 YAML 配置
// 使用场景：
//   - 应用配置文件加载
//   - 配置解析
package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadYaml 从指定路径加载 YAML 文件并解析到目标结构体
// 功能说明：
//   - 读取 YAML 文件内容
//   - 解析 YAML 内容到目标结构体
//
// 参数：
//   - path: YAML 文件路径
//   - out: 目标结构体指针（必须是指针类型）
//
// 返回值：
//   - nil: 加载成功
//   - error: 加载失败时的错误信息
//
// 使用示例：
//   type AppConfig struct {
//       Server struct {
//           Port int `yaml:"port"`
//       } `yaml:"server"`
//   }
//   var cfg AppConfig
//   err := config.LoadYaml("config.yaml", &cfg)
//   if err != nil {
//       log.Fatal(err)
//   }
func LoadYaml(path string, out interface{}) error {
	yamlFileBytes, readErr := os.ReadFile(path)
	if readErr != nil {
		return readErr
	}
	err := yaml.Unmarshal(yamlFileBytes, out)
	if err != nil {
		return errors.New("Cannot resolve [" + path + "] -- " + err.Error())
	}
	return nil
}

// LoadYamlByString 从 YAML 字符串解析到目标结构体
// 功能说明：
//   - 解析 YAML 格式的字符串到目标结构体
//   - 适用于从环境变量或网络获取的 YAML 内容
//
// 参数：
//   - yamlStr: YAML 格式的字符串
//   - out: 目标结构体指针（必须是指针类型）
//
// 返回值：
//   - nil: 解析成功
//   - error: 解析失败时的错误信息
//
// 使用示例：
//   yamlStr := "server:\n  port: 8080"
//   var cfg AppConfig
//   err := config.LoadYamlByString(yamlStr, &cfg)
func LoadYamlByString(yamlStr string, out interface{}) error {
	return yaml.Unmarshal([]byte(yamlStr), out)
}
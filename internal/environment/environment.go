package environment

import (
	"flag"
	"fmt"
	"go-server/config"
	"go-server/internal/constant"
	"strings"
)

var (
	active Environment
)

var _ Environment = (*environment)(nil)

type environment struct {
	value string
}

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsTest() bool
	IsPre() bool
	IsProd() bool
}

// 获取当前环境
func GetEnvironment() Environment {
	return active
}

// 获取当前环境值
func (e *environment) Value() string {
	return e.value
}

// 是否开发环境
func (e *environment) IsDev() bool {
	return e.value == constant.EnvironmentDev
}

// 是否测试环境
func (e *environment) IsTest() bool {
	return e.value == constant.EnvironmentTest
}

// 是否预发布环境
func (e *environment) IsPre() bool {
	return e.value == constant.EnvironmentPre
}

// 是否生产环境
func (e *environment) IsProd() bool {
	return e.value == constant.EnvironmentProd
}

// 初始化环境
func InitEnvironment() {
	// go run main.go -env=prod
	env := flag.String("env", "", fmt.Sprintf("请输入运行环境:\n %s:开发环境\n %s:测试环境\n %s:预上线环境\n %s:正式环境\n", constant.EnvironmentDev, constant.EnvironmentTest, constant.EnvironmentPre, constant.EnvironmentProd))
	flag.Parse()

	switch strings.ToLower(strings.TrimSpace(*env)) {
	case constant.EnvironmentDev:
		active = &environment{value: constant.EnvironmentDev}
	case constant.EnvironmentTest:
		active = &environment{value: constant.EnvironmentTest}
	case constant.EnvironmentPre:
		active = &environment{value: constant.EnvironmentPre}
	case constant.EnvironmentProd:
		active = &environment{value: constant.EnvironmentProd}
	default:
		config_env := config.Get().Environment
		if config_env == constant.EnvironmentDev {
			fmt.Println("Warning: '-" + constant.EnvironmentDev + "' cannot be found, or it is illegal. The default '" + constant.EnvironmentDev + "' will be used.")
		}
		active = &environment{value: config_env}
	}
}

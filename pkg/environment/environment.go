package environment

import (
	"strings"
)

const (
	EnvironmentDev = "dev"
	EnvironmentTest = "test"
	EnvironmentPre = "pre"
	EnvironmentProd = "prod"
)

var _ Environment = (*environment)(nil)

// Environment 环境配置
type Environment interface {
	Value() string
	IsDev() bool
	IsTest() bool
	IsPre() bool
	IsProd() bool
}

type environment struct {
	value string
}

// NewEnvironment 创建应用环境
func NewEnvironment(envname string) Environment {
	envname = strings.ToLower(envname)
	switch envname {
	case EnvironmentDev:
	case EnvironmentTest:
	case EnvironmentPre:
	case EnvironmentProd:
	default:
		envname = EnvironmentDev
	}

	var env = new(environment)
	env.value = envname
	return env
}

// 获取当前环境值
func (e *environment) Value() string {
	return e.value
}

// 是否开发环境
func (e *environment) IsDev() bool {
	return e.value == EnvironmentDev
}

// 是否测试环境
func (e *environment) IsTest() bool {
	return e.value == EnvironmentTest
}

// 是否预发布环境
func (e *environment) IsPre() bool {
	return e.value == EnvironmentPre
}

// 是否生产环境
func (e *environment) IsProd() bool {
	return e.value == EnvironmentProd
}

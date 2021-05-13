package utils

import (
	"fmt"
	"os"
)

// 判断目录是否存在,不存在则创建
func CreateDirectory(dir string) {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				panic(fmt.Errorf("Fatal create %v directory %v \n", dir, err))
			} else {
				fmt.Printf("create %v directory success \n", dir)
			}
		}
	}
}

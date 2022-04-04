package ut

import (
	"os"
)

// CreateDirectory 判断目录是否存在,不存在则创建
func CreateDirectory(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.ModePerm)
			return err
		}
	}

	return nil
}

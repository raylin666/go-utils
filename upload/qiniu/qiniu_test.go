package qiniu

import (
	"fmt"
	"testing"
)

func initOptions() *Options {
	return New(
		"ph2rb2GYAJu94WzdZ0AB5uZFTK1GJSPCP0SAri9_",
		"u0okNS0wLPwEILtOLUcfQRXiZY7l7TltINs38hmg",
		"raylin666",
		"huanan",
		&Config{},
	)
}

// 下载文件 - 公开空间
func TestMakePublicURL(t *testing.T) {
	opts := initOptions()
	url := opts.MakePublicURL("cdn.ls331.com", "micro/2021/0928/9f4ccfa9-44b2-4558-adcb-bc4780ec4716.jpe")
	fmt.Println(url)
}

func TestGetFileInfo(t *testing.T) {
	opts := initOptions()
	f, err := opts.GetFileInfo("micro/2021/0928/9f4ccfa9-44b2-4558-adcb-bc4780ec4716.jpe")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(f)
}

func TestListFiles(t *testing.T) {
	opts := initOptions()
	marker, entries, err := opts.ListFiles("", "", "eyJjIjowLCJrIjoic3RvcmFnZS8yMDIxLzA0MjgvV1gyMDE5MDYwNy0xNDMzNDVAMnguanBnIn0=",20)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(marker)
	fmt.Println(entries)
}

func TestFetchWithoutKey(t *testing.T) {
	opts := initOptions()
	r, err := opts.FetchWithoutKey("http://devtools.qiniu.com/qiniu.png")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(r)
}
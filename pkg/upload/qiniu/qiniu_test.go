package qiniu

import (
	"fmt"
	"testing"
	"time"
)

func initQiniu() Qiniu {
	return New(
		"ph2rb2GYAJu94WzdZ0AB5uxxxxxxxxx0SAri9_",
		"u0okNS0wLPwEILtOLUcxxxxxxTltINs38hmg",
		"xxxxxx",
		"huanan",
		&Config{},
	)
}

// 下载文件 - 公开空间
func TestMakePublicURLv2(t *testing.T) {
	qiniu := initQiniu()
	url := qiniu.MakePublicURLv2("cdn.xxx.com", "micro/2021/0928/9f4ccfa9-44b2-4558-adcb-bc4780ec4716.jpe", nil)
	fmt.Println(url)

	/**
		cdn.xxx.com/micro/2021/0928/9f4ccfa9-44b2-4558-adcb-bc4780ec4716.jpe
	 */
}

func TestGetFileInfo(t *testing.T) {
	qiniu := initQiniu()
	f, err := qiniu.GetFileInfo("micro/2021/0928/9f4ccfa9-44b2-4558-adcb-bc4780ec4716.jpe")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(f)

	/**
		Hash:     FnAdCe_IWHgA6sDCD6zGd3kEO_S2
		Fsize:    207904
		PutTime:  16196013289516606
		MimeType: image/jpeg
		Type:     0
		Status:   0
	*/
}

func TestListFiles(t *testing.T) {
	qiniu := initQiniu()
	marker, entries, err := qiniu.ListFiles("", "", "",20)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(marker)
	fmt.Println(entries)
}

func TestFetchWithoutKey(t *testing.T) {
	qiniu := initQiniu()
	r, err := qiniu.FetchWithoutKey("http://devtools.qiniu.com/qiniu.png")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(r)
}

func TestCreateBucket(t *testing.T) {
	qiniu := initQiniu()
	err := qiniu.CreateBucket("test", "z0")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("success")
}

func TestBuckets(t *testing.T) {
	qiniu := initQiniu()
	buckets, err := qiniu.Buckets(true)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(buckets)
}

func TestCreateTimestampAntileechURL(t *testing.T) {
	qiniu := initQiniu()
	deadline := time.Now().Add(time.Second * 3600).Unix()
	url, err := qiniu.CreateTimestampAntileechURL("http://cdn.xxx.com/storage/2021/0428/1537765080414.jpg", "abc123", deadline)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(url)

	/**
		http://cdn.xxx.com/storage/2021/0428/1537765080414.jpg?sign=a6b7ff088049c06af70643277dd8857c&t=c31f1de8
	 */
}
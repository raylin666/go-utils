package qiniu

import (
	"bytes"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type basicConfig struct {
	AccessKey string
	SecretKey string
	Bucket    string
}

type Config struct {
	*storage.PutPolicy
	*storage.Config
}

type Options struct {
	basicConfig
	*Config

	PutRet      interface{}
	PutExtra    interface{}
}

func New(accessKey string, secretKey string, bucket string, zone string, c *Config) *Options {
	var (
		zoneObj = &storage.ZoneHuanan
	)

	if c.PutPolicy == nil {
		c.PutPolicy = new(storage.PutPolicy)
	}

	if c.Config == nil {
		c.Config = new(storage.Config)
	}

	switch zone {
	case "huadong":
		zoneObj = &storage.ZoneHuadong
		break
	case "huabei":
		zoneObj = &storage.ZoneHuabei
		break
	case "huanan":
		zoneObj = &storage.ZoneHuanan
		break
	case "beimei":
		zoneObj = &storage.ZoneBeimei
		break
	case "xinjiapo":
		zoneObj = &storage.ZoneXinjiapo
		break
	}

	opts := &Options{
		basicConfig{
			AccessKey: accessKey,
			SecretKey: secretKey,
			Bucket:    bucket,
		},
		c,
		storage.PutRet{},
		storage.PutExtra{},
	}

	opts.Config.Zone = zoneObj
	opts.PutPolicy.Scope = bucket
	return opts
}

func (opt *Options) GetMac() *qbox.Mac {
	return qbox.NewMac(opt.AccessKey, opt.SecretKey)
}

func (opt *Options) GetUploadToken() string {
	mac := opt.GetMac()
	return opt.PutPolicy.UploadToken(mac)
}

// GetFormUploader 上传文件/字节数组上传/数据流上传（表单方式）
func (opt *Options) GetFormUploader() *storage.FormUploader {
	return storage.NewFormUploader(opt.Config.Config)
}

// GetResumeUploader 文件分片上传/文件断点续传（表单方式）
func (opt *Options) GetResumeUploader() *storage.ResumeUploaderV2 {
	return storage.NewResumeUploaderV2(opt.Config.Config)
}

// FormUploaderPutFile 文件上传
func (opt *Options) FormUploaderPutFile(localPathFile string, storagePathFile string) (interface{}, error) {
	var (
		extra storage.PutExtra
	)

	if _, ok := opt.PutExtra.(storage.PutExtra); ok {
		extra = opt.PutExtra.(storage.PutExtra)
	}

	err := opt.GetFormUploader().PutFile(context.Background(), &opt.PutRet, opt.GetUploadToken(), storagePathFile, localPathFile, &extra)
	if err != nil {
		return nil, err
	}

	return opt.PutRet, nil
}

// FormUploaderPut 字节数组上传/数据流上传
func (opt *Options) FormUploaderPut(data []byte, storagePathFile string) (interface{}, error) {
	var (
		extra storage.PutExtra
	)

	if _, ok := opt.PutExtra.(storage.PutExtra); ok {
		extra = opt.PutExtra.(storage.PutExtra)
	}

	dataLen := int64(len(data))
	err := opt.GetFormUploader().Put(context.Background(), &opt.PutRet, opt.GetUploadToken(), storagePathFile, bytes.NewReader(data), dataLen, &extra)
	if err != nil {
		return nil, err
	}

	return opt.PutRet, nil
}

// ResumeUploaderPutFile 文件分片上传/文件断点续传
func (opt *Options) ResumeUploaderPutFile(localPathFile string, storagePathFile string) (interface{}, error) {
	var (
		extra storage.RputV2Extra
	)

	if _, ok := opt.PutExtra.(storage.RputV2Extra); ok {
		extra = opt.PutExtra.(storage.RputV2Extra)
	}

	err := opt.GetResumeUploader().PutFile(context.Background(), &opt.PutRet, opt.GetUploadToken(), storagePathFile, localPathFile, &extra)
	if err != nil {
		return nil, err
	}

	return opt.PutRet, nil
}



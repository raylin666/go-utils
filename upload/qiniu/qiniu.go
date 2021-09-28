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

func (opt *Options) GetBucketManager() *storage.BucketManager {
	return storage.NewBucketManager(opt.GetMac(), opt.Config.Config)
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
func (opt *Options) FormUploaderPutFile(localFile string, key string) (interface{}, error) {
	var (
		extra storage.PutExtra
	)

	if _, ok := opt.PutExtra.(storage.PutExtra); ok {
		extra = opt.PutExtra.(storage.PutExtra)
	}

	err := opt.GetFormUploader().PutFile(context.Background(), &opt.PutRet, opt.GetUploadToken(), key, localFile, &extra)
	if err != nil {
		return nil, err
	}

	return opt.PutRet, nil
}

// FormUploaderPut 字节数组上传/数据流上传
func (opt *Options) FormUploaderPut(data []byte, key string) (interface{}, error) {
	var (
		extra storage.PutExtra
	)

	if _, ok := opt.PutExtra.(storage.PutExtra); ok {
		extra = opt.PutExtra.(storage.PutExtra)
	}

	dataLen := int64(len(data))
	err := opt.GetFormUploader().Put(context.Background(), &opt.PutRet, opt.GetUploadToken(), key, bytes.NewReader(data), dataLen, &extra)
	if err != nil {
		return nil, err
	}

	return opt.PutRet, nil
}

// ResumeUploaderPutFile 文件分片上传/文件断点续传
func (opt *Options) ResumeUploaderPutFile(localFile string, key string) (interface{}, error) {
	var (
		extra storage.RputV2Extra
	)

	if _, ok := opt.PutExtra.(storage.RputV2Extra); ok {
		extra = opt.PutExtra.(storage.RputV2Extra)
	}

	err := opt.GetResumeUploader().PutFile(context.Background(), &opt.PutRet, opt.GetUploadToken(), key, localFile, &extra)
	if err != nil {
		return nil, err
	}

	return opt.PutRet, nil
}

// MakePublicURL 下载文件 - 公开空间
func (opt *Options) MakePublicURL(domain string, key string) string {
	return storage.MakePublicURL(domain, key)
}

// MakePrivateURL 下载文件 - 私有空间
func (opt *Options) MakePrivateURL(domain string, key string, ttl int64) string {
	return storage.MakePrivateURL(opt.GetMac(), domain, key, ttl)
}

// GetFileInfo 获取文件信息
func (opt *Options) GetFileInfo(key string) (*storage.FileInfo, error) {
	fileInfo, err := opt.GetBucketManager().Stat(opt.Bucket, key)
	if err != nil {
		return nil, err
	}

	return &fileInfo, nil
}

// ChangeFileMimeType 修改文件MimeType
func (opt *Options) ChangeFileMimeType(key string, mimeType string) error {
	return opt.GetBucketManager().ChangeMime(opt.Bucket, key, mimeType)
}

// ChangeFileType 修改文件类型
func (opt *Options) ChangeFileType(key string, fileType int) error {
	return opt.GetBucketManager().ChangeType(opt.Bucket, key, fileType)
}

// FileMove 移动或重命名文件
func (opt *Options) Move(destBucket string, srcKey string, destKey string, force bool) error {
	return opt.GetBucketManager().Move(opt.Bucket, srcKey, destBucket, destKey, force)
}

// FileCopy 复制文件副本
func (opt *Options) Copy(destBucket string, srcKey string, destKey string, force bool) error {
	return opt.GetBucketManager().Copy(opt.Bucket, srcKey, destBucket, destKey, force)
}

// FileDelete 删除空间中的文件
func (opt *Options) Delete(key string) error {
	return opt.GetBucketManager().Delete(opt.Bucket, key)
}

// FileDeleteAfterDays 设置或更新文件的生存时间
func (opt *Options) DeleteAfterDays(key string, days int) error {
	return opt.GetBucketManager().DeleteAfterDays(opt.Bucket, key, days)
}

// ListFiles 获取指定前缀的文件列表
func (opt *Options) ListFiles(prefix string, delimiter string, marker string, limit int) (string, []storage.ListItem, error) {
	entries, _, nextMarker, hasNext, err := opt.GetBucketManager().ListFiles(opt.Bucket, prefix, delimiter, marker, limit)

	if hasNext {
		// 获取下个 marker
		marker = nextMarker
	}

	return marker, entries, err
}

// Fetch 抓取网络资源到空间 (指定保存的key)
func (opt *Options) Fetch(resURL string, key string) (storage.FetchRet, error) {
	return opt.GetBucketManager().Fetch(resURL, opt.Bucket, key)
}

// FetchWithoutKey 抓取网络资源到空间 (不指定保存的key，默认用文件hash作为文件名)
func (opt *Options) FetchWithoutKey(resURL string) (storage.FetchRet, error) {
	return opt.GetBucketManager().FetchWithoutKey(resURL, opt.Bucket)
}


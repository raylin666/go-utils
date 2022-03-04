package qiniu

import (
	"bytes"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/cdn"
	"github.com/qiniu/go-sdk/v7/storage"
	"net/url"
)

// 基础配置, 对内提供
type basicConfig struct {
	accessKey string
	secretKey string
	bucket    string
}

// 存储配置
type Config struct {
	*storage.PutPolicy
	*storage.Config
}

// 配置参数
type options struct {
	*basicConfig
	*Config
	putRet      interface{}
	putExtra    interface{}
}

var _ Qiniu = (*options)(nil)

type Qiniu interface {
	// GetMac 构建鉴权对象
	GetMac() *qbox.Mac
	// GetBucketManager 构建资源管理对象
	GetBucketManager() *storage.BucketManager
	// GetCdnManager 构建CDN资源管理对象
	GetCdnManager() *cdn.CdnManager
	// GetUploadToken 获取上传凭证
	GetUploadToken() string
	// GetFormUploader 上传文件/字节数组上传/数据流上传（表单方式）
	GetFormUploader() *storage.FormUploader
	// GetResumeUploader 文件分片上传/文件断点续传（表单方式）
	GetResumeUploader() *storage.ResumeUploaderV2
	// FormUploaderPutFile 文件上传
	FormUploaderPutFile(localFile string, key string) (interface{}, error)
	// FormUploaderPut 字节数组上传/数据流上传
	FormUploaderPut(data []byte, key string) (interface{}, error)
	// ResumeUploaderPutFile 文件分片上传/文件断点续传
	ResumeUploaderPutFile(localFile string, key string) (interface{}, error)
	// MakePublicURL 生成下载文件链接 - 公开空间
	MakePublicURL(domain string, key string) string
	// MakePublicURLv2 生成下载文件链接, 并且该方法确保 key 将会被 escape，并在 URL 后追加经过编码的查询参数 - 公开空间
	MakePublicURLv2(domain string, key string, query url.Values) string
	// MakePrivateURL 生成下载文件链接 - 私有空间
	MakePrivateURL(domain string, key string, ttl int64) string
	// MakePrivateURLv2 生成下载文件链接, 并且该方法确保 key 将会被 escape，并在 URL 后追加经过编码的查询参数 - 公开空间 - 私有空间
	MakePrivateURLv2(domain string, key string, ttl int64, query url.Values) string
	// GetFileInfo 获取文件信息
	GetFileInfo(key string) (*storage.FileInfo, error)
	// ChangeFileMimeType 修改文件MimeType
	ChangeFileMimeType(key string, mimeType string) error
	// ChangeFileType 修改文件类型
	ChangeFileType(key string, fileType int) error
	// Move 移动或重命名文件
	Move(destBucket string, srcKey string, destKey string, force bool) error
	// Copy 复制文件副本
	Copy(destBucket string, srcKey string, destKey string, force bool) error
	// Delete 删除空间中的文件
	Delete(key string) error
	// DeleteAfterDays 设置或更新文件的生存时间
	DeleteAfterDays(key string, days int) error
	// ListFiles 获取指定前缀的文件列表
	ListFiles(prefix string, delimiter string, marker string, limit int) (string, []storage.ListItem, error)
	// Fetch 抓取网络资源到空间 (指定保存的key)
	Fetch(resURL string, key string) (storage.FetchRet, error)
	// FetchWithoutKey 抓取网络资源到空间 (不指定保存的key，默认用文件hash作为文件名)
	FetchWithoutKey(resURL string) (storage.FetchRet, error)
	// UpdateObjectStatus 修改文件状态, 禁用和启用文件的可访问性
	UpdateObjectStatus(bucket string, key string, enable bool) error
	// RefreshUrls CDN 文件刷新, 单次请求链接不可以超过100个，如果超过，请分批发送请求
	RefreshUrls(urlsToRefresh []string) (result cdn.RefreshResp, err error)
	// CreateTimestampAntileechURL 构建时间戳防盗链访问链接
	CreateTimestampAntileechURL(urlStr string, encryptKey string, ttl int64) (antileechURL string, err error)
	// CreateBucket 创建一个七牛存储空间
	CreateBucket(bucket string, regionID storage.RegionID) error
	// DropBucket 删除七牛存储空间
	DropBucket(bucket string) error
	// Buckets 用来获取空间列表
	Buckets(shared bool) (buckets []string, err error)
}

// New 创建对象
func New(accessKey string, secretKey string, bucket string, zone string, c *Config) Qiniu {
	if c.PutPolicy == nil {
		c.PutPolicy = new(storage.PutPolicy)
	}

	if c.Config == nil {
		c.Config = new(storage.Config)
	}

	var opts = new(options)
	opts.basicConfig = &basicConfig{accessKey: accessKey, secretKey: secretKey, bucket: bucket}
	opts.Config = c
	opts.putRet = storage.PutRet{}
	opts.putExtra = storage.PutExtra{}
	opts.Config.Zone = opts.region(zone)
	opts.PutPolicy.Scope = bucket
	return opts
}

// region 配置区域
func (opts *options) region(zone string) *storage.Region {
	var zoneObj = &storage.ZoneHuanan
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

	return zoneObj
}

// GetMac 构建鉴权对象
func (opt *options) GetMac() *qbox.Mac {
	return qbox.NewMac(opt.accessKey, opt.secretKey)
}

// GetBucketManager 构建资源管理对象
func (opt *options) GetBucketManager() *storage.BucketManager {
	return storage.NewBucketManager(opt.GetMac(), opt.Config.Config)
}

// GetCdnManager 构建CDN资源管理对象
func (opt *options) GetCdnManager() *cdn.CdnManager {
	return cdn.NewCdnManager(opt.GetMac())
}

// GetUploadToken 获取上传凭证
func (opt *options) GetUploadToken() string {
	mac := opt.GetMac()
	return opt.PutPolicy.UploadToken(mac)
}

// GetFormUploader 上传文件/字节数组上传/数据流上传（表单方式）
func (opt *options) GetFormUploader() *storage.FormUploader {
	return storage.NewFormUploader(opt.Config.Config)
}

// GetResumeUploader 文件分片上传/文件断点续传（表单方式）
func (opt *options) GetResumeUploader() *storage.ResumeUploaderV2 {
	return storage.NewResumeUploaderV2(opt.Config.Config)
}

// FormUploaderPutFile 文件上传
func (opt *options) FormUploaderPutFile(localFile string, key string) (interface{}, error) {
	var (
		extra storage.PutExtra
	)

	if _, ok := opt.putExtra.(storage.PutExtra); ok {
		extra = opt.putExtra.(storage.PutExtra)
	}

	err := opt.GetFormUploader().PutFile(context.Background(), &opt.putRet, opt.GetUploadToken(), key, localFile, &extra)
	if err != nil {
		return nil, err
	}

	return opt.putRet, nil
}

// FormUploaderPut 字节数组上传/数据流上传
func (opt *options) FormUploaderPut(data []byte, key string) (interface{}, error) {
	var (
		extra storage.PutExtra
	)

	if _, ok := opt.putExtra.(storage.PutExtra); ok {
		extra = opt.putExtra.(storage.PutExtra)
	}

	dataLen := int64(len(data))
	err := opt.GetFormUploader().Put(context.Background(), &opt.putRet, opt.GetUploadToken(), key, bytes.NewReader(data), dataLen, &extra)
	if err != nil {
		return nil, err
	}

	return opt.putRet, nil
}

// ResumeUploaderPutFile 文件分片上传/文件断点续传
func (opt *options) ResumeUploaderPutFile(localFile string, key string) (interface{}, error) {
	var (
		extra storage.RputV2Extra
	)

	if _, ok := opt.putExtra.(storage.RputV2Extra); ok {
		extra = opt.putExtra.(storage.RputV2Extra)
	}

	err := opt.GetResumeUploader().PutFile(context.Background(), &opt.putRet, opt.GetUploadToken(), key, localFile, &extra)
	if err != nil {
		return nil, err
	}

	return opt.putRet, nil
}

// MakePublicURL 生成下载文件链接 - 公开空间
func (opt *options) MakePublicURL(domain string, key string) string {
	return storage.MakePublicURL(domain, key)
}

// MakePublicURLv2 生成下载文件链接, 并且该方法确保 key 将会被 escape，并在 URL 后追加经过编码的查询参数 - 公开空间
func (opt *options) MakePublicURLv2(domain string, key string, query url.Values) string {
	return storage.MakePublicURLv2WithQuery(domain, key, query)
}

// MakePrivateURL 生成下载文件链接 - 私有空间
func (opt *options) MakePrivateURL(domain string, key string, ttl int64) string {
	return storage.MakePrivateURL(opt.GetMac(), domain, key, ttl)
}

// MakePrivateURLv2 生成下载文件链接, 并且该方法确保 key 将会被 escape，并在 URL 后追加经过编码的查询参数 - 公开空间 - 私有空间
func (opt *options) MakePrivateURLv2(domain string, key string, ttl int64, query url.Values) string {
	return storage.MakePrivateURLv2WithQuery(opt.GetMac(), domain, key, query, ttl)
}

// GetFileInfo 获取文件信息
func (opt *options) GetFileInfo(key string) (*storage.FileInfo, error) {
	fileInfo, err := opt.GetBucketManager().Stat(opt.bucket, key)
	if err != nil {
		return nil, err
	}

	return &fileInfo, nil
}

// ChangeFileMimeType 修改文件MimeType
func (opt *options) ChangeFileMimeType(key string, mimeType string) error {
	return opt.GetBucketManager().ChangeMime(opt.bucket, key, mimeType)
}

// ChangeFileType 修改文件类型
func (opt *options) ChangeFileType(key string, fileType int) error {
	return opt.GetBucketManager().ChangeType(opt.bucket, key, fileType)
}

// Move 移动或重命名文件
func (opt *options) Move(destBucket string, srcKey string, destKey string, force bool) error {
	return opt.GetBucketManager().Move(opt.bucket, srcKey, destBucket, destKey, force)
}

// Copy 复制文件副本
func (opt *options) Copy(destBucket string, srcKey string, destKey string, force bool) error {
	return opt.GetBucketManager().Copy(opt.bucket, srcKey, destBucket, destKey, force)
}

// Delete 删除空间中的文件
func (opt *options) Delete(key string) error {
	return opt.GetBucketManager().Delete(opt.bucket, key)
}

// DeleteAfterDays 设置或更新文件的生存时间
func (opt *options) DeleteAfterDays(key string, days int) error {
	return opt.GetBucketManager().DeleteAfterDays(opt.bucket, key, days)
}

// ListFiles 获取指定前缀的文件列表
func (opt *options) ListFiles(prefix string, delimiter string, marker string, limit int) (string, []storage.ListItem, error) {
	entries, _, nextMarker, hasNext, err := opt.GetBucketManager().ListFiles(opt.bucket, prefix, delimiter, marker, limit)

	if hasNext {
		// 获取下个 marker
		marker = nextMarker
	}

	return marker, entries, err
}

// Fetch 抓取网络资源到空间 (指定保存的key)
func (opt *options) Fetch(resURL string, key string) (storage.FetchRet, error) {
	return opt.GetBucketManager().Fetch(resURL, opt.bucket, key)
}

// FetchWithoutKey 抓取网络资源到空间 (不指定保存的key，默认用文件hash作为文件名)
func (opt *options) FetchWithoutKey(resURL string) (storage.FetchRet, error) {
	return opt.GetBucketManager().FetchWithoutKey(resURL, opt.bucket)
}

// UpdateObjectStatus 修改文件状态, 禁用和启用文件的可访问性
func (opt *options) UpdateObjectStatus(bucket string, key string, enable bool) error {
	return opt.GetBucketManager().UpdateObjectStatus(bucket, key, enable)
}

// RefreshUrls CDN 文件刷新, 单次请求链接不可以超过100个，如果超过，请分批发送请求
func (opt *options) RefreshUrls(urlsToRefresh []string) (result cdn.RefreshResp, err error) {
	return opt.GetCdnManager().RefreshUrls(urlsToRefresh)
}

// CreateTimestampAntileechURL 构建时间戳防盗链访问链接
func (opt *options) CreateTimestampAntileechURL(urlStr string, encryptKey string, ttl int64) (antileechURL string, err error) {
	return cdn.CreateTimestampAntileechURL(urlStr, encryptKey, ttl)
}

// CreateBucket 创建一个七牛存储空间
func (opt *options) CreateBucket(bucket string, regionID storage.RegionID) error {
	return opt.GetBucketManager().CreateBucket(bucket, regionID)
}

// DropBucket 删除七牛存储空间
func (opt *options) DropBucket(bucket string) error {
	return opt.GetBucketManager().DropBucket(bucket)
}

// Buckets 用来获取空间列表
func (opt *options) Buckets(shared bool) (buckets []string, err error) {
	return opt.GetBucketManager().Buckets(shared)
}


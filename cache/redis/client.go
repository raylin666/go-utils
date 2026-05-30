package redis

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/raylin666/go-utils/cache/redis/cmd"
	"github.com/redis/go-redis/v9"

	ut "github.com/raylin666/go-utils"
)

var _ Client = (*client)(nil)

var (
	ErrAddrEmpty         = fmt.Errorf("Redis server address cannot be empty")
	ErrPortInvalid       = fmt.Errorf("Redis port must be in range 1-65535")
	ErrDBIndexInvalid    = fmt.Errorf("Redis database index must be in range 0-15")
	ErrPoolSizeInvalid   = fmt.Errorf("pool size must be greater than 0")
	ErrMaxRetriesInvalid = fmt.Errorf("max retries cannot be negative")
)

// Client Redis客户端接口，提供完整的Redis命令支持
// 功能分组：
//   - 连接管理：Close, Raw, Ping, HealthCheck, IsConnected
//   - 键操作：Keys, Del, Exists, Expire, TTL, Type, Rename, Move
//   - 字符串：Get, Set, SetEx, SetNX, SetXX, GetSet, GetRange, Append, StrLen, Incr, Decr
//   - 哈希：HGet, HSet, HGetAll, HDel, HExists, HIncrBy, HKeys, HVals, HLen
//   - 列表：LPush, RPush, LPop, RPop, LRange, LLen, LIndex, LRem, LSet, LTrim
//   - 集合：SAdd, SRem, SMembers, SIsMember, SCard, SInter, SUnion, SDiff
//   - 有序集合：ZAdd, ZRem, ZRange, ZRank, ZScore, ZCard, ZIncrBy, ZCount
//   - 位图：GetBit, SetBit, BitCount, BitOpAnd, BitOpOr, BitOpXor, BitPos
//   - HyperLogLog：PFAdd, PFCount, PFMerge
//   - 地理位置：GeoAdd, GeoRadius, GeoDist, GeoHash, GeoPos
//   - 事务：Watch, Pipeline, TxPipeline
//   - 发布订阅：Publish, PubSubChannels
//   - 集群：ClusterSlots, ClusterNodes, ClusterMeet, ClusterInfo
//   - 服务器管理：Info, DBSize, FlushDB, FlushAll, ConfigGet, ConfigSet
type Client interface {
	ut.HealthChecker

	// ========== 连接管理 ==========

	// Close 关闭Redis连接，释放资源
	Close() error

	// Raw 返回原生Redis客户端，供高级操作使用
	// 使用场景：需要使用原生Redis客户端的高级功能时
	Raw() *redis.Client

	// Ping 测试Redis连接是否正常
	// 返回值：PONG表示连接正常
	Ping(ctx context.Context) *cmd.StatusCmd

	// Command 获取Redis命令列表和详细信息
	Command(ctx context.Context) *cmd.CommandsInfoCmd

	// ClientGetName 获取当前客户端连接名称
	ClientGetName(ctx context.Context) *cmd.StringCmd

	// Echo 测试命令，返回输入的消息
	Echo(ctx context.Context, message interface{}) *cmd.StringCmd

	// ========== 键操作 ==========

	// Keys 查找所有符合给定模式的键
	// 参数：pattern - 键模式（如 "user:*"）
	// 注意：生产环境慎用，可能阻塞
	Keys(ctx context.Context, pattern string) *cmd.StringSliceCmd

	// SafeScanKeys 安全扫描所有符合给定模式的键（推荐替代Keys）
	// 功能说明：
	//   - 使用Scan命令迭代扫描，避免阻塞Redis
	//   - 适用于生产环境大量键扫描
	//   - 自动处理游标迭代
	//
	// 参数：
	//   - ctx: 上下文，用于超时控制
	//   - pattern: 键模式（如 "user:*"）
	//   - count: 每次迭代返回的键数量（建议100-1000）
	//
	// 返回值：
	//   - keys: 所有匹配的键列表
	//   - error: 扫描失败时的错误信息
	//
	// 使用示例：
	//   keys, err := client.SafeScanKeys(ctx, "user:*", 100)
	//   if err != nil {
	//       log.Printf("扫描失败: %v", err)
	//   }
	//
	// 性能建议：
	//   - 生产环境推荐使用此方法替代Keys
	//   - count参数建议设置为100-1000，避免单次返回过多键
	//   - 大量键扫描时建议使用context超时控制
	SafeScanKeys(ctx context.Context, pattern string, count int64) ([]string, error)

	// Dump 序列化给定键的值
	Dump(ctx context.Context, key string) *cmd.StringCmd

	// Get 获取键的值
	Get(ctx context.Context, key string) *cmd.StringCmd

	// Set 设置键值
	// 参数：expiration - 过期时间（0表示永不过期）
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.StatusCmd

	// SetEx 设置键值并指定过期时间（秒）
	SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.StatusCmd

	// SetNX 仅当键不存在时设置值
	// 返回值：true表示设置成功，false表示键已存在
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.BoolCmd

	// SetXX 仅当键存在时设置值
	// 返回值：true表示设置成功，false表示键不存在
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.BoolCmd

	// Del 删除键
	// 返回值：被删除键的数量
	Del(ctx context.Context, keys ...string) *cmd.IntCmd

	// GetRange 获取键值的指定范围子串
	GetRange(ctx context.Context, key string, start, end int64) *cmd.StringCmd

	// GetSet 设置新值并返回旧值
	GetSet(ctx context.Context, key string, value interface{}) *cmd.StringCmd

	// GetEx 获取键值并设置过期时间
	GetEx(ctx context.Context, key string, expiration time.Duration) *cmd.StringCmd

	// GetDel 获取键值并删除键
	GetDel(ctx context.Context, key string) *cmd.StringCmd

	// StrLen 获取键值的字符串长度
	StrLen(ctx context.Context, key string) *cmd.IntCmd

	// ========== 数值操作 ==========

	// Incr 键值自增1
	// 返回值：自增后的值
	Incr(ctx context.Context, key string) *cmd.IntCmd

	// Decr 键值自减1
	Decr(ctx context.Context, key string) *cmd.IntCmd

	// IncrBy 键值增加指定数值
	IncrBy(ctx context.Context, key string, value int64) *cmd.IntCmd

	// DecrBy 键值减少指定数值
	DecrBy(ctx context.Context, key string, decrement int64) *cmd.IntCmd

	// IncrByFloat 键值增加指定浮点数
	IncrByFloat(ctx context.Context, key string, value float64) *cmd.FloatCmd

	// ========== 过期时间 ==========

	// Expire 设置键的过期时间（秒）
	Expire(ctx context.Context, key string, expiration time.Duration) *cmd.BoolCmd

	// ExpireAt 设置键的过期时间戳
	ExpireAt(ctx context.Context, key string, tm time.Time) *cmd.BoolCmd

	// PExpire 设置键的过期时间（毫秒）
	PExpire(ctx context.Context, key string, expiration time.Duration) *cmd.BoolCmd

	// PExpireAt 设置键的过期时间戳（毫秒）
	PExpireAt(ctx context.Context, key string, tm time.Time) *cmd.BoolCmd

	// TTL 获取键的剩余过期时间（秒）
	TTL(ctx context.Context, key string) *cmd.DurationCmd

	// PTTL 获取键的剩余过期时间（毫秒）
	PTTL(ctx context.Context, key string) *cmd.DurationCmd

	// ========== 键检查 ==========

	// Exists 检查键是否存在
	// 返回值：存在的键数量
	Exists(ctx context.Context, keys ...string) *cmd.IntCmd

	// Unlink 异步删除键（非阻塞）
	Unlink(ctx context.Context, keys ...string) *cmd.IntCmd

	// Migrate 将键迁移到另一个Redis实例
	Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *cmd.StatusCmd

	// Move 将键移动到另一个数据库
	Move(ctx context.Context, key string, db int) *cmd.BoolCmd

	// ObjectRefCount 获取键值的引用计数
	ObjectRefCount(ctx context.Context, key string) *cmd.IntCmd

	// ObjectEncoding 获取键值的内部编码
	ObjectEncoding(ctx context.Context, key string) *cmd.StringCmd

	// ObjectIdleTime 获取键的空闲时间
	ObjectIdleTime(ctx context.Context, key string) *cmd.DurationCmd

	// RandomKey 随机返回一个键
	RandomKey(ctx context.Context) *cmd.StringCmd

	// Rename 重命名键
	Rename(ctx context.Context, key string, newkey string) *cmd.StatusCmd

	// RenameNX 仅当新键不存在时重命名
	RenameNX(ctx context.Context, key string, newkey string) *cmd.BoolCmd

	// Type 获取键的类型
	Type(ctx context.Context, key string) *cmd.StatusCmd

	// ========== 字符串操作 ==========

	// Append 追加值到键末尾
	// 返回值：追加后的字符串长度
	Append(ctx context.Context, key, value string) *cmd.IntCmd

	// MGet 批量获取多个键的值
	MGet(ctx context.Context, keys ...string) *cmd.SliceCmd

	// MSet 批量设置多个键值
	MSet(ctx context.Context, values ...interface{}) *cmd.StatusCmd

	// MSetNX 仅当所有键都不存在时批量设置
	MSetNX(ctx context.Context, values ...interface{}) *cmd.BoolCmd

	// ========== 位图操作 ==========

	// GetBit 获取键指定偏移量的位值
	GetBit(ctx context.Context, key string, offset int64) *cmd.IntCmd

	// SetBit 设置键指定偏移量的位值
	SetBit(ctx context.Context, key string, offset int64, value int) *cmd.IntCmd

	// BitCount 计算键中设置为1的位数
	BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *cmd.IntCmd

	// BitOpAnd 对多个键执行位AND操作
	BitOpAnd(ctx context.Context, destKey string, keys ...string) *cmd.IntCmd

	// BitOpOr 对多个键执行位OR操作
	BitOpOr(ctx context.Context, destKey string, keys ...string) *cmd.IntCmd

	// BitOpXor 对多个键执行位XOR操作
	BitOpXor(ctx context.Context, destKey string, keys ...string) *cmd.IntCmd

	// BitOpNot 对键执行位NOT操作
	BitOpNot(ctx context.Context, destKey string, key string) *cmd.IntCmd

	// BitPos 查找键中第一个设置为1或0的位位置
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *cmd.IntCmd

	// BitField 执行位域操作
	BitField(ctx context.Context, key string, args ...interface{}) *cmd.IntSliceCmd

	// SetArgs 使用SetArgs结构设置键值
	SetArgs(ctx context.Context, key string, value interface{}, args redis.SetArgs) *cmd.StatusCmd

	// ========== 扫描操作 ==========

	// Scan 扫描键空间
	// 参数：cursor - 游标，match - 匹配模式，count - 返回数量
	Scan(ctx context.Context, cursor uint64, match string, count int64) *cmd.ScanCmd

	// ScanType 扫描指定类型的键
	ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *cmd.ScanCmd

	// SScan 扫描集合元素
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *cmd.ScanCmd

	// HScan 扫描哈希字段
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *cmd.ScanCmd

	// ZScan 扫描有序集合成员
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *cmd.ScanCmd

	// ========== 哈希操作 ==========

	// HDel 删除哈希字段
	// 返回值：被删除字段的数量
	HDel(ctx context.Context, key string, fields ...string) *cmd.IntCmd

	// HExists 检查哈希字段是否存在
	HExists(ctx context.Context, key, field string) *cmd.BoolCmd

	// HGet 获取哈希字段值
	HGet(ctx context.Context, key, field string) *cmd.StringCmd

	// HGetAll 获取哈希所有字段和值
	HGetAll(ctx context.Context, key string) *cmd.StringStringMapCmd

	// HIncrBy 哈希字段值增加整数
	HIncrBy(ctx context.Context, key, field string, incr int64) *cmd.IntCmd

	// HIncrByFloat 哈希字段值增加浮点数
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *cmd.FloatCmd

	// HKeys 获取哈希所有字段名
	HKeys(ctx context.Context, key string) *cmd.StringSliceCmd

	// HLen 获取哈希字段数量
	HLen(ctx context.Context, key string) *cmd.IntCmd

	// HMGet 批量获取哈希多个字段值
	HMGet(ctx context.Context, key string, fields ...string) *cmd.SliceCmd

	// HSet 设置哈希字段值
	// 参数：values - 字段值对（如 "field1", "value1", "field2", "value2"）
	HSet(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd

	// HMSet 批量设置哈希多个字段值
	HMSet(ctx context.Context, key string, values ...interface{}) *cmd.BoolCmd

	// HSetNX 仅当哈希字段不存在时设置值
	HSetNX(ctx context.Context, key, field string, value interface{}) *cmd.BoolCmd

	// HVals 获取哈希所有字段值
	HVals(ctx context.Context, key string) *cmd.StringSliceCmd

	// HRandField 随机获取哈希字段名
	HRandField(ctx context.Context, key string, count int) *cmd.StringSliceCmd

	// HRandFieldWithValues 随机获取哈希字段名和值
	HRandFieldWithValues(ctx context.Context, key string, count int) *cmd.KeyValueSliceCmd

	// ========== 列表操作 ==========

	// BLPop 阻塞式左侧弹出列表元素
	// 参数：timeout - 阻塞超时时间（0表示永久阻塞）
	BLPop(ctx context.Context, timeout time.Duration, keys ...string) *cmd.StringSliceCmd

	// BRPop 阻塞式右侧弹出列表元素
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) *cmd.StringSliceCmd

	// BRPopLPush 阻塞式从源列表右侧弹出并推入目标列表左侧
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *cmd.StringCmd

	// LIndex 获取列表指定索引位置的元素
	LIndex(ctx context.Context, key string, index int64) *cmd.StringCmd

	// LInsert 在列表指定元素前/后插入新元素
	// 参数：op - "BEFORE"或"AFTER"
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) *cmd.IntCmd

	// LInsertBefore 在列表指定元素前插入新元素
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *cmd.IntCmd

	// LInsertAfter 在列表指定元素后插入新元素
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *cmd.IntCmd

	// LLen 获取列表长度
	LLen(ctx context.Context, key string) *cmd.IntCmd

	// LPop 从列表左侧弹出元素
	LPop(ctx context.Context, key string) *cmd.StringCmd

	// LPopCount 从列表左侧弹出多个元素
	LPopCount(ctx context.Context, key string, count int) *cmd.StringSliceCmd

	// LPos 查找列表中元素的位置
	LPos(ctx context.Context, key string, value string, a redis.LPosArgs) *cmd.IntCmd

	// LPosCount 查找列表中元素的多个位置
	LPosCount(ctx context.Context, key string, value string, count int64, a redis.LPosArgs) *cmd.IntSliceCmd

	// LPush 从列表左侧推入元素
	// 返回值：推入后的列表长度
	LPush(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd

	// LPushX 仅当列表存在时从左侧推入元素
	LPushX(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd

	// LRange 获取列表指定范围的元素
	// 参数：start, stop - 索引范围（负数表示从末尾开始）
	LRange(ctx context.Context, key string, start, stop int64) *cmd.StringSliceCmd

	// LRem 删除列表中指定值的元素
	// 参数：count - 删除数量（0表示删除所有匹配元素）
	LRem(ctx context.Context, key string, count int64, value interface{}) *cmd.IntCmd

	// LSet 设置列表指定索引位置的元素
	LSet(ctx context.Context, key string, index int64, value interface{}) *cmd.StatusCmd

	// LTrim 修剪列表，只保留指定范围的元素
	LTrim(ctx context.Context, key string, start int64, stop int64) *cmd.StatusCmd

	// RPop 从列表右侧弹出元素
	RPop(ctx context.Context, key string) *cmd.StringCmd

	// RPopCount 从列表右侧弹出多个元素
	RPopCount(ctx context.Context, key string, count int) *cmd.StringSliceCmd

	// RPopLPush 从源列表右侧弹出并推入目标列表左侧
	RPopLPush(ctx context.Context, source, destination string) *cmd.StringCmd

	// RPush 从列表右侧推入元素
	RPush(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd

	// RPushX 仅当列表存在时从右侧推入元素
	RPushX(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd

	// LMove 从源列表移动元素到目标列表
	// 参数：srcpos, destpos - "LEFT"或"RIGHT"
	LMove(ctx context.Context, source, destination, srcpos, destpos string) *cmd.StringCmd

	// BLMove 阻塞式从源列表移动元素到目标列表
	BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *cmd.StringCmd

	// ========== 集合操作 ==========

	// SAdd 添加集合成员
	// 返回值：成功添加的成员数量
	SAdd(ctx context.Context, key string, members ...interface{}) *cmd.IntCmd

	// SCard 获取集合成员数量
	SCard(ctx context.Context, key string) *cmd.IntCmd

	// SDiff 获取多个集合的差集
	SDiff(ctx context.Context, keys ...string) *cmd.StringSliceCmd

	// SDiffStore 将多个集合的差集存储到目标集合
	SDiffStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd

	// SInter 获取多个集合的交集
	SInter(ctx context.Context, keys ...string) *cmd.StringSliceCmd

	// SInterStore 将多个集合的交集存储到目标集合
	SInterStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd

	// SIsMember 检查元素是否是集合成员
	SIsMember(ctx context.Context, key string, member interface{}) *cmd.BoolCmd

	// SMIsMember 批量检查多个元素是否是集合成员
	SMIsMember(ctx context.Context, key string, members ...interface{}) *cmd.BoolSliceCmd

	// SMembers 获取集合所有成员
	SMembers(ctx context.Context, key string) *cmd.StringSliceCmd

	// SMembersMap 获取集合所有成员（以map形式返回）
	SMembersMap(ctx context.Context, key string) *cmd.StringStructMapCmd

	// SMove 将成员从源集合移动到目标集合
	SMove(ctx context.Context, source, destination string, member interface{}) *cmd.BoolCmd

	// SPop 随机弹出集合成员
	SPop(ctx context.Context, key string) *cmd.StringCmd

	// SPopN 随机弹出多个集合成员
	SPopN(ctx context.Context, key string, count int64) *cmd.StringSliceCmd

	// SRandMember 随机获取集合成员（不删除）
	SRandMember(ctx context.Context, key string) *cmd.StringCmd

	// SRandMemberN 随机获取多个集合成员（不删除）
	SRandMemberN(ctx context.Context, key string, count int64) *cmd.StringSliceCmd

	// SRem 删除集合成员
	// 返回值：成功删除的成员数量
	SRem(ctx context.Context, key string, members ...interface{}) *cmd.IntCmd

	// SUnion 获取多个集合的并集
	SUnion(ctx context.Context, keys ...string) *cmd.StringSliceCmd

	// SUnionStore 将多个集合的并集存储到目标集合
	SUnionStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd

	// ========== 有序集合操作 ==========

	// ZAddArgs 使用ZAddArgs结构添加有序集合成员
	ZAddArgs(ctx context.Context, key string, args redis.ZAddArgs) *cmd.IntCmd

	// ZAddArgsIncr 使用ZAddArgs结构增加有序集合成员分数
	ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *cmd.FloatCmd

	// ZAdd 添加有序集合成员
	// 参数：members - 成员和分数（如 redis.Z{Member: "user1", Score: 100}）
	ZAdd(ctx context.Context, key string, members ...redis.Z) *cmd.IntCmd

	// ZAddNX 仅当成员不存在时添加
	ZAddNX(ctx context.Context, key string, members ...redis.Z) *cmd.IntCmd

	// ZAddXX 仅当成员存在时更新分数
	ZAddXX(ctx context.Context, key string, members ...redis.Z) *cmd.IntCmd

	// ZCard 获取有序集合成员数量
	ZCard(ctx context.Context, key string) *cmd.IntCmd

	// ZCount 统计分数范围内的成员数量
	ZCount(ctx context.Context, key, min, max string) *cmd.IntCmd

	// ZLexCount 统计字典范围内的成员数量
	ZLexCount(ctx context.Context, key, min, max string) *cmd.IntCmd

	// ZIncrBy 增加有序集合成员分数
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *cmd.FloatCmd

	// ZInterStore 将多个有序集合的交集存储到目标集合
	ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *cmd.IntCmd

	// ZInter 获取多个有序集合的交集
	ZInter(ctx context.Context, store *redis.ZStore) *cmd.StringSliceCmd

	// ZInterWithScores 获取多个有序集合的交集（包含分数）
	ZInterWithScores(ctx context.Context, store *redis.ZStore) *cmd.ZSliceCmd

	// ZMScore 批量获取多个成员的分数
	ZMScore(ctx context.Context, key string, members ...string) *cmd.FloatSliceCmd

	// ZPopMax 弹出分数最高的成员
	ZPopMax(ctx context.Context, key string, count ...int64) *cmd.ZSliceCmd

	// ZPopMin 弹出分数最低的成员
	ZPopMin(ctx context.Context, key string, count ...int64) *cmd.ZSliceCmd

	// ZRangeArgs 使用ZRangeArgs结构获取有序集合成员
	ZRangeArgs(ctx context.Context, z redis.ZRangeArgs) *cmd.StringSliceCmd

	// ZRangeArgsWithScores 使用ZRangeArgs结构获取有序集合成员（包含分数）
	ZRangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *cmd.ZSliceCmd

	// ZRange 获取有序集合指定范围的成员（按分数升序）
	ZRange(ctx context.Context, key string, start, stop int64) *cmd.StringSliceCmd

	// ZRangeWithScores 获取有序集合指定范围的成员和分数
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *cmd.ZSliceCmd

	// ZRangeByScore 获取分数范围内的成员
	ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd

	// ZRangeByLex 获取字典范围内的成员
	ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd

	// ZRangeByScoreWithScores 获取分数范围内的成员和分数
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.ZSliceCmd

	// ZRangeStore 将有序集合范围存储到目标集合
	ZRangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *cmd.IntCmd

	// ZRank 获取成员的排名（按分数升序，从0开始）
	ZRank(ctx context.Context, key, member string) *cmd.IntCmd

	// ZRem 删除有序集合成员
	ZRem(ctx context.Context, key string, members ...interface{}) *cmd.IntCmd

	// ZRemRangeByRank 删除排名范围内的成员
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *cmd.IntCmd

	// ZRemRangeByScore 删除分数范围内的成员
	ZRemRangeByScore(ctx context.Context, key, min, max string) *cmd.IntCmd

	// ZRemRangeByLex 删除字典范围内的成员
	ZRemRangeByLex(ctx context.Context, key, min, max string) *cmd.IntCmd

	// ZRevRange 获取有序集合指定范围的成员（按分数降序）
	ZRevRange(ctx context.Context, key string, start, stop int64) *cmd.StringSliceCmd

	// ZRevRangeWithScores 获取有序集合指定范围的成员和分数（降序）
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *cmd.ZSliceCmd

	// ZRevRangeByScore 获取分数范围内的成员（降序）
	ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd

	// ZRevRangeByLex 获取字典范围内的成员（降序）
	ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd

	// ZRevRangeByScoreWithScores 获取分数范围内的成员和分数（降序）
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.ZSliceCmd

	// ZRevRank 获取成员的排名（按分数降序，从0开始）
	ZRevRank(ctx context.Context, key, member string) *cmd.IntCmd

	// ZScore 获取成员的分数
	ZScore(ctx context.Context, key, member string) *cmd.FloatCmd

	// ZUnion 获取多个有序集合的并集
	ZUnion(ctx context.Context, store redis.ZStore) *cmd.StringSliceCmd

	// ZUnionWithScores 获取多个有序集合的并集（包含分数）
	ZUnionWithScores(ctx context.Context, store redis.ZStore) *cmd.ZSliceCmd

	// ZUnionStore 将多个有序集合的并集存储到目标集合
	ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *cmd.IntCmd

	// ZRandMember 随机获取有序集合成员
	ZRandMember(ctx context.Context, key string, count int) *cmd.StringSliceCmd

	// ZRandMemberWithScores 随机获取有序集合成员和分数
	ZRandMemberWithScores(ctx context.Context, key string, count int) *cmd.ZSliceCmd

	// ZDiff 获取多个有序集合的差集
	ZDiff(ctx context.Context, keys ...string) *cmd.StringSliceCmd

	// ZDiffWithScores 获取多个有序集合的差集（包含分数）
	ZDiffWithScores(ctx context.Context, keys ...string) *cmd.ZSliceCmd

	// ZDiffStore 将多个有序集合的差集存储到目标集合
	ZDiffStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd

	// ========== HyperLogLog操作 ==========

	// PFAdd 添加元素到HyperLogLog
	// 返回值：1表示基数估计值改变，0表示未改变
	PFAdd(ctx context.Context, key string, els ...interface{}) *cmd.IntCmd

	// PFCount 获取HyperLogLog的基数估计值
	PFCount(ctx context.Context, keys ...string) *cmd.IntCmd

	// PFMerge 合并多个HyperLogLog
	PFMerge(ctx context.Context, dest string, keys ...string) *cmd.StatusCmd

	// ========== 服务器管理 ==========

	// BgRewriteAOF 异步重写AOF文件
	BgRewriteAOF(ctx context.Context) *cmd.StatusCmd

	// BgSave 异步保存数据到磁盘
	BgSave(ctx context.Context) *cmd.StatusCmd

	// ========== 地理位置 ==========

	// GeoAdd 添加地理位置
	// 参数：geoLocation - 经纬度和成员名称
	GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *cmd.IntCmd

	// GeoRadius 获取指定经纬度范围内的地理位置
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *cmd.GeoLocationCmd

	// GeoRadiusStore 将指定经纬度范围内的地理位置存储到目标集合
	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *cmd.IntCmd

	// GeoRadiusByMember 获取指定成员范围内的地理位置
	GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *cmd.GeoLocationCmd

	// GeoRadiusByMemberStore 将指定成员范围内的地理位置存储到目标集合
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *cmd.IntCmd

	// GeoSearch 搜索地理位置
	GeoSearch(ctx context.Context, key string, q *redis.GeoSearchQuery) *cmd.StringSliceCmd

	// GeoSearchLocation 搜索地理位置（包含坐标和距离）
	GeoSearchLocation(ctx context.Context, key string, q *redis.GeoSearchLocationQuery) *cmd.GeoSearchLocationCmd

	// GeoSearchStore 将搜索的地理位置存储到目标集合
	GeoSearchStore(ctx context.Context, key, store string, q *redis.GeoSearchStoreQuery) *cmd.IntCmd

	// GeoDist 计算两个地理位置之间的距离
	// 参数：unit - 单位（"m", "km", "mi", "ft"）
	GeoDist(ctx context.Context, key string, member1, member2, unit string) *cmd.FloatCmd

	// GeoHash 获取地理位置的Geohash值
	GeoHash(ctx context.Context, key string, members ...string) *cmd.StringSliceCmd

	// GeoPos 获取地理位置的经纬度
	GeoPos(ctx context.Context, key string, members ...string) *cmd.GeoPosCmd

	// ========== 事务和管道 ==========

	// Watch 监视键，用于事务
	// 参数：fn - 事务函数，keys - 监视的键
	Watch(ctx context.Context, fn func(tx *redis.Tx) error, keys ...string) error

	// Wait 等待指定数量的副本确认写操作
	Wait(ctx context.Context, numSlaves int, timeout time.Duration) *cmd.IntCmd

	// Pipeline 执行管道命令（批量发送命令，减少网络往返）
	// 参数：fn - 管道函数
	// 使用场景：需要执行多个命令且不依赖中间结果时
	Pipeline(ctx context.Context, fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error)

	// TxPipeline 执行事务管道命令（使用MULTI/EXEC包裹）
	TxPipeline(ctx context.Context, fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error)

	// ========== 客户端管理 ==========

	// ClientKill 杀死指定客户端连接
	ClientKill(ctx context.Context, ipPort string) *cmd.StatusCmd

	// ClientKillByFilter 按条件杀死客户端连接
	ClientKillByFilter(ctx context.Context, keys ...string) *cmd.IntCmd

	// ClientList 获取客户端连接列表
	ClientList(ctx context.Context) *cmd.StringCmd

	// ClientPause 暂停所有客户端命令
	ClientPause(ctx context.Context, dur time.Duration) *cmd.BoolCmd

	// ClientID 获取当前客户端连接ID
	ClientID(ctx context.Context) *cmd.IntCmd

	// ClientUnblock 解除阻塞客户端
	ClientUnblock(ctx context.Context, id int64) *cmd.IntCmd

	// ClientUnblockWithError 解除阻塞客户端并返回错误
	ClientUnblockWithError(ctx context.Context, id int64) *cmd.IntCmd

	// ========== 配置管理 ==========

	// ConfigGet 获取Redis配置参数
	ConfigGet(ctx context.Context, parameter string) *cmd.StringStringMapCmd

	// ConfigResetStat 重置统计信息
	ConfigResetStat(ctx context.Context) *cmd.StatusCmd

	// ConfigSet 设置Redis配置参数
	ConfigSet(ctx context.Context, parameter, value string) *cmd.StatusCmd

	// ConfigRewrite 将配置写入配置文件
	ConfigRewrite(ctx context.Context) *cmd.StatusCmd

	// ========== 数据库管理 ==========

	// DBSize 获取数据库键数量
	DBSize(ctx context.Context) *cmd.IntCmd

	// FlushAll 清空所有数据库（同步）
	FlushAll(ctx context.Context) *cmd.StatusCmd

	// FlushAllAsync 异步清空所有数据库
	FlushAllAsync(ctx context.Context) *cmd.StatusCmd

	// FlushDB 清空当前数据库（同步）
	FlushDB(ctx context.Context) *cmd.StatusCmd

	// FlushDBAsync 异步清空当前数据库
	FlushDBAsync(ctx context.Context) *cmd.StatusCmd

	// Info 获取Redis服务器信息
	// 参数：section - 信息类型（如 "server", "memory", "stats"）
	Info(ctx context.Context, section ...string) *cmd.StringCmd

	// LastSave 获取最后一次保存时间戳
	LastSave(ctx context.Context) *cmd.IntCmd

	// Save 同步保存数据到磁盘
	Save(ctx context.Context) *cmd.StatusCmd

	// Shutdown 关闭Redis服务器
	Shutdown(ctx context.Context) *cmd.StatusCmd

	// ShutdownSave 关闭Redis服务器并保存数据
	ShutdownSave(ctx context.Context) *cmd.StatusCmd

	// ShutdownNoSave 关闭Redis服务器不保存数据
	ShutdownNoSave(ctx context.Context) *cmd.StatusCmd

	// SlaveOf 设置主从复制
	SlaveOf(ctx context.Context, host, port string) *cmd.StatusCmd

	// SlowLogGet 获取慢查询日志
	SlowLogGet(ctx context.Context, num int64) *cmd.SlowLogCmd

	// Time 获取Redis服务器时间
	Time(ctx context.Context) *cmd.TimeCmd

	// DebugObject 获取键的调试信息
	DebugObject(ctx context.Context, key string) *cmd.StringCmd

	// ReadOnly 设置只读模式（用于从节点）
	ReadOnly(ctx context.Context) *cmd.StatusCmd

	// ReadWrite 设置读写模式
	ReadWrite(ctx context.Context) *cmd.StatusCmd

	// MemoryUsage 获取键的内存使用量
	MemoryUsage(ctx context.Context, key string, samples ...int) *cmd.IntCmd

	// ========== Lua脚本 ==========

	// Eval 执行Lua脚本
	// 参数：script - Lua脚本，keys - 键列表，args - 参数列表
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *cmd.Cmd

	// EvalSha 执行缓存的Lua脚本（通过SHA1）
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *cmd.Cmd

	// ScriptExists 检查脚本是否已缓存
	ScriptExists(ctx context.Context, hashes ...string) *cmd.BoolSliceCmd

	// ScriptFlush 清空脚本缓存
	ScriptFlush(ctx context.Context) *cmd.StatusCmd

	// ScriptKill 杀死正在执行的脚本
	ScriptKill(ctx context.Context) *cmd.StatusCmd

	// ScriptLoad 缓存脚本并返回SHA1
	ScriptLoad(ctx context.Context, script string) *cmd.StringCmd

	// ========== 发布订阅 ==========

	// Publish 发布消息到频道
	// 返回值：接收到消息的订阅者数量
	Publish(ctx context.Context, channel string, message interface{}) *cmd.IntCmd

	// PubSubChannels 获取活跃频道列表
	PubSubChannels(ctx context.Context, pattern string) *cmd.StringSliceCmd

	// PubSubNumSub 获取频道订阅者数量
	PubSubNumSub(ctx context.Context, channels ...string) *cmd.StringIntMapCmd

	// PubSubNumPat 获取模式订阅者数量
	PubSubNumPat(ctx context.Context) *cmd.IntCmd

	// ========== 集群操作 ==========

	// ClusterSlots 获取集群槽位信息
	ClusterSlots(ctx context.Context) *cmd.ClusterSlotsCmd

	// ClusterNodes 获取集群节点信息
	ClusterNodes(ctx context.Context) *cmd.StringCmd

	// ClusterMeet 连接集群节点
	ClusterMeet(ctx context.Context, host, port string) *cmd.StatusCmd

	// ClusterForget 移除集群节点
	ClusterForget(ctx context.Context, nodeID string) *cmd.StatusCmd

	// ClusterReplicate 设置节点为指定节点的副本
	ClusterReplicate(ctx context.Context, nodeID string) *cmd.StatusCmd

	// ClusterResetSoft 软重置集群节点
	ClusterResetSoft(ctx context.Context) *cmd.StatusCmd

	// ClusterResetHard 硬重置集群节点
	ClusterResetHard(ctx context.Context) *cmd.StatusCmd

	// ClusterInfo 获取集群信息
	ClusterInfo(ctx context.Context) *cmd.StringCmd

	// ClusterKeySlot 获取键的槽位
	ClusterKeySlot(ctx context.Context, key string) *cmd.IntCmd

	// ClusterGetKeysInSlot 获取槽位中的键
	ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *cmd.StringSliceCmd

	// ClusterCountFailureReports 获取节点故障报告数量
	ClusterCountFailureReports(ctx context.Context, nodeID string) *cmd.IntCmd

	// ClusterCountKeysInSlot 获取槽位中的键数量
	ClusterCountKeysInSlot(ctx context.Context, slot int) *cmd.IntCmd

	// ClusterDelSlots 删除槽位
	ClusterDelSlots(ctx context.Context, slots ...int) *cmd.StatusCmd

	// ClusterDelSlotsRange 删除槽位范围
	ClusterDelSlotsRange(ctx context.Context, min, max int) *cmd.StatusCmd

	// ClusterSaveConfig 保存集群配置
	ClusterSaveConfig(ctx context.Context) *cmd.StatusCmd

	// ClusterSlaves 获取节点的副本列表
	ClusterSlaves(ctx context.Context, nodeID string) *cmd.StringSliceCmd

	// ClusterFailover 执行故障转移
	ClusterFailover(ctx context.Context) *cmd.StatusCmd

	// ClusterAddSlots 添加槽位
	ClusterAddSlots(ctx context.Context, slots ...int) *cmd.StatusCmd

	// ClusterAddSlotsRange 添加槽位范围
	ClusterAddSlotsRange(ctx context.Context, min, max int) *cmd.StatusCmd
}

type Options struct {
	redis.Options
}

type client struct {
	client    *redis.Client
	options   *Options
	connected atomic.Bool
}

func NewClient(ctx context.Context, opt *Options) (Client, error) {
	if opt == nil {
		return nil, fmt.Errorf("Redis configuration options cannot be empty")
	}

	if opt.Addr == "" {
		return nil, ErrAddrEmpty
	}

	if opt.DB < 0 || opt.DB > 15 {
		return nil, ErrDBIndexInvalid
	}

	if opt.PoolSize < 0 {
		return nil, ErrPoolSizeInvalid
	}

	if opt.MaxRetries < 0 {
		return nil, ErrMaxRetriesInvalid
	}

	var c = new(client)
	c.options = opt
	c.client = redis.NewClient(&opt.Options)

	if err := c.Ping(ctx).Err(); err != nil {
		c.connected.Store(false)
		return nil, fmt.Errorf("Redis connection failed: %w", err)
	}

	c.connected.Store(true)
	return c, nil
}

func (c *client) HealthCheck(ctx context.Context) error {
	if c.client == nil {
		c.connected.Store(false)
		return fmt.Errorf("Redis client not initialized")
	}

	err := c.Ping(ctx).Err()
	if err != nil {
		c.connected.Store(false)
		return fmt.Errorf("Redis health check failed: %w", err)
	}

	c.connected.Store(true)
	return nil
}

func (c *client) IsConnected() bool {
	return c.connected.Load()
}

func (c *client) Close() error {
	if c.client != nil {
		err := c.client.Close()
		if err != nil {
			return err
		}
		c.connected.Store(false)
	}
	return nil
}

func (c *client) Raw() *redis.Client {
	return c.client
}

func (c *client) Ping(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Ping(ctx))
}

func (c *client) Command(ctx context.Context) *cmd.CommandsInfoCmd {
	return cmd.NewCommandsInfoCMD(c.client.Command(ctx))
}

func (c *client) ClientGetName(ctx context.Context) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ClientGetName(ctx))
}

func (c *client) Echo(ctx context.Context, message interface{}) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.Echo(ctx, message))
}

func (c *client) Keys(ctx context.Context, pattern string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.Keys(ctx, pattern))
}

func (c *client) SafeScanKeys(ctx context.Context, pattern string, count int64) ([]string, error) {
	var keys []string
	var cursor uint64 = 0

	for {
		iter := c.Scan(ctx, cursor, pattern, count)
		if iter.Err() != nil {
			return nil, fmt.Errorf("Redis scan failed: %w", iter.Err())
		}

		pageKeys, nextCursor := iter.Val()
		keys = append(keys, pageKeys...)

		if nextCursor == 0 {
			break
		}

		cursor = nextCursor

		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("Redis scan timeout: %w", ctx.Err())
		default:
		}
	}

	return keys, nil
}

func (c *client) Dump(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.Dump(ctx, key))
}

func (c *client) Get(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.Get(ctx, key))
}

func (c *client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Set(ctx, key, value, expiration))
}

func (c *client) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.SetEx(ctx, key, value, expiration))
}

func (c *client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.SetNX(ctx, key, value, expiration))
}

func (c *client) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.SetXX(ctx, key, value, expiration))
}

func (c *client) Del(ctx context.Context, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Del(ctx, keys...))
}

func (c *client) GetRange(ctx context.Context, key string, start, end int64) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.GetRange(ctx, key, start, end))
}

func (c *client) GetSet(ctx context.Context, key string, value interface{}) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.GetSet(ctx, key, value))
}

func (c *client) GetEx(ctx context.Context, key string, expiration time.Duration) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.GetEx(ctx, key, expiration))
}

func (c *client) GetDel(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.GetDel(ctx, key))
}

func (c *client) StrLen(ctx context.Context, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.StrLen(ctx, key))
}

func (c *client) Incr(ctx context.Context, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Incr(ctx, key))
}

func (c *client) Decr(ctx context.Context, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Decr(ctx, key))
}

func (c *client) IncrBy(ctx context.Context, key string, value int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.IncrBy(ctx, key, value))
}

func (c *client) DecrBy(ctx context.Context, key string, decrement int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.DecrBy(ctx, key, decrement))
}

func (c *client) IncrByFloat(ctx context.Context, key string, value float64) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.IncrByFloat(ctx, key, value))
}

func (c *client) Expire(ctx context.Context, key string, expiration time.Duration) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.Expire(ctx, key, expiration))
}

func (c *client) ExpireAt(ctx context.Context, key string, tm time.Time) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.ExpireAt(ctx, key, tm))
}

func (c *client) PExpire(ctx context.Context, key string, expiration time.Duration) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.PExpire(ctx, key, expiration))
}

func (c *client) PExpireAt(ctx context.Context, key string, tm time.Time) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.PExpireAt(ctx, key, tm))
}

func (c *client) TTL(ctx context.Context, key string) *cmd.DurationCmd {
	return cmd.NewDurationCMD(c.client.TTL(ctx, key))
}

func (c *client) PTTL(ctx context.Context, key string) *cmd.DurationCmd {
	return cmd.NewDurationCMD(c.client.PTTL(ctx, key))
}

func (c *client) Exists(ctx context.Context, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Exists(ctx, keys...))
}

func (c *client) Unlink(ctx context.Context, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Unlink(ctx, keys...))
}

func (c *client) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Migrate(ctx, host, port, key, db, timeout))
}

func (c *client) Move(ctx context.Context, key string, db int) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.Move(ctx, key, db))
}

func (c *client) ObjectRefCount(ctx context.Context, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ObjectRefCount(ctx, key))
}

func (c *client) ObjectEncoding(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ObjectEncoding(ctx, key))
}

func (c *client) ObjectIdleTime(ctx context.Context, key string) *cmd.DurationCmd {
	return cmd.NewDurationCMD(c.client.ObjectIdleTime(ctx, key))
}

func (c *client) RandomKey(ctx context.Context) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.RandomKey(ctx))
}

func (c *client) Rename(ctx context.Context, key string, newkey string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Rename(ctx, key, newkey))
}

func (c *client) RenameNX(ctx context.Context, key string, newkey string) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.RenameNX(ctx, key, newkey))
}

func (c *client) Type(ctx context.Context, key string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Type(ctx, key))
}

func (c *client) Append(ctx context.Context, key, value string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Append(ctx, key, value))
}

func (c *client) MGet(ctx context.Context, keys ...string) *cmd.SliceCmd {
	return cmd.NewSliceCMD(c.client.MGet(ctx, keys...))
}

func (c *client) MSet(ctx context.Context, values ...interface{}) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.MSet(ctx, values...))
}

func (c *client) MSetNX(ctx context.Context, values ...interface{}) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.MSetNX(ctx, values...))
}

func (c *client) GetBit(ctx context.Context, key string, offset int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.GetBit(ctx, key, offset))
}

func (c *client) SetBit(ctx context.Context, key string, offset int64, value int) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SetBit(ctx, key, offset, value))
}

func (c *client) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitCount(ctx, key, bitCount))
}

func (c *client) BitOpAnd(ctx context.Context, destKey string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitOpAnd(ctx, destKey, keys...))
}

func (c *client) BitOpOr(ctx context.Context, destKey string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitOpOr(ctx, destKey, keys...))
}

func (c *client) BitOpXor(ctx context.Context, destKey string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitOpXor(ctx, destKey, keys...))
}

func (c *client) BitOpNot(ctx context.Context, destKey string, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitOpNot(ctx, destKey, key))
}

func (c *client) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitPos(ctx, key, bit, pos...))
}

func (c *client) BitField(ctx context.Context, key string, args ...interface{}) *cmd.IntSliceCmd {
	return cmd.NewIntSliceCMD(c.client.BitField(ctx, key, args...))
}

func (c *client) SetArgs(ctx context.Context, key string, value interface{}, args redis.SetArgs) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.SetArgs(ctx, key, value, args))
}

func (c *client) Scan(ctx context.Context, cursor uint64, match string, count int64) *cmd.ScanCmd {
	return cmd.NewScanCMD(c.client.Scan(ctx, cursor, match, count))
}

func (c *client) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *cmd.ScanCmd {
	return cmd.NewScanCMD(c.client.ScanType(ctx, cursor, match, count, keyType))
}

func (c *client) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *cmd.ScanCmd {
	return cmd.NewScanCMD(c.client.SScan(ctx, key, cursor, match, count))
}

func (c *client) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *cmd.ScanCmd {
	return cmd.NewScanCMD(c.client.HScan(ctx, key, cursor, match, count))
}

func (c *client) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *cmd.ScanCmd {
	return cmd.NewScanCMD(c.client.ZScan(ctx, key, cursor, match, count))
}

func (c *client) HDel(ctx context.Context, key string, fields ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.HDel(ctx, key, fields...))
}

func (c *client) HExists(ctx context.Context, key, field string) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.HExists(ctx, key, field))
}

func (c *client) HGet(ctx context.Context, key, field string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.HGet(ctx, key, field))
}

func (c *client) HGetAll(ctx context.Context, key string) *cmd.StringStringMapCmd {
	return cmd.NewStringStringMapCMD(c.client.HGetAll(ctx, key))
}

func (c *client) HIncrBy(ctx context.Context, key, field string, incr int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.HIncrBy(ctx, key, field, incr))
}

func (c *client) HIncrByFloat(ctx context.Context, key, field string, incr float64) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.HIncrByFloat(ctx, key, field, incr))
}

func (c *client) HKeys(ctx context.Context, key string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.HKeys(ctx, key))
}

func (c *client) HLen(ctx context.Context, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.HLen(ctx, key))
}

func (c *client) HMGet(ctx context.Context, key string, fields ...string) *cmd.SliceCmd {
	return cmd.NewSliceCMD(c.client.HMGet(ctx, key, fields...))
}

func (c *client) HSet(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.HSet(ctx, key, values...))
}

func (c *client) HMSet(ctx context.Context, key string, values ...interface{}) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.HMSet(ctx, key, values...))
}

func (c *client) HSetNX(ctx context.Context, key, field string, value interface{}) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.HSetNX(ctx, key, field, value))
}

func (c *client) HVals(ctx context.Context, key string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.HVals(ctx, key))
}

func (c *client) HRandField(ctx context.Context, key string, count int) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.HRandField(ctx, key, count))
}

func (c *client) HRandFieldWithValues(ctx context.Context, key string, count int) *cmd.KeyValueSliceCmd {
	return cmd.NewKeyValueSliceCMD(c.client.HRandFieldWithValues(ctx, key, count))
}

func (c *client) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.BLPop(ctx, timeout, keys...))
}

func (c *client) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.BRPop(ctx, timeout, keys...))
}

func (c *client) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.BRPopLPush(ctx, source, destination, timeout))
}

func (c *client) LIndex(ctx context.Context, key string, index int64) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.LIndex(ctx, key, index))
}

func (c *client) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LInsert(ctx, key, op, pivot, value))
}

func (c *client) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LInsertBefore(ctx, key, pivot, value))
}

func (c *client) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LInsertAfter(ctx, key, pivot, value))
}

func (c *client) LLen(ctx context.Context, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LLen(ctx, key))
}

func (c *client) LPop(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.LPop(ctx, key))
}

func (c *client) LPopCount(ctx context.Context, key string, count int) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.LPopCount(ctx, key, count))
}

func (c *client) LPos(ctx context.Context, key string, value string, a redis.LPosArgs) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LPos(ctx, key, value, a))
}

func (c *client) LPosCount(ctx context.Context, key string, value string, count int64, a redis.LPosArgs) *cmd.IntSliceCmd {
	return cmd.NewIntSliceCMD(c.client.LPosCount(ctx, key, value, count, a))
}

func (c *client) LPush(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LPush(ctx, key, values...))
}

func (c *client) LPushX(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LPushX(ctx, key, values...))
}

func (c *client) LRange(ctx context.Context, key string, start, stop int64) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.LRange(ctx, key, start, stop))
}

func (c *client) LRem(ctx context.Context, key string, count int64, value interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LRem(ctx, key, count, value))
}

func (c *client) LSet(ctx context.Context, key string, index int64, value interface{}) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.LSet(ctx, key, index, value))
}

func (c *client) LTrim(ctx context.Context, key string, start int64, stop int64) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.LTrim(ctx, key, start, stop))
}

func (c *client) RPop(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.RPop(ctx, key))
}

func (c *client) RPopCount(ctx context.Context, key string, count int) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.RPopCount(ctx, key, count))
}

func (c *client) RPopLPush(ctx context.Context, source, destination string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.RPopLPush(ctx, source, destination))
}

func (c *client) RPush(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.RPush(ctx, key, values...))
}

func (c *client) RPushX(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.RPushX(ctx, key, values...))
}

func (c *client) LMove(ctx context.Context, source, destination, srcpos, destpos string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.LMove(ctx, source, destination, srcpos, destpos))
}

func (c *client) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.BLMove(ctx, source, destination, srcpos, destpos, timeout))
}

func (c *client) SAdd(ctx context.Context, key string, members ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SAdd(ctx, key, members...))
}

func (c *client) SCard(ctx context.Context, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SCard(ctx, key))
}

func (c *client) SDiff(ctx context.Context, keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SDiff(ctx, keys...))
}

func (c *client) SDiffStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SDiffStore(ctx, destination, keys...))
}

func (c *client) SInter(ctx context.Context, keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SInter(ctx, keys...))
}

func (c *client) SInterStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SInterStore(ctx, destination, keys...))
}

func (c *client) SIsMember(ctx context.Context, key string, member interface{}) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.SIsMember(ctx, key, member))
}

func (c *client) SMIsMember(ctx context.Context, key string, members ...interface{}) *cmd.BoolSliceCmd {
	return cmd.NewBoolSliceCMD(c.client.SMIsMember(ctx, key, members...))
}

func (c *client) SMembers(ctx context.Context, key string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SMembers(ctx, key))
}

func (c *client) SMembersMap(ctx context.Context, key string) *cmd.StringStructMapCmd {
	return cmd.NewStringStructMapCMD(c.client.SMembersMap(ctx, key))
}

func (c *client) SMove(ctx context.Context, source, destination string, member interface{}) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.SMove(ctx, source, destination, member))
}

func (c *client) SPop(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.SPop(ctx, key))
}

func (c *client) SPopN(ctx context.Context, key string, count int64) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SPopN(ctx, key, count))
}

func (c *client) SRandMember(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.SRandMember(ctx, key))
}

func (c *client) SRandMemberN(ctx context.Context, key string, count int64) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SRandMemberN(ctx, key, count))
}

func (c *client) SRem(ctx context.Context, key string, members ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SRem(ctx, key, members...))
}

func (c *client) SUnion(ctx context.Context, keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SUnion(ctx, keys...))
}

func (c *client) SUnionStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SUnionStore(ctx, destination, keys...))
}

func (c *client) ZAddArgs(ctx context.Context, key string, args redis.ZAddArgs) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZAddArgs(ctx, key, args))
}

func (c *client) ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.ZAddArgsIncr(ctx, key, args))
}

func (c *client) ZAdd(ctx context.Context, key string, members ...redis.Z) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZAdd(ctx, key, members...))
}

func (c *client) ZAddNX(ctx context.Context, key string, members ...redis.Z) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZAddNX(ctx, key, members...))
}

func (c *client) ZAddXX(ctx context.Context, key string, members ...redis.Z) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZAddXX(ctx, key, members...))
}

func (c *client) ZCard(ctx context.Context, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZCard(ctx, key))
}

func (c *client) ZCount(ctx context.Context, key, min, max string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZCount(ctx, key, min, max))
}

func (c *client) ZLexCount(ctx context.Context, key, min, max string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZLexCount(ctx, key, min, max))
}

func (c *client) ZIncrBy(ctx context.Context, key string, increment float64, member string) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.ZIncrBy(ctx, key, increment, member))
}

func (c *client) ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZInterStore(ctx, destination, store))
}

func (c *client) ZInter(ctx context.Context, store *redis.ZStore) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZInter(ctx, store))
}

func (c *client) ZInterWithScores(ctx context.Context, store *redis.ZStore) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZInterWithScores(ctx, store))
}

func (c *client) ZMScore(ctx context.Context, key string, members ...string) *cmd.FloatSliceCmd {
	return cmd.NewFloatSliceCMD(c.client.ZMScore(ctx, key, members...))
}

func (c *client) ZPopMax(ctx context.Context, key string, count ...int64) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZPopMax(ctx, key, count...))
}

func (c *client) ZPopMin(ctx context.Context, key string, count ...int64) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZPopMin(ctx, key, count...))
}

func (c *client) ZRangeArgs(ctx context.Context, z redis.ZRangeArgs) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRangeArgs(ctx, z))
}

func (c *client) ZRangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRangeArgsWithScores(ctx, z))
}

func (c *client) ZRange(ctx context.Context, key string, start, stop int64) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRange(ctx, key, start, stop))
}

func (c *client) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRangeWithScores(ctx, key, start, stop))
}

func (c *client) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRangeByScore(ctx, key, opt))
}

func (c *client) ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRangeByLex(ctx, key, opt))
}

func (c *client) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRangeByScoreWithScores(ctx, key, opt))
}

func (c *client) ZRangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRangeStore(ctx, dst, z))
}

func (c *client) ZRank(ctx context.Context, key, member string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRank(ctx, key, member))
}

func (c *client) ZRem(ctx context.Context, key string, members ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRem(ctx, key, members...))
}

func (c *client) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRemRangeByRank(ctx, key, start, stop))
}

func (c *client) ZRemRangeByScore(ctx context.Context, key, min, max string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRemRangeByScore(ctx, key, min, max))
}

func (c *client) ZRemRangeByLex(ctx context.Context, key, min, max string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRemRangeByLex(ctx, key, min, max))
}

func (c *client) ZRevRange(ctx context.Context, key string, start, stop int64) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRevRange(ctx, key, start, stop))
}

func (c *client) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRevRangeWithScores(ctx, key, start, stop))
}

func (c *client) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRevRangeByScore(ctx, key, opt))
}

func (c *client) ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRevRangeByLex(ctx, key, opt))
}

func (c *client) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRevRangeByScoreWithScores(ctx, key, opt))
}

func (c *client) ZRevRank(ctx context.Context, key, member string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRevRank(ctx, key, member))
}

func (c *client) ZScore(ctx context.Context, key, member string) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.ZScore(ctx, key, member))
}

func (c *client) ZUnion(ctx context.Context, store redis.ZStore) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZUnion(ctx, store))
}

func (c *client) ZUnionWithScores(ctx context.Context, store redis.ZStore) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZUnionWithScores(ctx, store))
}

func (c *client) ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZUnionStore(ctx, dest, store))
}

func (c *client) ZRandMember(ctx context.Context, key string, count int) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRandMember(ctx, key, count))
}

func (c *client) ZRandMemberWithScores(ctx context.Context, key string, count int) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRandMemberWithScores(ctx, key, count))
}

func (c *client) ZDiff(ctx context.Context, keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZDiff(ctx, keys...))
}

func (c *client) ZDiffWithScores(ctx context.Context, keys ...string) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZDiffWithScores(ctx, keys...))
}

func (c *client) ZDiffStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZDiffStore(ctx, destination, keys...))
}

func (c *client) PFAdd(ctx context.Context, key string, els ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.PFAdd(ctx, key, els...))
}

func (c *client) PFCount(ctx context.Context, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.PFCount(ctx, keys...))
}

func (c *client) PFMerge(ctx context.Context, dest string, keys ...string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.PFMerge(ctx, dest, keys...))
}

func (c *client) BgRewriteAOF(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.BgRewriteAOF(ctx))
}

func (c *client) BgSave(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.BgSave(ctx))
}

func (c *client) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.GeoAdd(ctx, key, geoLocation...))
}

func (c *client) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *cmd.GeoLocationCmd {
	return cmd.NewGeoLocationCMD(c.client.GeoRadius(ctx, key, longitude, latitude, query))
}

func (c *client) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.GeoRadiusStore(ctx, key, longitude, latitude, query))
}

func (c *client) GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *cmd.GeoLocationCmd {
	return cmd.NewGeoLocationCMD(c.client.GeoRadiusByMember(ctx, key, member, query))
}

func (c *client) GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.GeoRadiusByMemberStore(ctx, key, member, query))
}

func (c *client) GeoSearch(ctx context.Context, key string, q *redis.GeoSearchQuery) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.GeoSearch(ctx, key, q))
}

func (c *client) GeoSearchLocation(ctx context.Context, key string, q *redis.GeoSearchLocationQuery) *cmd.GeoSearchLocationCmd {
	return cmd.NewGeoSearchLocationCMD(c.client.GeoSearchLocation(ctx, key, q))
}

func (c *client) GeoSearchStore(ctx context.Context, key, store string, q *redis.GeoSearchStoreQuery) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.GeoSearchStore(ctx, key, store, q))
}

func (c *client) GeoDist(ctx context.Context, key string, member1, member2, unit string) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.GeoDist(ctx, key, member1, member2, unit))
}

func (c *client) GeoHash(ctx context.Context, key string, members ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.GeoHash(ctx, key, members...))
}

func (c *client) GeoPos(ctx context.Context, key string, members ...string) *cmd.GeoPosCmd {
	return cmd.NewGeoPosCMD(c.client.GeoPos(ctx, key, members...))
}

func (c *client) Watch(ctx context.Context, fn func(tx *redis.Tx) error, keys ...string) error {
	return c.client.Watch(ctx, fn, keys...)
}

func (c *client) Wait(ctx context.Context, numSlaves int, timeout time.Duration) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Wait(ctx, numSlaves, timeout))
}

func (c *client) Pipeline(ctx context.Context, fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error) {
	pipeliner := c.client.Pipeline()
	fn(pipeliner)
	return pipeliner.Exec(ctx)
}

func (c *client) TxPipeline(ctx context.Context, fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error) {
	txPipeline := c.client.TxPipeline()
	fn(txPipeline)
	return txPipeline.Exec(ctx)
}

func (c *client) ClientKill(ctx context.Context, ipPort string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClientKill(ctx, ipPort))
}

func (c *client) ClientKillByFilter(ctx context.Context, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClientKillByFilter(ctx, keys...))
}

func (c *client) ClientList(ctx context.Context) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ClientList(ctx))
}

func (c *client) ClientPause(ctx context.Context, dur time.Duration) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.ClientPause(ctx, dur))
}

func (c *client) ClientID(ctx context.Context) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClientID(ctx))
}

func (c *client) ClientUnblock(ctx context.Context, id int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClientUnblock(ctx, id))
}

func (c *client) ClientUnblockWithError(ctx context.Context, id int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClientUnblockWithError(ctx, id))
}

func (c *client) ConfigGet(ctx context.Context, parameter string) *cmd.StringStringMapCmd {
	return cmd.NewStringStringMapCMD(c.client.ConfigGet(ctx, parameter))
}

func (c *client) ConfigResetStat(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ConfigResetStat(ctx))
}

func (c *client) ConfigSet(ctx context.Context, parameter, value string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ConfigSet(ctx, parameter, value))
}

func (c *client) ConfigRewrite(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ConfigRewrite(ctx))
}

func (c *client) DBSize(ctx context.Context) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.DBSize(ctx))
}

func (c *client) FlushAll(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.FlushAll(ctx))
}

func (c *client) FlushAllAsync(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.FlushAllAsync(ctx))
}

func (c *client) FlushDB(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.FlushDB(ctx))
}

func (c *client) FlushDBAsync(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.FlushDBAsync(ctx))
}

func (c *client) Info(ctx context.Context, section ...string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.Info(ctx, section...))
}

func (c *client) LastSave(ctx context.Context) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LastSave(ctx))
}

func (c *client) Save(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Save(ctx))
}

func (c *client) Shutdown(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Shutdown(ctx))
}

func (c *client) ShutdownSave(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ShutdownSave(ctx))
}

func (c *client) ShutdownNoSave(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ShutdownNoSave(ctx))
}

func (c *client) SlaveOf(ctx context.Context, host, port string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.SlaveOf(ctx, host, port))
}

func (c *client) SlowLogGet(ctx context.Context, num int64) *cmd.SlowLogCmd {
	return cmd.NewSlowLogCMD(c.client.SlowLogGet(ctx, num))
}

func (c *client) Time(ctx context.Context) *cmd.TimeCmd {
	return cmd.NewTimeCMD(c.client.Time(ctx))
}

func (c *client) DebugObject(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.DebugObject(ctx, key))
}

func (c *client) ReadOnly(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ReadOnly(ctx))
}

func (c *client) ReadWrite(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ReadWrite(ctx))
}

func (c *client) MemoryUsage(ctx context.Context, key string, samples ...int) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.MemoryUsage(ctx, key, samples...))
}

func (c *client) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *cmd.Cmd {
	return cmd.NewCMD(c.client.Eval(ctx, script, keys, args...))
}

func (c *client) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *cmd.Cmd {
	return cmd.NewCMD(c.client.EvalSha(ctx, sha1, keys, args...))
}

func (c *client) ScriptExists(ctx context.Context, hashes ...string) *cmd.BoolSliceCmd {
	return cmd.NewBoolSliceCMD(c.client.ScriptExists(ctx, hashes...))
}

func (c *client) ScriptFlush(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ScriptFlush(ctx))
}

func (c *client) ScriptKill(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ScriptKill(ctx))
}

func (c *client) ScriptLoad(ctx context.Context, script string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ScriptLoad(ctx, script))
}

func (c *client) Publish(ctx context.Context, channel string, message interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Publish(ctx, channel, message))
}

func (c *client) PubSubChannels(ctx context.Context, pattern string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.PubSubChannels(ctx, pattern))
}

func (c *client) PubSubNumSub(ctx context.Context, channels ...string) *cmd.StringIntMapCmd {
	return cmd.NewStringIntMapCMD(c.client.PubSubNumSub(ctx, channels...))
}

func (c *client) PubSubNumPat(ctx context.Context) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.PubSubNumPat(ctx))
}

func (c *client) ClusterSlots(ctx context.Context) *cmd.ClusterSlotsCmd {
	return cmd.NewClusterSlotsCMD(c.client.ClusterSlots(ctx))
}

func (c *client) ClusterNodes(ctx context.Context) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ClusterNodes(ctx))
}

func (c *client) ClusterMeet(ctx context.Context, host, port string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterMeet(ctx, host, port))
}

func (c *client) ClusterForget(ctx context.Context, nodeID string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterForget(ctx, nodeID))
}

func (c *client) ClusterReplicate(ctx context.Context, nodeID string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterReplicate(ctx, nodeID))
}

func (c *client) ClusterResetSoft(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterResetSoft(ctx))
}

func (c *client) ClusterResetHard(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterResetHard(ctx))
}

func (c *client) ClusterInfo(ctx context.Context) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ClusterInfo(ctx))
}

func (c *client) ClusterKeySlot(ctx context.Context, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClusterKeySlot(ctx, key))
}

func (c *client) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ClusterGetKeysInSlot(ctx, slot, count))
}

func (c *client) ClusterCountFailureReports(ctx context.Context, nodeID string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClusterCountFailureReports(ctx, nodeID))
}

func (c *client) ClusterCountKeysInSlot(ctx context.Context, slot int) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClusterCountKeysInSlot(ctx, slot))
}

func (c *client) ClusterDelSlots(ctx context.Context, slots ...int) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterDelSlots(ctx, slots...))
}

func (c *client) ClusterDelSlotsRange(ctx context.Context, min, max int) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterDelSlotsRange(ctx, min, max))
}

func (c *client) ClusterSaveConfig(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterSaveConfig(ctx))
}

func (c *client) ClusterSlaves(ctx context.Context, nodeID string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ClusterSlaves(ctx, nodeID))
}

func (c *client) ClusterFailover(ctx context.Context) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterFailover(ctx))
}

func (c *client) ClusterAddSlots(ctx context.Context, slots ...int) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterAddSlots(ctx, slots...))
}

func (c *client) ClusterAddSlotsRange(ctx context.Context, min, max int) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterAddSlotsRange(ctx, min, max))
}

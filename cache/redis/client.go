package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/raylin666/go-utils/cache/redis/cmd"
	"time"
)

var _ Client = (*client)(nil)

type Client interface {
	Ping() *cmd.StatusCmd
	Command() *cmd.CommandsInfoCmd
	ClientGetName() *cmd.StringCmd
	Echo(message interface{}) *cmd.StringCmd
	Keys(pattern string) *cmd.StringSliceCmd
	Dump(key string) *cmd.StringCmd
	Get(key string) *cmd.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *cmd.StatusCmd
	SetEX(key string, value interface{}, expiration time.Duration) *cmd.StatusCmd
	SetNX(key string, value interface{}, expiration time.Duration) *cmd.BoolCmd
	SetXX(key string, value interface{}, expiration time.Duration) *cmd.BoolCmd
	Del(keys ...string) *cmd.IntCmd
	GetRange(key string, start, end int64) *cmd.StringCmd
	GetSet(key string, value interface{}) *cmd.StringCmd
	GetEx(key string, expiration time.Duration) *cmd.StringCmd
	GetDel(key string) *cmd.StringCmd
	StrLen(key string) *cmd.IntCmd
	Incr(key string) *cmd.IntCmd
	Decr(key string) *cmd.IntCmd
	IncrBy(key string, value int64) *cmd.IntCmd
	DecrBy(key string, decrement int64) *cmd.IntCmd
	IncrByFloat(key string, value float64) *cmd.FloatCmd
	Expire(key string, expiration time.Duration) *cmd.BoolCmd
	ExpireAt(key string, tm time.Time) *cmd.BoolCmd
	PExpire(key string, expiration time.Duration) *cmd.BoolCmd
	PExpireAt(key string, tm time.Time) *cmd.BoolCmd
	TTL(key string) *cmd.DurationCmd
	PTTL(key string) *cmd.DurationCmd
	Exists(keys ...string) *cmd.IntCmd
	Unlink(keys ...string) *cmd.IntCmd
	Migrate(host, port, key string, db int, timeout time.Duration) *cmd.StatusCmd
	Move(key string, db int) *cmd.BoolCmd
	ObjectRefCount(key string) *cmd.IntCmd
	ObjectEncoding(key string) *cmd.StringCmd
	ObjectIdleTime(key string) *cmd.DurationCmd
	RandomKey() *cmd.StringCmd
	Rename(key string, newkey string) *cmd.StatusCmd
	RenameNX(key string, newkey string) *cmd.BoolCmd
	Type(key string) *cmd.StatusCmd
	Append(key, value string) *cmd.IntCmd
	MGet(keys ...string) *cmd.SliceCmd
	MSet(values ...interface{}) *cmd.StatusCmd
	MSetNX(values ...interface{}) *cmd.BoolCmd
	GetBit(key string, offset int64) *cmd.IntCmd
	SetBit(key string, offset int64, value int) *cmd.IntCmd
	BitCount(key string, bitCount *redis.BitCount) *cmd.IntCmd
	BitOpAnd(destKey string, keys ...string) *cmd.IntCmd
	BitOpOr(destKey string, keys ...string) *cmd.IntCmd
	BitOpXor(destKey string, keys ...string) *cmd.IntCmd
	BitOpNot(destKey string, key string) *cmd.IntCmd
	BitPos(key string, bit int64, pos ...int64) *cmd.IntCmd
	BitField(key string, args ...interface{}) *cmd.IntSliceCmd
	SetArgs(key string, value interface{}, args redis.SetArgs) *cmd.StatusCmd
	Scan(cursor uint64, match string, count int64) *cmd.ScanCmd
	ScanType(cursor uint64, match string, count int64, keyType string) *cmd.ScanCmd
	SScan(key string, cursor uint64, match string, count int64) *cmd.ScanCmd
	HScan(key string, cursor uint64, match string, count int64) *cmd.ScanCmd
	ZScan(key string, cursor uint64, match string, count int64) *cmd.ScanCmd
	HDel(key string, fields ...string) *cmd.IntCmd
	HExists(key, field string) *cmd.BoolCmd
	HGet(key, field string) *cmd.StringCmd
	HGetAll(key string) *cmd.StringStringMapCmd
	HIncrBy(key, field string, incr int64) *cmd.IntCmd
	HIncrByFloat(key, field string, incr float64) *cmd.FloatCmd
	HKeys(key string) *cmd.StringSliceCmd
	HLen(key string) *cmd.IntCmd
	HMGet(key string, fields ...string) *cmd.SliceCmd
	HSet(key string, values ...interface{}) *cmd.IntCmd
	HMSet(key string, values ...interface{}) *cmd.BoolCmd
	HSetNX(key, field string, value interface{}) *cmd.BoolCmd
	HVals(key string) *cmd.StringSliceCmd
	HRandField(key string, count int, withValues bool) *cmd.StringSliceCmd
	BLPop(timeout time.Duration, keys ...string) *cmd.StringSliceCmd
	BRPop(timeout time.Duration, keys ...string) *cmd.StringSliceCmd
	BRPopLPush(source, destination string, timeout time.Duration) *cmd.StringCmd
	LIndex(key string, index int64) *cmd.StringCmd
	LInsert(key, op string, pivot, value interface{}) *cmd.IntCmd
	LInsertBefore(key string, pivot, value interface{}) *cmd.IntCmd
	LInsertAfter(key string, pivot, value interface{}) *cmd.IntCmd
	LLen(key string) *cmd.IntCmd
	LPop(key string) *cmd.StringCmd
	LPopCount(key string, count int) *cmd.StringSliceCmd
	LPos(key string, value string, a redis.LPosArgs) *cmd.IntCmd
	LPosCount(key string, value string, count int64, a redis.LPosArgs) *cmd.IntSliceCmd
	LPush(key string, values ...interface{}) *cmd.IntCmd
	LPushX(key string, values ...interface{}) *cmd.IntCmd
	LRange(key string, start, stop int64) *cmd.StringSliceCmd
	LRem(key string, count int64, value interface{}) *cmd.IntCmd
	LSet(key string, index int64, value interface{}) *cmd.StatusCmd
	LTrim(key string, start int64, stop int64) *cmd.StatusCmd
	RPop(key string) *cmd.StringCmd
	RPopCount(key string, count int) *cmd.StringSliceCmd
	RPopLPush(source, destination string) *cmd.StringCmd
	RPush(key string, values ...interface{}) *cmd.IntCmd
	RPushX(key string, values ...interface{}) *cmd.IntCmd
	LMove(source, destination, srcpos, destpos string) *cmd.StringCmd
	BLMove(source, destination, srcpos, destpos string, timeout time.Duration) *cmd.StringCmd
	SAdd(key string, members ...interface{}) *cmd.IntCmd
	SCard(key string) *cmd.IntCmd
	SDiff(keys ...string) *cmd.StringSliceCmd
	SDiffStore(destination string, keys ...string) *cmd.IntCmd
	SInter(keys ...string) *cmd.StringSliceCmd
	SInterStore(destination string, keys ...string) *cmd.IntCmd
	SIsMember(key string, member interface{}) *cmd.BoolCmd
	SMIsMember(key string, members ...interface{}) *cmd.BoolSliceCmd
	SMembers(key string) *cmd.StringSliceCmd
	SMembersMap(key string) *cmd.StringStructMapCmd
	SMove(source, destination string, member interface{}) *cmd.BoolCmd
	SPop(key string) *cmd.StringCmd
	SPopN(key string, count int64) *cmd.StringSliceCmd
	SRandMember(key string) *cmd.StringCmd
	SRandMemberN(key string, count int64) *cmd.StringSliceCmd
	SRem(key string, members ...interface{}) *cmd.IntCmd
	SUnion(keys ...string) *cmd.StringSliceCmd
	SUnionStore(destination string, keys ...string) *cmd.IntCmd
	ZAddArgs(key string, args redis.ZAddArgs) *cmd.IntCmd
	ZAddArgsIncr(key string, args redis.ZAddArgs) *cmd.FloatCmd
	ZAdd(key string, members ...*redis.Z) *cmd.IntCmd
	ZAddNX(key string, members ...*redis.Z) *cmd.IntCmd
	ZAddXX(key string, members ...*redis.Z) *cmd.IntCmd
	ZCard(key string) *cmd.IntCmd
	ZCount(key, min, max string) *cmd.IntCmd
	ZLexCount(key, min, max string) *cmd.IntCmd
	ZIncrBy(key string, increment float64, member string) *cmd.FloatCmd
	ZInterStore(destination string, store *redis.ZStore) *cmd.IntCmd
	ZInter(store *redis.ZStore) *cmd.StringSliceCmd
	ZInterWithScores(store *redis.ZStore) *cmd.ZSliceCmd
	ZMScore(key string, members ...string) *cmd.FloatSliceCmd
	ZPopMax(key string, count ...int64) *cmd.ZSliceCmd
	ZPopMin(key string, count ...int64) *cmd.ZSliceCmd
	ZRangeArgs(z redis.ZRangeArgs) *cmd.StringSliceCmd
	ZRangeArgsWithScores(z redis.ZRangeArgs) *cmd.ZSliceCmd
	ZRange(key string, start, stop int64) *cmd.StringSliceCmd
	ZRangeWithScores(key string, start, stop int64) *cmd.ZSliceCmd
	ZRangeByScore(key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd
	ZRangeByLex(key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd
	ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *cmd.ZSliceCmd
	ZRangeStore(dst string, z redis.ZRangeArgs) *cmd.IntCmd
	ZRank(key, member string) *cmd.IntCmd
	ZRem(key string, members ...interface{}) *cmd.IntCmd
	ZRemRangeByRank(key string, start, stop int64) *cmd.IntCmd
	ZRemRangeByScore(key, min, max string) *cmd.IntCmd
	ZRemRangeByLex(key, min, max string) *cmd.IntCmd
	ZRevRange(key string, start, stop int64) *cmd.StringSliceCmd
	ZRevRangeWithScores(key string, start, stop int64) *cmd.ZSliceCmd
	ZRevRangeByScore(key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd
	ZRevRangeByLex(key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd
	ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *cmd.ZSliceCmd
	ZRevRank(key, member string) *cmd.IntCmd
	ZScore(key, member string) *cmd.FloatCmd
	ZUnion(store redis.ZStore) *cmd.StringSliceCmd
	ZUnionWithScores(store redis.ZStore) *cmd.ZSliceCmd
	ZUnionStore(dest string, store *redis.ZStore) *cmd.IntCmd
	ZRandMember(key string, count int, withScores bool) *cmd.StringSliceCmd
	ZDiff(keys ...string) *cmd.StringSliceCmd
	ZDiffWithScores(keys ...string) *cmd.ZSliceCmd
	ZDiffStore(destination string, keys ...string) *cmd.IntCmd
	PFAdd(key string, els ...interface{}) *cmd.IntCmd
	PFCount(keys ...string) *cmd.IntCmd
	PFMerge(dest string, keys ...string) *cmd.StatusCmd
	BgRewriteAOF() *cmd.StatusCmd
	BgSave() *cmd.StatusCmd
	GeoAdd(key string, geoLocation ...*redis.GeoLocation) *cmd.IntCmd
	GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *cmd.GeoLocationCmd
	GeoRadiusStore(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *cmd.IntCmd
	GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *cmd.GeoLocationCmd
	GeoRadiusByMemberStore(key, member string, query *redis.GeoRadiusQuery) *cmd.IntCmd
	GeoSearch(key string, q *redis.GeoSearchQuery) *cmd.StringSliceCmd
	GeoSearchLocation(key string, q *redis.GeoSearchLocationQuery) *cmd.GeoSearchLocationCmd
	GeoSearchStore(key, store string, q *redis.GeoSearchStoreQuery) *cmd.IntCmd
	GeoDist(key string, member1, member2, unit string) *cmd.FloatCmd
	GeoHash(key string, members ...string) *cmd.StringSliceCmd
	GeoPos(key string, members ...string) *cmd.GeoPosCmd
	Watch(fn func(tx *redis.Tx) error, keys ...string) error
	Wait(numSlaves int, timeout time.Duration) *cmd.IntCmd
	Pipeline(fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error)
	TxPipeline(fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error)
	ClientKill(ipPort string) *cmd.StatusCmd
	ClientKillByFilter(keys ...string) *cmd.IntCmd
	ClientList() *cmd.StringCmd
	ClientPause(dur time.Duration) *cmd.BoolCmd
	ClientID() *cmd.IntCmd
	ClientUnblock(id int64) *cmd.IntCmd
	ClientUnblockWithError(id int64) *cmd.IntCmd
	ConfigGet(parameter string) *cmd.SliceCmd
	ConfigResetStat() *cmd.StatusCmd
	ConfigSet(parameter, value string) *cmd.StatusCmd
	ConfigRewrite() *cmd.StatusCmd
	DBSize() *cmd.IntCmd
	FlushAll() *cmd.StatusCmd
	FlushAllAsync() *cmd.StatusCmd
	FlushDB() *cmd.StatusCmd
	FlushDBAsync() *cmd.StatusCmd
	Info(section ...string) *cmd.StringCmd
	LastSave() *cmd.IntCmd
	Save() *cmd.StatusCmd
	Shutdown() *cmd.StatusCmd
	ShutdownSave() *cmd.StatusCmd
	ShutdownNoSave() *cmd.StatusCmd
	SlaveOf(host, port string) *cmd.StatusCmd
	SlowLogGet(num int64) *cmd.SlowLogCmd
	Time() *cmd.TimeCmd
	DebugObject(key string) *cmd.StringCmd
	ReadOnly() *cmd.StatusCmd
	ReadWrite() *cmd.StatusCmd
	MemoryUsage(key string, samples ...int) *cmd.IntCmd
	Eval(script string, keys []string, args ...interface{}) *cmd.Cmd
	EvalSha(sha1 string, keys []string, args ...interface{}) *cmd.Cmd
	ScriptExists(hashes ...string) *cmd.BoolSliceCmd
	ScriptFlush() *cmd.StatusCmd
	ScriptKill() *cmd.StatusCmd
	ScriptLoad(script string) *cmd.StringCmd
	Publish(channel string, message interface{}) *cmd.IntCmd
	PubSubChannels(pattern string) *cmd.StringSliceCmd
	PubSubNumSub(channels ...string) *cmd.StringIntMapCmd
	PubSubNumPat() *cmd.IntCmd
	ClusterSlots() *cmd.ClusterSlotsCmd
	ClusterNodes() *cmd.StringCmd
	ClusterMeet(host, port string) *cmd.StatusCmd
	ClusterForget(nodeID string) *cmd.StatusCmd
	ClusterReplicate(nodeID string) *cmd.StatusCmd
	ClusterResetSoft() *cmd.StatusCmd
	ClusterResetHard() *cmd.StatusCmd
	ClusterInfo() *cmd.StringCmd
	ClusterKeySlot(key string) *cmd.IntCmd
	ClusterGetKeysInSlot(slot int, count int) *cmd.StringSliceCmd
	ClusterCountFailureReports(nodeID string) *cmd.IntCmd
	ClusterCountKeysInSlot(slot int) *cmd.IntCmd
	ClusterDelSlots(slots ...int) *cmd.StatusCmd
	ClusterDelSlotsRange(min, max int) *cmd.StatusCmd
	ClusterSaveConfig() *cmd.StatusCmd
	ClusterSlaves(nodeID string) *cmd.StringSliceCmd
	ClusterFailover() *cmd.StatusCmd
	ClusterAddSlots(slots ...int) *cmd.StatusCmd
	ClusterAddSlotsRange(min, max int) *cmd.StatusCmd
}

type Options struct {
	redis.Options
}

type client struct {
	client  *redis.Client
	options *Options
	context context.Context
}

func NewClient(ctx context.Context, opt *Options) (Client, error) {
	var c = new(client)
	c.context = ctx
	c.options = opt
	c.client = redis.NewClient(&opt.Options)
	if err := c.Ping().Err(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *client) Ping() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Ping(c.context))
}

func (c *client) Command() *cmd.CommandsInfoCmd {
	return cmd.NewCommandsInfoCMD(c.client.Command(c.context))
}

func (c *client) ClientGetName() *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ClientGetName(c.context))
}

func (c *client) Echo(message interface{}) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.Echo(c.context, message))
}

func (c *client) Keys(pattern string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.Keys(c.context, pattern))
}

func (c *client) Dump(key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.Dump(c.context, key))
}

func (c *client) Get(key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.Get(c.context, key))
}

func (c *client) Set(key string, value interface{}, expiration time.Duration) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Set(c.context, key, value, expiration))
}

func (c *client) SetEX(key string, value interface{}, expiration time.Duration) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.SetEX(c.context, key, value, expiration))
}

func (c *client) SetNX(key string, value interface{}, expiration time.Duration) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.SetNX(c.context, key, value, expiration))
}

func (c *client) SetXX(key string, value interface{}, expiration time.Duration) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.SetXX(c.context, key, value, expiration))
}

func (c *client) Del(keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Del(c.context, keys...))
}

func (c *client) GetRange(key string, start, end int64) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.GetRange(c.context, key, start, end))
}

func (c *client) GetSet(key string, value interface{}) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.GetSet(c.context, key, value))
}

func (c *client) GetEx(key string, expiration time.Duration) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.GetEx(c.context, key, expiration))
}

func (c *client) GetDel(key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.GetDel(c.context, key))
}

func (c *client) StrLen(key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.StrLen(c.context, key))
}

func (c *client) Incr(key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Incr(c.context, key))
}

func (c *client) Decr(key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Decr(c.context, key))
}

func (c *client) IncrBy(key string, value int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.IncrBy(c.context, key, value))
}

func (c *client) DecrBy(key string, decrement int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.DecrBy(c.context, key, decrement))
}

func (c *client) IncrByFloat(key string, value float64) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.IncrByFloat(c.context, key, value))
}

func (c *client) Expire(key string, expiration time.Duration) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.Expire(c.context, key, expiration))
}

func (c *client) ExpireAt(key string, tm time.Time) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.ExpireAt(c.context, key, tm))
}

func (c *client) PExpire(key string, expiration time.Duration) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.PExpire(c.context, key, expiration))
}

func (c *client) PExpireAt(key string, tm time.Time) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.PExpireAt(c.context, key, tm))
}

func (c *client) TTL(key string) *cmd.DurationCmd {
	return cmd.NewDurationCMD(c.client.TTL(c.context, key))
}

func (c *client) PTTL(key string) *cmd.DurationCmd {
	return cmd.NewDurationCMD(c.client.PTTL(c.context, key))
}

func (c *client) Exists(keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Exists(c.context, keys...))
}

func (c *client) Unlink(keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Unlink(c.context, keys...))
}

func (c *client) Migrate(host, port, key string, db int, timeout time.Duration) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Migrate(c.context, host, port, key, db, timeout))
}

func (c *client) Move(key string, db int) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.Move(c.context, key, db))
}

func (c *client) ObjectRefCount(key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ObjectRefCount(c.context, key))
}

func (c *client) ObjectEncoding(key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ObjectEncoding(c.context, key))
}

func (c *client) ObjectIdleTime(key string) *cmd.DurationCmd {
	return cmd.NewDurationCMD(c.client.ObjectIdleTime(c.context, key))
}

func (c *client) RandomKey() *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.RandomKey(c.context))
}

func (c *client) Rename(key string, newkey string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Rename(c.context, key, newkey))
}

func (c *client) RenameNX(key string, newkey string) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.RenameNX(c.context, key, newkey))
}

func (c *client) Type(key string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Type(c.context, key))
}

func (c *client) Append(key, value string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Append(c.context, key, value))
}

func (c *client) MGet(keys ...string) *cmd.SliceCmd {
	return cmd.NewSliceCMD(c.client.MGet(c.context, keys...))
}

func (c *client) MSet(values ...interface{}) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.MSet(c.context, values...))
}

func (c *client) MSetNX(values ...interface{}) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.MSetNX(c.context, values...))
}

func (c *client) GetBit(key string, offset int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.GetBit(c.context, key, offset))
}

func (c *client) SetBit(key string, offset int64, value int) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SetBit(c.context, key, offset, value))
}

func (c *client) BitCount(key string, bitCount *redis.BitCount) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitCount(c.context, key, bitCount))
}

func (c *client) BitOpAnd(destKey string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitOpAnd(c.context, destKey, keys...))
}

func (c *client) BitOpOr(destKey string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitOpOr(c.context, destKey, keys...))
}

func (c *client) BitOpXor(destKey string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitOpXor(c.context, destKey, keys...))
}

func (c *client) BitOpNot(destKey string, key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitOpNot(c.context, destKey, key))
}

func (c *client) BitPos(key string, bit int64, pos ...int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.BitPos(c.context, key, bit, pos...))
}

func (c *client) BitField(key string, args ...interface{}) *cmd.IntSliceCmd {
	return cmd.NewIntSliceCMD(c.client.BitField(c.context, key, args...))
}

func (c *client) SetArgs(key string, value interface{}, args redis.SetArgs) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.SetArgs(c.context, key, value, args))
}

func (c *client) Scan(cursor uint64, match string, count int64) *cmd.ScanCmd {
	return cmd.NewScanCMD(c.client.Scan(c.context, cursor, match, count))
}

func (c *client) ScanType(cursor uint64, match string, count int64, keyType string) *cmd.ScanCmd {
	return cmd.NewScanCMD(c.client.ScanType(c.context, cursor, match, count, keyType))
}

func (c *client) SScan(key string, cursor uint64, match string, count int64) *cmd.ScanCmd {
	return cmd.NewScanCMD(c.client.SScan(c.context, key, cursor, match, count))
}

func (c *client) HScan(key string, cursor uint64, match string, count int64) *cmd.ScanCmd {
	return cmd.NewScanCMD(c.client.HScan(c.context, key, cursor, match, count))
}

func (c *client) ZScan(key string, cursor uint64, match string, count int64) *cmd.ScanCmd {
	return cmd.NewScanCMD(c.client.ZScan(c.context, key, cursor, match, count))
}

func (c *client) HDel(key string, fields ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.HDel(c.context, key, fields...))
}

func (c *client) HExists(key, field string) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.HExists(c.context, key, field))
}

func (c *client) HGet(key, field string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.HGet(c.context, key, field))
}

func (c *client) HGetAll(key string) *cmd.StringStringMapCmd {
	return cmd.NewStringStringMapCMD(c.client.HGetAll(c.context, key))
}

func (c *client) HIncrBy(key, field string, incr int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.HIncrBy(c.context, key, field, incr))
}

func (c *client) HIncrByFloat(key, field string, incr float64) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.HIncrByFloat(c.context, key, field, incr))
}

func (c *client) HKeys(key string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.HKeys(c.context, key))
}

func (c *client) HLen(key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.HLen(c.context, key))
}

func (c *client) HMGet(key string, fields ...string) *cmd.SliceCmd {
	return cmd.NewSliceCMD(c.client.HMGet(c.context, key, fields...))
}

func (c *client) HSet(key string, values ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.HSet(c.context, key, values...))
}

func (c *client) HMSet(key string, values ...interface{}) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.HMSet(c.context, key, values...))
}

func (c *client) HSetNX(key, field string, value interface{}) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.HSetNX(c.context, key, field, value))
}

func (c *client) HVals(key string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.HVals(c.context, key))
}

func (c *client) HRandField(key string, count int, withValues bool) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.HRandField(c.context, key, count, withValues))
}

func (c *client) BLPop(timeout time.Duration, keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.BLPop(c.context, timeout, keys...))
}

func (c *client) BRPop(timeout time.Duration, keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.BRPop(c.context, timeout, keys...))
}

func (c *client) BRPopLPush(source, destination string, timeout time.Duration) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.BRPopLPush(c.context, source, destination, timeout))
}

func (c *client) LIndex(key string, index int64) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.LIndex(c.context, key, index))
}

func (c *client) LInsert(key, op string, pivot, value interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LInsert(c.context, key, op, pivot, value))
}

func (c *client) LInsertBefore(key string, pivot, value interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LInsertBefore(c.context, key, pivot, value))
}

func (c *client) LInsertAfter(key string, pivot, value interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LInsertAfter(c.context, key, pivot, value))
}

func (c *client) LLen(key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LLen(c.context, key))
}

func (c *client) LPop(key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.LPop(c.context, key))
}

func (c *client) LPopCount(key string, count int) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.LPopCount(c.context, key, count))
}

func (c *client) LPos(key string, value string, a redis.LPosArgs) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LPos(c.context, key, value, a))
}

func (c *client) LPosCount(key string, value string, count int64, a redis.LPosArgs) *cmd.IntSliceCmd {
	return cmd.NewIntSliceCMD(c.client.LPosCount(c.context, key, value, count, a))
}

func (c *client) LPush(key string, values ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LPush(c.context, key, values...))
}

func (c *client) LPushX(key string, values ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LPushX(c.context, key, values...))
}

func (c *client) LRange(key string, start, stop int64) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.LRange(c.context, key, start, stop))
}

func (c *client) LRem(key string, count int64, value interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LRem(c.context, key, count, value))
}

func (c *client) LSet(key string, index int64, value interface{}) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.LSet(c.context, key, index, value))
}

func (c *client) LTrim(key string, start int64, stop int64) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.LTrim(c.context, key, start, stop))
}

func (c *client) RPop(key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.RPop(c.context, key))
}

func (c *client) RPopCount(key string, count int) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.RPopCount(c.context, key, count))
}

func (c *client) RPopLPush(source, destination string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.RPopLPush(c.context, source, destination))
}

func (c *client) RPush(key string, values ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.RPush(c.context, key, values...))
}

func (c *client) RPushX(key string, values ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.RPushX(c.context, key, values...))
}

func (c *client) LMove(source, destination, srcpos, destpos string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.LMove(c.context, source, destination, srcpos, destpos))
}

func (c *client) BLMove(source, destination, srcpos, destpos string, timeout time.Duration) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.BLMove(c.context, source, destination, srcpos, destpos, timeout))
}

func (c *client) SAdd(key string, members ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SAdd(c.context, key, members...))
}

func (c *client) SCard(key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SCard(c.context, key))
}

func (c *client) SDiff(keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SDiff(c.context, keys...))
}

func (c *client) SDiffStore(destination string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SDiffStore(c.context, destination, keys...))
}

func (c *client) SInter(keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SInter(c.context, keys...))
}

func (c *client) SInterStore(destination string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SInterStore(c.context, destination, keys...))
}

func (c *client) SIsMember(key string, member interface{}) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.SIsMember(c.context, key, member))
}

func (c *client) SMIsMember(key string, members ...interface{}) *cmd.BoolSliceCmd {
	return cmd.NewBoolSliceCMD(c.client.SMIsMember(c.context, key, members...))
}

func (c *client) SMembers(key string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SMembers(c.context, key))
}

func (c *client) SMembersMap(key string) *cmd.StringStructMapCmd {
	return cmd.NewStringStructMapCMD(c.client.SMembersMap(c.context, key))
}

func (c *client) SMove(source, destination string, member interface{}) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.SMove(c.context, source, destination, member))
}

func (c *client) SPop(key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.SPop(c.context, key))
}

func (c *client) SPopN(key string, count int64) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SPopN(c.context, key, count))
}

func (c *client) SRandMember(key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.SRandMember(c.context, key))
}

func (c *client) SRandMemberN(key string, count int64) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SRandMemberN(c.context, key, count))
}

func (c *client) SRem(key string, members ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SRem(c.context, key, members...))
}

func (c *client) SUnion(keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.SUnion(c.context, keys...))
}

func (c *client) SUnionStore(destination string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.SUnionStore(c.context, destination, keys...))
}

func (c *client) ZAddArgs(key string, args redis.ZAddArgs) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZAddArgs(c.context, key, args))
}

func (c *client) ZAddArgsIncr(key string, args redis.ZAddArgs) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.ZAddArgsIncr(c.context, key, args))
}

func (c *client) ZAdd(key string, members ...*redis.Z) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZAdd(c.context, key, members...))
}

func (c *client) ZAddNX(key string, members ...*redis.Z) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZAddNX(c.context, key, members...))
}

func (c *client) ZAddXX(key string, members ...*redis.Z) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZAddXX(c.context, key, members...))
}

func (c *client) ZCard(key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZCard(c.context, key))
}

func (c *client) ZCount(key, min, max string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZCount(c.context, key, min, max))
}

func (c *client) ZLexCount(key, min, max string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZLexCount(c.context, key, min, max))
}

func (c *client) ZIncrBy(key string, increment float64, member string) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.ZIncrBy(c.context, key, increment, member))
}

func (c *client) ZInterStore(destination string, store *redis.ZStore) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZInterStore(c.context, destination, store))
}

func (c *client) ZInter(store *redis.ZStore) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZInter(c.context, store))
}

func (c *client) ZInterWithScores(store *redis.ZStore) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZInterWithScores(c.context, store))
}

func (c *client) ZMScore(key string, members ...string) *cmd.FloatSliceCmd {
	return cmd.NewFloatSliceCMD(c.client.ZMScore(c.context, key, members...))
}

func (c *client) ZPopMax(key string, count ...int64) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZPopMax(c.context, key, count...))
}

func (c *client) ZPopMin(key string, count ...int64) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZPopMin(c.context, key, count...))
}

func (c *client) ZRangeArgs(z redis.ZRangeArgs) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRangeArgs(c.context, z))
}

func (c *client) ZRangeArgsWithScores(z redis.ZRangeArgs) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRangeArgsWithScores(c.context, z))
}

func (c *client) ZRange(key string, start, stop int64) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRange(c.context, key, start, stop))
}

func (c *client) ZRangeWithScores(key string, start, stop int64) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRangeWithScores(c.context, key, start, stop))
}

func (c *client) ZRangeByScore(key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRangeByScore(c.context, key, opt))
}

func (c *client) ZRangeByLex(key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRangeByLex(c.context, key, opt))
}

func (c *client) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRangeByScoreWithScores(c.context, key, opt))
}

func (c *client) ZRangeStore(dst string, z redis.ZRangeArgs) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRangeStore(c.context, dst, z))
}

func (c *client) ZRank(key, member string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRank(c.context, key, member))
}

func (c *client) ZRem(key string, members ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRem(c.context, key, members...))
}

func (c *client) ZRemRangeByRank(key string, start, stop int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRemRangeByRank(c.context, key, start, stop))
}

func (c *client) ZRemRangeByScore(key, min, max string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRemRangeByScore(c.context, key, min, max))
}

func (c *client) ZRemRangeByLex(key, min, max string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRemRangeByLex(c.context, key, min, max))
}

func (c *client) ZRevRange(key string, start, stop int64) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRevRange(c.context, key, start, stop))
}

func (c *client) ZRevRangeWithScores(key string, start, stop int64) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRevRangeWithScores(c.context, key, start, stop))
}

func (c *client) ZRevRangeByScore(key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRevRangeByScore(c.context, key, opt))
}

func (c *client) ZRevRangeByLex(key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRevRangeByLex(c.context, key, opt))
}

func (c *client) ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZRevRangeByScoreWithScores(c.context, key, opt))
}

func (c *client) ZRevRank(key, member string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZRevRank(c.context, key, member))
}

func (c *client) ZScore(key, member string) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.ZScore(c.context, key, member))
}

func (c *client) ZUnion(store redis.ZStore) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZUnion(c.context, store))
}

func (c *client) ZUnionWithScores(store redis.ZStore) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZUnionWithScores(c.context, store))
}

func (c *client) ZUnionStore(dest string, store *redis.ZStore) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZUnionStore(c.context, dest, store))
}

func (c *client) ZRandMember(key string, count int, withScores bool) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRandMember(c.context, key, count, withScores))
}

func (c *client) ZDiff(keys ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZDiff(c.context, keys...))
}

func (c *client) ZDiffWithScores(keys ...string) *cmd.ZSliceCmd {
	return cmd.NewZSliceCMD(c.client.ZDiffWithScores(c.context, keys...))
}

func (c *client) ZDiffStore(destination string, keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZDiffStore(c.context, destination, keys...))
}

func (c *client) PFAdd(key string, els ...interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.PFAdd(c.context, key, els...))
}

func (c *client) PFCount(keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.PFCount(c.context, keys...))
}

func (c *client) PFMerge(dest string, keys ...string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.PFMerge(c.context, dest, keys...))
}

func (c *client) BgRewriteAOF() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.BgRewriteAOF(c.context))
}

func (c *client) BgSave() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.BgSave(c.context))
}

func (c *client) GeoAdd(key string, geoLocation ...*redis.GeoLocation) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.GeoAdd(c.context, key, geoLocation...))
}

func (c *client) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *cmd.GeoLocationCmd {
	return cmd.NewGeoLocationCMD(c.client.GeoRadius(c.context, key, longitude, latitude, query))
}

func (c *client) GeoRadiusStore(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.GeoRadiusStore(c.context, key, longitude, latitude, query))
}

func (c *client) GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) *cmd.GeoLocationCmd {
	return cmd.NewGeoLocationCMD(c.client.GeoRadiusByMember(c.context, key, member, query))
}

func (c *client) GeoRadiusByMemberStore(key, member string, query *redis.GeoRadiusQuery) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.GeoRadiusByMemberStore(c.context, key, member, query))
}

func (c *client) GeoSearch(key string, q *redis.GeoSearchQuery) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.GeoSearch(c.context, key, q))
}

func (c *client) GeoSearchLocation(key string, q *redis.GeoSearchLocationQuery) *cmd.GeoSearchLocationCmd {
	return cmd.NewGeoSearchLocationCMD(c.client.GeoSearchLocation(c.context, key, q))
}

func (c *client) GeoSearchStore(key, store string, q *redis.GeoSearchStoreQuery) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.GeoSearchStore(c.context, key, store, q))
}

func (c *client) GeoDist(key string, member1, member2, unit string) *cmd.FloatCmd {
	return cmd.NewFloatCMD(c.client.GeoDist(c.context, key, member1, member2, unit))
}

func (c *client) GeoHash(key string, members ...string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.GeoHash(c.context, key, members...))
}

func (c *client) GeoPos(key string, members ...string) *cmd.GeoPosCmd {
	return cmd.NewGeoPosCMD(c.client.GeoPos(c.context, key, members...))
}

func (c *client) Watch(fn func(tx *redis.Tx) error, keys ...string) error {
	return c.client.Watch(c.context, fn, keys...)
}

func (c *client) Wait(numSlaves int, timeout time.Duration) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Wait(c.context, numSlaves, timeout))
}

// Pipeline 网络优化器, 通过 fn 函数缓冲一堆命令并一次性将它们发送到服务器执行, 好处是 节省了每个命令的网络往返时间（RTT）
func (c *client) Pipeline(fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error) {
	pipeliner := c.client.Pipeline()
	fn(pipeliner)
	return pipeliner.Exec(c.context)
}

// TxPipeline 事务 - 类似 Pipeline, 但是它内部会使用 MULTI/EXEC 包裹排队的命令
func (c *client) TxPipeline(fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error) {
	txPipeline := c.client.TxPipeline()
	fn(txPipeline)
	return txPipeline.Exec(c.context)
}

func (c *client) ClientKill(ipPort string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClientKill(c.context, ipPort))
}

func (c *client) ClientKillByFilter(keys ...string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClientKillByFilter(c.context, keys...))
}

func (c *client) ClientList() *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ClientList(c.context))
}

func (c *client) ClientPause(dur time.Duration) *cmd.BoolCmd {
	return cmd.NewBoolCMD(c.client.ClientPause(c.context, dur))
}

func (c *client) ClientID() *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClientID(c.context))
}

func (c *client) ClientUnblock(id int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClientUnblock(c.context, id))
}

func (c *client) ClientUnblockWithError(id int64) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClientUnblockWithError(c.context, id))
}

func (c *client) ConfigGet(parameter string) *cmd.SliceCmd {
	return cmd.NewSliceCMD(c.client.ConfigGet(c.context, parameter))
}

func (c *client) ConfigResetStat() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ConfigResetStat(c.context))
}

func (c *client) ConfigSet(parameter, value string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ConfigSet(c.context, parameter, value))
}

func (c *client) ConfigRewrite() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ConfigRewrite(c.context))
}

func (c *client) DBSize() *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.DBSize(c.context))
}

func (c *client) FlushAll() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.FlushAll(c.context))
}

func (c *client) FlushAllAsync() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.FlushAllAsync(c.context))
}

func (c *client) FlushDB() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.FlushDB(c.context))
}

func (c *client) FlushDBAsync() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.FlushDBAsync(c.context))
}

func (c *client) Info(section ...string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.Info(c.context, section...))
}

func (c *client) LastSave() *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.LastSave(c.context))
}

func (c *client) Save() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Save(c.context))
}

func (c *client) Shutdown() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Shutdown(c.context))
}

func (c *client) ShutdownSave() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ShutdownSave(c.context))
}

func (c *client) ShutdownNoSave() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ShutdownNoSave(c.context))
}

func (c *client) SlaveOf(host, port string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.SlaveOf(c.context, host, port))
}

func (c *client) SlowLogGet(num int64) *cmd.SlowLogCmd {
	return cmd.NewSlowLogCMD(c.client.SlowLogGet(c.context, num))
}

func (c *client) Time() *cmd.TimeCmd {
	return cmd.NewTimeCMD(c.client.Time(c.context))
}

func (c *client) DebugObject(key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.DebugObject(c.context, key))
}

func (c *client) ReadOnly() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ReadOnly(c.context))
}

func (c *client) ReadWrite() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ReadWrite(c.context))
}

func (c *client) MemoryUsage(key string, samples ...int) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.MemoryUsage(c.context, key, samples...))
}

func (c *client) Eval(script string, keys []string, args ...interface{}) *cmd.Cmd {
	return cmd.NewCMD(c.client.Eval(c.context, script, keys, args...))
}

func (c *client) EvalSha(sha1 string, keys []string, args ...interface{}) *cmd.Cmd {
	return cmd.NewCMD(c.client.EvalSha(c.context, sha1, keys, args...))
}

func (c *client) ScriptExists(hashes ...string) *cmd.BoolSliceCmd {
	return cmd.NewBoolSliceCMD(c.client.ScriptExists(c.context, hashes...))
}

func (c *client) ScriptFlush() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ScriptFlush(c.context))
}

func (c *client) ScriptKill() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ScriptKill(c.context))
}

func (c *client) ScriptLoad(script string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ScriptLoad(c.context, script))
}

func (c *client) Publish(channel string, message interface{}) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.Publish(c.context, channel, message))
}

func (c *client) PubSubChannels(pattern string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.PubSubChannels(c.context, pattern))
}

func (c *client) PubSubNumSub(channels ...string) *cmd.StringIntMapCmd {
	return cmd.NewStringIntMapCMD(c.client.PubSubNumSub(c.context, channels...))
}

func (c *client) PubSubNumPat() *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.PubSubNumPat(c.context))
}

func (c *client) ClusterSlots() *cmd.ClusterSlotsCmd {
	return cmd.NewClusterSlotsCMD(c.client.ClusterSlots(c.context))
}

func (c *client) ClusterNodes() *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ClusterNodes(c.context))
}

func (c *client) ClusterMeet(host, port string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterMeet(c.context, host, port))
}

func (c *client) ClusterForget(nodeID string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterForget(c.context, nodeID))
}

func (c *client) ClusterReplicate(nodeID string) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterReplicate(c.context, nodeID))
}

func (c *client) ClusterResetSoft() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterResetSoft(c.context))
}

func (c *client) ClusterResetHard() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterResetHard(c.context))
}

func (c *client) ClusterInfo() *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.ClusterInfo(c.context))
}

func (c *client) ClusterKeySlot(key string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClusterKeySlot(c.context, key))
}

func (c *client) ClusterGetKeysInSlot(slot int, count int) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ClusterGetKeysInSlot(c.context, slot, count))
}

func (c *client) ClusterCountFailureReports(nodeID string) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClusterCountFailureReports(c.context, nodeID))
}

func (c *client) ClusterCountKeysInSlot(slot int) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ClusterCountKeysInSlot(c.context, slot))
}

func (c *client) ClusterDelSlots(slots ...int) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterDelSlots(c.context, slots...))
}

func (c *client) ClusterDelSlotsRange(min, max int) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterDelSlotsRange(c.context, min, max))
}

func (c *client) ClusterSaveConfig() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterSaveConfig(c.context))
}

func (c *client) ClusterSlaves(nodeID string) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ClusterSlaves(c.context, nodeID))
}

func (c *client) ClusterFailover() *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterFailover(c.context))
}

func (c *client) ClusterAddSlots(slots ...int) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterAddSlots(c.context, slots...))
}

func (c *client) ClusterAddSlotsRange(min, max int) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.ClusterAddSlotsRange(c.context, min, max))
}

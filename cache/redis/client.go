package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/raylin666/go-utils/cache/redis/cmd"
	"time"
)

var _ Client = (*client)(nil)

type Client interface {
	Close() error
	Ping(ctx context.Context) *cmd.StatusCmd
	Command(ctx context.Context) *cmd.CommandsInfoCmd
	ClientGetName(ctx context.Context) *cmd.StringCmd
	Echo(ctx context.Context, message interface{}) *cmd.StringCmd
	Keys(ctx context.Context, pattern string) *cmd.StringSliceCmd
	Dump(ctx context.Context, key string) *cmd.StringCmd
	Get(ctx context.Context, key string) *cmd.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.StatusCmd
	SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.StatusCmd
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.BoolCmd
	SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.BoolCmd
	Del(ctx context.Context, keys ...string) *cmd.IntCmd
	GetRange(ctx context.Context, key string, start, end int64) *cmd.StringCmd
	GetSet(ctx context.Context, key string, value interface{}) *cmd.StringCmd
	GetEx(ctx context.Context, key string, expiration time.Duration) *cmd.StringCmd
	GetDel(ctx context.Context, key string) *cmd.StringCmd
	StrLen(ctx context.Context, key string) *cmd.IntCmd
	Incr(ctx context.Context, key string) *cmd.IntCmd
	Decr(ctx context.Context, key string) *cmd.IntCmd
	IncrBy(ctx context.Context, key string, value int64) *cmd.IntCmd
	DecrBy(ctx context.Context, key string, decrement int64) *cmd.IntCmd
	IncrByFloat(ctx context.Context, key string, value float64) *cmd.FloatCmd
	Expire(ctx context.Context, key string, expiration time.Duration) *cmd.BoolCmd
	ExpireAt(ctx context.Context, key string, tm time.Time) *cmd.BoolCmd
	PExpire(ctx context.Context, key string, expiration time.Duration) *cmd.BoolCmd
	PExpireAt(ctx context.Context, key string, tm time.Time) *cmd.BoolCmd
	TTL(ctx context.Context, key string) *cmd.DurationCmd
	PTTL(ctx context.Context, key string) *cmd.DurationCmd
	Exists(ctx context.Context, keys ...string) *cmd.IntCmd
	Unlink(ctx context.Context, keys ...string) *cmd.IntCmd
	Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *cmd.StatusCmd
	Move(ctx context.Context, key string, db int) *cmd.BoolCmd
	ObjectRefCount(ctx context.Context, key string) *cmd.IntCmd
	ObjectEncoding(ctx context.Context, key string) *cmd.StringCmd
	ObjectIdleTime(ctx context.Context, key string) *cmd.DurationCmd
	RandomKey(ctx context.Context) *cmd.StringCmd
	Rename(ctx context.Context, key string, newkey string) *cmd.StatusCmd
	RenameNX(ctx context.Context, key string, newkey string) *cmd.BoolCmd
	Type(ctx context.Context, key string) *cmd.StatusCmd
	Append(ctx context.Context, key, value string) *cmd.IntCmd
	MGet(ctx context.Context, keys ...string) *cmd.SliceCmd
	MSet(ctx context.Context, values ...interface{}) *cmd.StatusCmd
	MSetNX(ctx context.Context, values ...interface{}) *cmd.BoolCmd
	GetBit(ctx context.Context, key string, offset int64) *cmd.IntCmd
	SetBit(ctx context.Context, key string, offset int64, value int) *cmd.IntCmd
	BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *cmd.IntCmd
	BitOpAnd(ctx context.Context, destKey string, keys ...string) *cmd.IntCmd
	BitOpOr(ctx context.Context, destKey string, keys ...string) *cmd.IntCmd
	BitOpXor(ctx context.Context, destKey string, keys ...string) *cmd.IntCmd
	BitOpNot(ctx context.Context, destKey string, key string) *cmd.IntCmd
	BitPos(ctx context.Context, key string, bit int64, pos ...int64) *cmd.IntCmd
	BitField(ctx context.Context, key string, args ...interface{}) *cmd.IntSliceCmd
	SetArgs(ctx context.Context, key string, value interface{}, args redis.SetArgs) *cmd.StatusCmd
	Scan(ctx context.Context, cursor uint64, match string, count int64) *cmd.ScanCmd
	ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *cmd.ScanCmd
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *cmd.ScanCmd
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *cmd.ScanCmd
	ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *cmd.ScanCmd
	HDel(ctx context.Context, key string, fields ...string) *cmd.IntCmd
	HExists(ctx context.Context, key, field string) *cmd.BoolCmd
	HGet(ctx context.Context, key, field string) *cmd.StringCmd
	HGetAll(ctx context.Context, key string) *cmd.StringStringMapCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) *cmd.IntCmd
	HIncrByFloat(ctx context.Context, key, field string, incr float64) *cmd.FloatCmd
	HKeys(ctx context.Context, key string) *cmd.StringSliceCmd
	HLen(ctx context.Context, key string) *cmd.IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *cmd.SliceCmd
	HSet(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *cmd.BoolCmd
	HSetNX(ctx context.Context, key, field string, value interface{}) *cmd.BoolCmd
	HVals(ctx context.Context, key string) *cmd.StringSliceCmd
	HRandField(ctx context.Context, key string, count int, withValues bool) *cmd.StringSliceCmd
	BLPop(ctx context.Context, timeout time.Duration, keys ...string) *cmd.StringSliceCmd
	BRPop(ctx context.Context, timeout time.Duration, keys ...string) *cmd.StringSliceCmd
	BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *cmd.StringCmd
	LIndex(ctx context.Context, key string, index int64) *cmd.StringCmd
	LInsert(ctx context.Context, key, op string, pivot, value interface{}) *cmd.IntCmd
	LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *cmd.IntCmd
	LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *cmd.IntCmd
	LLen(ctx context.Context, key string) *cmd.IntCmd
	LPop(ctx context.Context, key string) *cmd.StringCmd
	LPopCount(ctx context.Context, key string, count int) *cmd.StringSliceCmd
	LPos(ctx context.Context, key string, value string, a redis.LPosArgs) *cmd.IntCmd
	LPosCount(ctx context.Context, key string, value string, count int64, a redis.LPosArgs) *cmd.IntSliceCmd
	LPush(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd
	LPushX(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd
	LRange(ctx context.Context, key string, start, stop int64) *cmd.StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) *cmd.IntCmd
	LSet(ctx context.Context, key string, index int64, value interface{}) *cmd.StatusCmd
	LTrim(ctx context.Context, key string, start int64, stop int64) *cmd.StatusCmd
	RPop(ctx context.Context, key string) *cmd.StringCmd
	RPopCount(ctx context.Context, key string, count int) *cmd.StringSliceCmd
	RPopLPush(ctx context.Context, source, destination string) *cmd.StringCmd
	RPush(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd
	RPushX(ctx context.Context, key string, values ...interface{}) *cmd.IntCmd
	LMove(ctx context.Context, source, destination, srcpos, destpos string) *cmd.StringCmd
	BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *cmd.StringCmd
	SAdd(ctx context.Context, key string, members ...interface{}) *cmd.IntCmd
	SCard(ctx context.Context, key string) *cmd.IntCmd
	SDiff(ctx context.Context, keys ...string) *cmd.StringSliceCmd
	SDiffStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd
	SInter(ctx context.Context, keys ...string) *cmd.StringSliceCmd
	SInterStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd
	SIsMember(ctx context.Context, key string, member interface{}) *cmd.BoolCmd
	SMIsMember(ctx context.Context, key string, members ...interface{}) *cmd.BoolSliceCmd
	SMembers(ctx context.Context, key string) *cmd.StringSliceCmd
	SMembersMap(ctx context.Context, key string) *cmd.StringStructMapCmd
	SMove(ctx context.Context, source, destination string, member interface{}) *cmd.BoolCmd
	SPop(ctx context.Context, key string) *cmd.StringCmd
	SPopN(ctx context.Context, key string, count int64) *cmd.StringSliceCmd
	SRandMember(ctx context.Context, key string) *cmd.StringCmd
	SRandMemberN(ctx context.Context, key string, count int64) *cmd.StringSliceCmd
	SRem(ctx context.Context, key string, members ...interface{}) *cmd.IntCmd
	SUnion(ctx context.Context, keys ...string) *cmd.StringSliceCmd
	SUnionStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd
	ZAddArgs(ctx context.Context, key string, args redis.ZAddArgs) *cmd.IntCmd
	ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *cmd.FloatCmd
	ZAdd(ctx context.Context, key string, members ...*redis.Z) *cmd.IntCmd
	ZAddNX(ctx context.Context, key string, members ...*redis.Z) *cmd.IntCmd
	ZAddXX(ctx context.Context, key string, members ...*redis.Z) *cmd.IntCmd
	ZCard(ctx context.Context, key string) *cmd.IntCmd
	ZCount(ctx context.Context, key, min, max string) *cmd.IntCmd
	ZLexCount(ctx context.Context, key, min, max string) *cmd.IntCmd
	ZIncrBy(ctx context.Context, key string, increment float64, member string) *cmd.FloatCmd
	ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *cmd.IntCmd
	ZInter(ctx context.Context, store *redis.ZStore) *cmd.StringSliceCmd
	ZInterWithScores(ctx context.Context, store *redis.ZStore) *cmd.ZSliceCmd
	ZMScore(ctx context.Context, key string, members ...string) *cmd.FloatSliceCmd
	ZPopMax(ctx context.Context, key string, count ...int64) *cmd.ZSliceCmd
	ZPopMin(ctx context.Context, key string, count ...int64) *cmd.ZSliceCmd
	ZRangeArgs(ctx context.Context, z redis.ZRangeArgs) *cmd.StringSliceCmd
	ZRangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *cmd.ZSliceCmd
	ZRange(ctx context.Context, key string, start, stop int64) *cmd.StringSliceCmd
	ZRangeWithScores(ctx context.Context, key string, start, stop int64) *cmd.ZSliceCmd
	ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd
	ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd
	ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.ZSliceCmd
	ZRangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *cmd.IntCmd
	ZRank(ctx context.Context, key, member string) *cmd.IntCmd
	ZRem(ctx context.Context, key string, members ...interface{}) *cmd.IntCmd
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *cmd.IntCmd
	ZRemRangeByScore(ctx context.Context, key, min, max string) *cmd.IntCmd
	ZRemRangeByLex(ctx context.Context, key, min, max string) *cmd.IntCmd
	ZRevRange(ctx context.Context, key string, start, stop int64) *cmd.StringSliceCmd
	ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *cmd.ZSliceCmd
	ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd
	ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.StringSliceCmd
	ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *cmd.ZSliceCmd
	ZRevRank(ctx context.Context, key, member string) *cmd.IntCmd
	ZScore(ctx context.Context, key, member string) *cmd.FloatCmd
	ZUnion(ctx context.Context, store redis.ZStore) *cmd.StringSliceCmd
	ZUnionWithScores(ctx context.Context, store redis.ZStore) *cmd.ZSliceCmd
	ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *cmd.IntCmd
	ZRandMember(ctx context.Context, key string, count int, withScores bool) *cmd.StringSliceCmd
	ZDiff(ctx context.Context, keys ...string) *cmd.StringSliceCmd
	ZDiffWithScores(ctx context.Context, keys ...string) *cmd.ZSliceCmd
	ZDiffStore(ctx context.Context, destination string, keys ...string) *cmd.IntCmd
	PFAdd(ctx context.Context, key string, els ...interface{}) *cmd.IntCmd
	PFCount(ctx context.Context, keys ...string) *cmd.IntCmd
	PFMerge(ctx context.Context, dest string, keys ...string) *cmd.StatusCmd
	BgRewriteAOF(ctx context.Context) *cmd.StatusCmd
	BgSave(ctx context.Context) *cmd.StatusCmd
	GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *cmd.IntCmd
	GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *cmd.GeoLocationCmd
	GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *cmd.IntCmd
	GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *cmd.GeoLocationCmd
	GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *cmd.IntCmd
	GeoSearch(ctx context.Context, key string, q *redis.GeoSearchQuery) *cmd.StringSliceCmd
	GeoSearchLocation(ctx context.Context, key string, q *redis.GeoSearchLocationQuery) *cmd.GeoSearchLocationCmd
	GeoSearchStore(ctx context.Context, key, store string, q *redis.GeoSearchStoreQuery) *cmd.IntCmd
	GeoDist(ctx context.Context, key string, member1, member2, unit string) *cmd.FloatCmd
	GeoHash(ctx context.Context, key string, members ...string) *cmd.StringSliceCmd
	GeoPos(ctx context.Context, key string, members ...string) *cmd.GeoPosCmd
	Watch(ctx context.Context, fn func(tx *redis.Tx) error, keys ...string) error
	Wait(ctx context.Context, numSlaves int, timeout time.Duration) *cmd.IntCmd
	Pipeline(ctx context.Context, fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error)
	TxPipeline(ctx context.Context, fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error)
	ClientKill(ctx context.Context, ipPort string) *cmd.StatusCmd
	ClientKillByFilter(ctx context.Context, keys ...string) *cmd.IntCmd
	ClientList(ctx context.Context) *cmd.StringCmd
	ClientPause(ctx context.Context, dur time.Duration) *cmd.BoolCmd
	ClientID(ctx context.Context) *cmd.IntCmd
	ClientUnblock(ctx context.Context, id int64) *cmd.IntCmd
	ClientUnblockWithError(ctx context.Context, id int64) *cmd.IntCmd
	ConfigGet(ctx context.Context, parameter string) *cmd.SliceCmd
	ConfigResetStat(ctx context.Context) *cmd.StatusCmd
	ConfigSet(ctx context.Context, parameter, value string) *cmd.StatusCmd
	ConfigRewrite(ctx context.Context) *cmd.StatusCmd
	DBSize(ctx context.Context) *cmd.IntCmd
	FlushAll(ctx context.Context) *cmd.StatusCmd
	FlushAllAsync(ctx context.Context) *cmd.StatusCmd
	FlushDB(ctx context.Context) *cmd.StatusCmd
	FlushDBAsync(ctx context.Context) *cmd.StatusCmd
	Info(ctx context.Context, section ...string) *cmd.StringCmd
	LastSave(ctx context.Context) *cmd.IntCmd
	Save(ctx context.Context) *cmd.StatusCmd
	Shutdown(ctx context.Context) *cmd.StatusCmd
	ShutdownSave(ctx context.Context) *cmd.StatusCmd
	ShutdownNoSave(ctx context.Context) *cmd.StatusCmd
	SlaveOf(ctx context.Context, host, port string) *cmd.StatusCmd
	SlowLogGet(ctx context.Context, num int64) *cmd.SlowLogCmd
	Time(ctx context.Context) *cmd.TimeCmd
	DebugObject(ctx context.Context, key string) *cmd.StringCmd
	ReadOnly(ctx context.Context) *cmd.StatusCmd
	ReadWrite(ctx context.Context) *cmd.StatusCmd
	MemoryUsage(ctx context.Context, key string, samples ...int) *cmd.IntCmd
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *cmd.Cmd
	EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *cmd.Cmd
	ScriptExists(ctx context.Context, hashes ...string) *cmd.BoolSliceCmd
	ScriptFlush(ctx context.Context) *cmd.StatusCmd
	ScriptKill(ctx context.Context) *cmd.StatusCmd
	ScriptLoad(ctx context.Context, script string) *cmd.StringCmd
	Publish(ctx context.Context, channel string, message interface{}) *cmd.IntCmd
	PubSubChannels(ctx context.Context, pattern string) *cmd.StringSliceCmd
	PubSubNumSub(ctx context.Context, channels ...string) *cmd.StringIntMapCmd
	PubSubNumPat(ctx context.Context) *cmd.IntCmd
	ClusterSlots(ctx context.Context) *cmd.ClusterSlotsCmd
	ClusterNodes(ctx context.Context) *cmd.StringCmd
	ClusterMeet(ctx context.Context, host, port string) *cmd.StatusCmd
	ClusterForget(ctx context.Context, nodeID string) *cmd.StatusCmd
	ClusterReplicate(ctx context.Context, nodeID string) *cmd.StatusCmd
	ClusterResetSoft(ctx context.Context) *cmd.StatusCmd
	ClusterResetHard(ctx context.Context) *cmd.StatusCmd
	ClusterInfo(ctx context.Context) *cmd.StringCmd
	ClusterKeySlot(ctx context.Context, key string) *cmd.IntCmd
	ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *cmd.StringSliceCmd
	ClusterCountFailureReports(ctx context.Context, nodeID string) *cmd.IntCmd
	ClusterCountKeysInSlot(ctx context.Context, slot int) *cmd.IntCmd
	ClusterDelSlots(ctx context.Context, slots ...int) *cmd.StatusCmd
	ClusterDelSlotsRange(ctx context.Context, min, max int) *cmd.StatusCmd
	ClusterSaveConfig(ctx context.Context) *cmd.StatusCmd
	ClusterSlaves(ctx context.Context, nodeID string) *cmd.StringSliceCmd
	ClusterFailover(ctx context.Context) *cmd.StatusCmd
	ClusterAddSlots(ctx context.Context, slots ...int) *cmd.StatusCmd
	ClusterAddSlotsRange(ctx context.Context, min, max int) *cmd.StatusCmd
}

type Options struct {
	redis.Options
}

type client struct {
	client  *redis.Client
	options *Options
}

func NewClient(ctx context.Context, opt *Options) (Client, error) {
	var c = new(client)
	c.options = opt
	c.client = redis.NewClient(&opt.Options)
	if err := c.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *client) Close() error {
	return c.client.Close()
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

func (c *client) Dump(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.Dump(ctx, key))
}

func (c *client) Get(ctx context.Context, key string) *cmd.StringCmd {
	return cmd.NewStringCMD(c.client.Get(ctx, key))
}

func (c *client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.Set(ctx, key, value, expiration))
}

func (c *client) SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *cmd.StatusCmd {
	return cmd.NewStatusCMD(c.client.SetEX(ctx, key, value, expiration))
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

func (c *client) HRandField(ctx context.Context, key string, count int, withValues bool) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.HRandField(ctx, key, count, withValues))
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

func (c *client) ZAdd(ctx context.Context, key string, members ...*redis.Z) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZAdd(ctx, key, members...))
}

func (c *client) ZAddNX(ctx context.Context, key string, members ...*redis.Z) *cmd.IntCmd {
	return cmd.NewIntCMD(c.client.ZAddNX(ctx, key, members...))
}

func (c *client) ZAddXX(ctx context.Context, key string, members ...*redis.Z) *cmd.IntCmd {
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

func (c *client) ZRandMember(ctx context.Context, key string, count int, withScores bool) *cmd.StringSliceCmd {
	return cmd.NewStringSliceCMD(c.client.ZRandMember(ctx, key, count, withScores))
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

// Pipeline 网络优化器, 通过 fn 函数缓冲一堆命令并一次性将它们发送到服务器执行, 好处是 节省了每个命令的网络往返时间（RTT）
func (c *client) Pipeline(ctx context.Context, fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error) {
	pipeliner := c.client.Pipeline()
	fn(pipeliner)
	return pipeliner.Exec(ctx)
}

// TxPipeline 事务 - 类似 Pipeline, 但是它内部会使用 MULTI/EXEC 包裹排队的命令
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

func (c *client) ConfigGet(ctx context.Context, parameter string) *cmd.SliceCmd {
	return cmd.NewSliceCMD(c.client.ConfigGet(ctx, parameter))
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

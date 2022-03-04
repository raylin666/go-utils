package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Client struct {
	Conn 	*redis.Client
	Options *Options
	Context context.Context
}

type Options struct {
	redis.Options
}

func New(ctx context.Context, options *Options) (*Client, error) {
	var client = new(Client)
	client.Context = ctx
	client.Options = options
	client.Conn = redis.NewClient(&options.Options)
	_, err := client.Ping()
	return client, err
}

func (c *Client) Ping() (string, error) {
	return c.Conn.Ping(c.Context).Result()
}

func (c *Client) Command() (map[string]*redis.CommandInfo, error) {
	return c.Conn.Command(c.Context).Result()
}

func (c *Client) ClientGetName() *redis.StringCmd {
	return c.Conn.ClientGetName(c.Context)
}

func (c *Client) Echo(message interface{}) *redis.StringCmd {
	return c.Conn.Echo(c.Context, message)
}

func (c *Client) Keys(pattern string) ([]string, error) {
	return c.Conn.Keys(c.Context, pattern).Result()
}

func (c *Client) Dump(key string) *redis.StringCmd {
	return c.Conn.Dump(c.Context, key)
}

func (c *Client) Get(key string) *redis.StringCmd {
	return c.Conn.Get(c.Context, key)
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return c.Conn.Set(c.Context, key, value, expiration).Result()
}

func (c *Client) SetEX(key string, value interface{}, expiration time.Duration) (string, error) {
	return c.Conn.SetEX(c.Context, key, value, expiration).Result()
}

func (c *Client) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.Conn.SetNX(c.Context, key, value, expiration).Result()
}

func (c *Client) SetXX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.Conn.SetXX(c.Context, key, value, expiration).Result()
}

func (c *Client) Del(keys ...string) (int64, error) {
	return c.Conn.Del(c.Context, keys...).Result()
}

func (c *Client) GetRange(key string, start, end int64) *redis.StringCmd {
	return c.Conn.GetRange(c.Context, key, start, end)
}

func (c *Client) GetSet(key string, value interface{}) *redis.StringCmd {
	return c.Conn.GetSet(c.Context, key, value)
}

func (c *Client) GetEx(key string, expiration time.Duration) *redis.StringCmd {
	return c.Conn.GetEx(c.Context, key, expiration)
}

func (c *Client) GetDel(key string) *redis.StringCmd {
	return c.Conn.GetDel(c.Context, key)
}

func (c *Client) StrLen(key string) (int64, error) {
	return c.Conn.StrLen(c.Context, key).Result()
}

func (c *Client) Incr(key string) (int64, error) {
	return c.Conn.Incr(c.Context, key).Result()
}

func (c *Client) Decr(key string) (int64, error) {
	return c.Conn.Decr(c.Context, key).Result()
}

func (c *Client) IncrBy(key string, value int64) (int64, error) {
	return c.Conn.IncrBy(c.Context, key, value).Result()
}

func (c *Client) DecrBy(key string, decrement int64) (int64, error) {
	return c.Conn.DecrBy(c.Context, key, decrement).Result()
}

func (c *Client) IncrByFloat(key string, value float64) (float64, error) {
	return c.Conn.IncrByFloat(c.Context, key, value).Result()
}

func (c *Client) Expire(key string, expiration time.Duration) (bool, error) {
	return c.Conn.Expire(c.Context, key, expiration).Result()
}

func (c *Client) ExpireAt(key string, tm time.Time) (bool, error) {
	return c.Conn.ExpireAt(c.Context, key, tm).Result()
}

func (c *Client) PExpire(key string, expiration time.Duration) (bool, error) {
	return c.Conn.PExpire(c.Context, key, expiration).Result()
}

func (c *Client) PExpireAt(key string, tm time.Time) (bool, error) {
	return c.Conn.PExpireAt(c.Context, key, tm).Result()
}

func (c *Client) TTL(key string) (time.Duration, error) {
	return c.Conn.TTL(c.Context, key).Result()
}

func (c *Client) PTTL(key string) (time.Duration, error) {
	return c.Conn.PTTL(c.Context, key).Result()
}

func (c *Client) Exists(keys ...string) (int64, error) {
	return c.Conn.Exists(c.Context, keys...).Result()
}

func (c *Client) Unlink(keys ...string) (int64, error) {
	return c.Conn.Unlink(c.Context, keys...).Result()
}

func (c *Client) Migrate(host, port, key string, db int, timeout time.Duration) (string, error) {
	return c.Conn.Migrate(c.Context, host, port, key, db, timeout).Result()
}

func (c *Client) Move(key string, db int) (bool, error) {
	return c.Conn.Move(c.Context, key, db).Result()
}

func (c *Client) ObjectRefCount(key string) (int64, error) {
	return c.Conn.ObjectRefCount(c.Context, key).Result()
}

func (c *Client) ObjectEncoding(key string) *redis.StringCmd {
	return c.Conn.ObjectEncoding(c.Context, key)
}

func (c *Client) ObjectIdleTime(key string) (time.Duration, error) {
	return c.Conn.ObjectIdleTime(c.Context, key).Result()
}

func (c *Client) RandomKey() *redis.StringCmd {
	return c.Conn.RandomKey(c.Context)
}

func (c *Client) Rename(key string, newkey string) (string, error) {
	return c.Conn.Rename(c.Context, key, newkey).Result()
}

func (c *Client) RenameNX(key string, newkey string) (bool, error) {
	return c.Conn.RenameNX(c.Context, key, newkey).Result()
}

func (c *Client) Type(key string) (string, error) {
	return c.Conn.Type(c.Context, key).Result()
}

func (c *Client) Append(key, value string) (int64, error) {
	return c.Conn.Append(c.Context, key, value).Result()
}

func (c *Client) MGet(keys ...string) ([]interface{}, error) {
	return c.Conn.MGet(c.Context, keys...).Result()
}

func (c *Client) MSet(values ...interface{}) (string, error) {
	return c.Conn.MSet(c.Context, values...).Result()
}

func (c *Client) MSetNX(values ...interface{}) (bool, error) {
	return c.Conn.MSetNX(c.Context, values...).Result()
}

func (c *Client) GetBit(key string, offset int64) (int64, error) {
	return c.Conn.GetBit(c.Context, key, offset).Result()
}

func (c *Client) SetBit(key string, offset int64, value int) (int64, error) {
	return c.Conn.SetBit(c.Context, key, offset, value).Result()
}

func (c *Client) BitCount(key string, bitCount *redis.BitCount) (int64, error) {
	return c.Conn.BitCount(c.Context, key, bitCount).Result()
}

func (c *Client) BitOpAnd(destKey string, keys ...string) (int64, error) {
	return c.Conn.BitOpAnd(c.Context, destKey, keys...).Result()
}

func (c *Client) BitOpOr(destKey string, keys ...string) (int64, error) {
	return c.Conn.BitOpOr(c.Context, destKey, keys...).Result()
}

func (c *Client) BitOpXor(destKey string, keys ...string) (int64, error) {
	return c.Conn.BitOpXor(c.Context, destKey, keys...).Result()
}

func (c *Client) BitOpNot(destKey string, key string) (int64, error) {
	return c.Conn.BitOpNot(c.Context, destKey, key).Result()
}

func (c *Client) BitPos(key string, bit int64, pos ...int64) (int64, error) {
	return c.Conn.BitPos(c.Context, key, bit, pos...).Result()
}

func (c *Client) BitField(key string, args ...interface{}) ([]int64, error) {
	return c.Conn.BitField(c.Context, key, args...).Result()
}

func (c *Client) SetArgs(key string, value interface{}, args redis.SetArgs) (string, error) {
	return c.Conn.SetArgs(c.Context, key, value, args).Result()
}

func (c *Client) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.Conn.Scan(c.Context, cursor, match, count).Result()
}

func (c *Client) ScanType(cursor uint64, match string, count int64, keyType string) ([]string, uint64, error) {
	return c.Conn.ScanType(c.Context, cursor, match, count, keyType).Result()
}

func (c *Client) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.Conn.SScan(c.Context, key, cursor, match, count).Result()
}

func (c *Client) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.Conn.HScan(c.Context, key, cursor, match, count).Result()
}

func (c *Client) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.Conn.ZScan(c.Context, key, cursor, match, count).Result()
}

func (c *Client) HDel(key string, fields ...string) (int64, error) {
	return c.Conn.HDel(c.Context, key, fields...).Result()
}

func (c *Client) HExists(key, field string) (bool, error) {
	return c.Conn.HExists(c.Context, key, field).Result()
}

func (c *Client) HGet(key, field string) *redis.StringCmd {
	return c.Conn.HGet(c.Context, key, field)
}

func (c *Client) HGetAll(key string) (map[string]string, error) {
	return c.Conn.HGetAll(c.Context, key).Result()
}

func (c *Client) HIncrBy(key, field string, incr int64) (int64, error) {
	return c.Conn.HIncrBy(c.Context, key, field, incr).Result()
}

func (c *Client) HIncrByFloat(key, field string, incr float64) (float64, error) {
	return c.Conn.HIncrByFloat(c.Context, key, field, incr).Result()
}

func (c *Client) HKeys(key string) ([]string, error) {
	return c.Conn.HKeys(c.Context, key).Result()
}

func (c *Client) HLen(key string) (int64, error) {
	return c.Conn.HLen(c.Context, key).Result()
}

func (c *Client) HMGet(key string, fields ...string) ([]interface{}, error) {
	return c.Conn.HMGet(c.Context, key, fields...).Result()
}

func (c *Client) HSet(key string, values ...interface{}) (int64, error) {
	return c.Conn.HSet(c.Context, key, values...).Result()
}

func (c *Client) HMSet(key string, values ...interface{}) (bool, error) {
	return c.Conn.HMSet(c.Context, key, values...).Result()
}

func (c *Client) HSetNX(key, field string, value interface{}) (bool, error) {
	return c.Conn.HSetNX(c.Context, field, key, value).Result()
}

func (c *Client) HVals(key string) ([]string, error) {
	return c.Conn.HVals(c.Context, key).Result()
}

func (c *Client) HRandField(key string, count int, withValues bool) ([]string, error) {
	return c.Conn.HRandField(c.Context, key, count, withValues).Result()
}

func (c *Client) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return c.Conn.BLPop(c.Context, timeout, keys...).Result()
}

func (c *Client) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	return c.Conn.BRPop(c.Context, timeout, keys...).Result()
}

func (c *Client) BRPopLPush(source, destination string, timeout time.Duration) *redis.StringCmd {
	return c.Conn.BRPopLPush(c.Context, source, destination, timeout)
}

func (c *Client) LIndex(key string, index int64) *redis.StringCmd {
	return c.Conn.LIndex(c.Context, key, index)
}

func (c *Client) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	return c.Conn.LInsert(c.Context, key, op, pivot, value).Result()
}

func (c *Client) LInsertBefore(key string, pivot, value interface{}) (int64, error) {
	return c.Conn.LInsertBefore(c.Context, key, pivot, value).Result()
}

func (c *Client) LInsertAfter(key string, pivot, value interface{}) (int64, error) {
	return c.Conn.LInsertAfter(c.Context, key, pivot, value).Result()
}

func (c *Client) LLen(key string) (int64, error) {
	return c.Conn.LLen(c.Context, key).Result()
}

func (c *Client) LPop(key string) *redis.StringCmd {
	return c.Conn.LPop(c.Context, key)
}

func (c *Client) LPopCount(key string, count int) ([]string, error) {
	return c.Conn.LPopCount(c.Context, key, count).Result()
}

func (c *Client) LPos(key string, value string, a redis.LPosArgs) (int64, error) {
	return c.Conn.LPos(c.Context, key, value, a).Result()
}

func (c *Client) LPosCount(key string, value string, count int64, a redis.LPosArgs) ([]int64, error) {
	return c.Conn.LPosCount(c.Context, key, value, count, a).Result()
}

func (c *Client) LPush(key string, values ...interface{}) (int64, error) {
	return c.Conn.LPush(c.Context, key, values...).Result()
}

func (c *Client) LPushX(key string, values ...interface{}) (int64, error) {
	return c.Conn.LPushX(c.Context, key, values...).Result()
}

func (c *Client) LRange(key string, start, stop int64) ([]string, error) {
	return c.Conn.LRange(c.Context, key, start, stop).Result()
}

func (c *Client) LRem(key string, count int64, value interface{}) (int64, error) {
	return c.Conn.LRem(c.Context, key, count, value).Result()
}

func (c *Client) LSet(key string, index int64, value interface{}) (string, error) {
	return c.Conn.LSet(c.Context, key, index, value).Result()
}

func (c *Client) LTrim(key string, start int64, stop int64) (string, error) {
	return c.Conn.LTrim(c.Context, key, start, stop).Result()
}

func (c *Client) RPop(key string) *redis.StringCmd {
	return c.Conn.RPop(c.Context, key)
}

func (c *Client) RPopCount(key string, count int) ([]string, error) {
	return c.Conn.RPopCount(c.Context, key, count).Result()
}

func (c *Client) RPopLPush(source, destination string) *redis.StringCmd {
	return c.Conn.RPopLPush(c.Context, source, destination)
}

func (c *Client) RPush(key string, values ...interface{}) (int64, error) {
	return c.Conn.RPush(c.Context, key, values...).Result()
}

func (c *Client) RPushX(key string, values ...interface{}) (int64, error) {
	return c.Conn.RPushX(c.Context, key, values...).Result()
}

func (c *Client) LMove(source, destination, srcpos, destpos string) *redis.StringCmd {
	return c.Conn.LMove(c.Context, source, destination, srcpos, destpos)
}

func (c *Client) BLMove(source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd {
	return c.Conn.BLMove(c.Context, source, destination, srcpos, destpos, timeout)
}

func (c *Client) SAdd(key string, members ...interface{}) (int64, error) {
	return c.Conn.SAdd(c.Context, key, members...).Result()
}

func (c *Client) SCard(key string) (int64, error) {
	return c.Conn.SCard(c.Context, key).Result()
}

func (c *Client) SDiff(keys ...string) ([]string, error) {
	return c.Conn.SDiff(c.Context, keys...).Result()
}

func (c *Client) SDiffStore(destination string, keys ...string) (int64, error) {
	return c.Conn.SDiffStore(c.Context, destination, keys...).Result()
}

func (c *Client) SInter(keys ...string) ([]string, error) {
	return c.Conn.SInter(c.Context, keys...).Result()
}

func (c *Client) SInterStore(destination string, keys ...string) (int64, error) {
	return c.Conn.SInterStore(c.Context, destination, keys...).Result()
}

func (c *Client) SIsMember(key string, member interface{}) (bool, error) {
	return c.Conn.SIsMember(c.Context, key, member).Result()
}

func (c *Client) SMIsMember(key string, members ...interface{}) ([]bool, error) {
	return c.Conn.SMIsMember(c.Context, key, members...).Result()
}

func (c *Client) SMembers(key string) ([]string, error) {
	return c.Conn.SMembers(c.Context, key).Result()
}

func (c *Client) SMembersMap(key string) (map[string]struct{}, error) {
	return c.Conn.SMembersMap(c.Context, key).Result()
}

func (c *Client) SMove(source, destination string, member interface{}) (bool, error) {
	return c.Conn.SMove(c.Context, source, destination, member).Result()
}

func (c *Client) SPop(key string) *redis.StringCmd {
	return c.Conn.SPop(c.Context, key)
}

func (c *Client) SPopN(key string, count int64) ([]string, error) {
	return c.Conn.SPopN(c.Context, key, count).Result()
}

func (c *Client) SRandMember(key string) *redis.StringCmd {
	return c.Conn.SRandMember(c.Context, key)
}

func (c *Client) SRandMemberN(key string, count int64) ([]string, error) {
	return c.Conn.SRandMemberN(c.Context, key, count).Result()
}

func (c *Client) SRem(key string, members ...interface{}) (int64, error) {
	return c.Conn.SRem(c.Context, key, members...).Result()
}

func (c *Client) SUnion(keys ...string) ([]string, error) {
	return c.Conn.SUnion(c.Context, keys...).Result()
}

func (c *Client) SUnionStore(destination string, keys ...string) (int64, error) {
	return c.Conn.SUnionStore(c.Context, destination, keys...).Result()
}

func (c *Client) ZAddArgs(key string, args redis.ZAddArgs) (int64, error) {
	return c.Conn.ZAddArgs(c.Context, key, args).Result()
}

func (c *Client) ZAddArgsIncr(key string, args redis.ZAddArgs) (float64, error) {
	return c.Conn.ZAddArgsIncr(c.Context, key, args).Result()
}

func (c *Client) ZAdd(key string, members ...*redis.Z) (int64, error) {
	return c.Conn.ZAdd(c.Context, key, members...).Result()
}

func (c *Client) ZAddNX(key string, members ...*redis.Z) (int64, error) {
	return c.Conn.ZAddNX(c.Context, key, members...).Result()
}

func (c *Client) ZAddXX(key string, members ...*redis.Z) (int64, error) {
	return c.Conn.ZAddXX(c.Context, key, members...).Result()
}

func (c *Client) ZCard(key string) (int64, error) {
	return c.Conn.ZCard(c.Context, key).Result()
}

func (c *Client) ZCount(key, min, max string) (int64, error) {
	return c.Conn.ZCount(c.Context, key, min, max).Result()
}

func (c *Client) ZLexCount(key, min, max string) (int64, error) {
	return c.Conn.ZLexCount(c.Context, key, min, max).Result()
}

func (c *Client) ZIncrBy(key string, increment float64, member string) (float64, error) {
	return c.Conn.ZIncrBy(c.Context, key, increment, member).Result()
}

func (c *Client) ZInterStore(destination string, store *redis.ZStore) (int64, error) {
	return c.Conn.ZInterStore(c.Context, destination, store).Result()
}

func (c *Client) ZInter(store *redis.ZStore) ([]string, error) {
	return c.Conn.ZInter(c.Context, store).Result()
}

func (c *Client) ZInterWithScores(store *redis.ZStore) ([]redis.Z, error) {
	return c.Conn.ZInterWithScores(c.Context, store).Result()
}

func (c *Client) ZMScore(key string, members ...string) ([]float64, error) {
	return c.Conn.ZMScore(c.Context, key, members...).Result()
}

func (c *Client) ZPopMax(key string, count ...int64) ([]redis.Z, error) {
	return c.Conn.ZPopMax(c.Context, key, count...).Result()
}

func (c *Client) ZPopMin(key string, count ...int64) ([]redis.Z, error) {
	return c.Conn.ZPopMin(c.Context, key, count...).Result()
}

func (c *Client) ZRangeArgs(z redis.ZRangeArgs) ([]string, error) {
	return c.Conn.ZRangeArgs(c.Context, z).Result()
}

func (c *Client) ZRangeArgsWithScores(z redis.ZRangeArgs) ([]redis.Z, error) {
	return c.Conn.ZRangeArgsWithScores(c.Context, z).Result()
}

func (c *Client) ZRange(key string, start, stop int64) ([]string, error) {
	return c.Conn.ZRange(c.Context, key, start, stop).Result()
}

func (c *Client) ZRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return c.Conn.ZRangeWithScores(c.Context, key, start, stop).Result()
}

func (c *Client) ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.Conn.ZRangeByScore(c.Context, key, opt).Result()
}

func (c *Client) ZRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.Conn.ZRangeByLex(c.Context, key, opt).Result()
}

func (c *Client) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return c.Conn.ZRangeByScoreWithScores(c.Context, key, opt).Result()
}

func (c *Client) ZRangeStore(dst string, z redis.ZRangeArgs) (int64, error) {
	return c.Conn.ZRangeStore(c.Context, dst, z).Result()
}

func (c *Client) ZRank(key, member string) (int64, error) {
	return c.Conn.ZRank(c.Context, key, member).Result()
}

func (c *Client) ZRem(key string, members ...interface{}) (int64, error) {
	return c.Conn.ZRem(c.Context, key, members...).Result()
}

func (c *Client) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	return c.Conn.ZRemRangeByRank(c.Context, key, start, stop).Result()
}

func (c *Client) ZRemRangeByScore(key, min, max string) (int64, error) {
	return c.Conn.ZRemRangeByScore(c.Context, key, min, max).Result()
}

func (c *Client) ZRemRangeByLex(key, min, max string) (int64, error) {
	return c.Conn.ZRemRangeByLex(c.Context, key, min, max).Result()
}

func (c *Client) ZRevRange(key string, start, stop int64) ([]string, error) {
	return c.Conn.ZRevRange(c.Context, key, start, stop).Result()
}

func (c *Client) ZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return c.Conn.ZRevRangeWithScores(c.Context, key, start, stop).Result()
}

func (c *Client) ZRevRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.Conn.ZRevRangeByScore(c.Context, key, opt).Result()
}

func (c *Client) ZRevRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.Conn.ZRevRangeByLex(c.Context, key, opt).Result()
}

func (c *Client) ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return c.Conn.ZRevRangeByScoreWithScores(c.Context, key, opt).Result()
}

func (c *Client) ZRevRank(key, member string) (int64, error) {
	return c.Conn.ZRevRank(c.Context, key, member).Result()
}

func (c *Client) ZScore(key, member string) (float64, error) {
	return c.Conn.ZScore(c.Context, key, member).Result()
}

func (c *Client) ZUnion(store redis.ZStore) ([]string, error) {
	return c.Conn.ZUnion(c.Context, store).Result()
}

func (c *Client) ZUnionWithScores(store redis.ZStore) ([]redis.Z, error) {
	return c.Conn.ZUnionWithScores(c.Context, store).Result()
}

func (c *Client) ZUnionStore(dest string, store *redis.ZStore) (int64, error) {
	return c.Conn.ZUnionStore(c.Context, dest, store).Result()
}

func (c *Client) ZRandMember(key string, count int, withScores bool) ([]string, error) {
	return c.Conn.ZRandMember(c.Context, key, count, withScores).Result()
}

func (c *Client) ZDiff(keys ...string) ([]string, error) {
	return c.Conn.ZDiff(c.Context, keys...).Result()
}

func (c *Client) ZDiffWithScores(keys ...string) ([]redis.Z, error) {
	return c.Conn.ZDiffWithScores(c.Context, keys...).Result()
}

func (c *Client) ZDiffStore(destination string, keys ...string) (int64, error) {
	return c.Conn.ZDiffStore(c.Context, destination, keys...).Result()
}

func (c *Client) PFAdd(key string, els ...interface{}) (int64, error) {
	return c.Conn.PFAdd(c.Context, key, els...).Result()
}

func (c *Client) PFCount(keys ...string) (int64, error) {
	return c.Conn.PFCount(c.Context, keys...).Result()
}

func (c *Client) PFMerge(dest string, keys ...string) (string, error) {
	return c.Conn.PFMerge(c.Context, dest, keys...).Result()
}

func (c *Client) BgRewriteAOF() (string, error) {
	return c.Conn.BgRewriteAOF(c.Context).Result()
}

func (c *Client) BgSave() (string, error) {
	return c.Conn.BgSave(c.Context).Result()
}

func (c *Client) GeoAdd(key string, geoLocation ...*redis.GeoLocation) (int64, error) {
	return c.Conn.GeoAdd(c.Context, key, geoLocation...).Result()
}

func (c *Client) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) ([]redis.GeoLocation, error) {
	return c.Conn.GeoRadius(c.Context, key, longitude, latitude, query).Result()
}

func (c *Client) GeoRadiusStore(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) (int64, error) {
	return c.Conn.GeoRadiusStore(c.Context, key, longitude, latitude, query).Result()
}

func (c *Client) GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) ([]redis.GeoLocation, error) {
	return c.Conn.GeoRadiusByMember(c.Context, key, member, query).Result()
}

func (c *Client) GeoRadiusByMemberStore(key, member string, query *redis.GeoRadiusQuery) (int64, error) {
	return c.Conn.GeoRadiusByMemberStore(c.Context, key, member, query).Result()
}

func (c *Client) GeoSearch(key string, q *redis.GeoSearchQuery) ([]string, error) {
	return c.Conn.GeoSearch(c.Context, key, q).Result()
}

func (c *Client) GeoSearchLocation(key string, q *redis.GeoSearchLocationQuery) ([]redis.GeoLocation, error) {
	return c.Conn.GeoSearchLocation(c.Context, key, q).Result()
}

func (c *Client) GeoSearchStore(key, store string, q *redis.GeoSearchStoreQuery) (int64, error) {
	return c.Conn.GeoSearchStore(c.Context, key, store, q).Result()
}

func (c *Client) GeoDist(key string, member1, member2, unit string) (float64, error) {
	return c.Conn.GeoDist(c.Context, key, member1, member2, unit).Result()
}

func (c *Client) GeoHash(key string, members ...string) ([]string, error) {
	return c.Conn.GeoHash(c.Context, key, members...).Result()
}

func (c *Client) GeoPos(key string, members ...string) ([]*redis.GeoPos, error) {
	return c.Conn.GeoPos(c.Context, key, members...).Result()
}

func (c *Client) Watch(fn func(tx *redis.Tx) error, keys ...string) error {
	return c.Conn.Watch(c.Context, fn, keys...)
}

func (c *Client) Wait(numSlaves int, timeout time.Duration) (int64, error) {
	return c.Conn.Wait(c.Context, numSlaves, timeout).Result()
}

// 网络优化器, 通过 fn 函数缓冲一堆命令并一次性将它们发送到服务器执行, 好处是 节省了每个命令的网络往返时间（RTT）
func (c *Client) Pipeline(fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error) {
	pipeliner := c.Conn.Pipeline()
	fn(pipeliner)
	return pipeliner.Exec(c.Context)
}

// 事务 - 类似 Pipeline, 但是它内部会使用 MULTI/EXEC 包裹排队的命令
func (c *Client) TxPipeline(fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error) {
	txPipeline := c.Conn.TxPipeline()
	fn(txPipeline)
	return txPipeline.Exec(c.Context)
}


func (c *Client) ClientKill(ipPort string) (string, error) {
	return c.Conn.ClientKill(c.Context, ipPort).Result()
}

func (c *Client) ClientKillByFilter(keys ...string) (int64, error) {
	return c.Conn.ClientKillByFilter(c.Context, keys...).Result()
}

func (c *Client) ClientList() *redis.StringCmd {
	return c.Conn.ClientList(c.Context)
}

func (c *Client) ClientPause(dur time.Duration) (bool, error) {
	return c.Conn.ClientPause(c.Context, dur).Result()
}

func (c *Client) ClientID() (int64, error) {
	return c.Conn.ClientID(c.Context).Result()
}

func (c *Client) ClientUnblock(id int64) (int64, error) {
	return c.Conn.ClientUnblock(c.Context, id).Result()
}

func (c *Client) ClientUnblockWithError(id int64) (int64, error) {
	return c.Conn.ClientUnblockWithError(c.Context, id).Result()
}

func (c *Client) ConfigGet(parameter string) ([]interface{}, error) {
	return c.Conn.ConfigGet(c.Context, parameter).Result()
}

func (c *Client) ConfigResetStat() (string, error) {
	return c.Conn.ConfigResetStat(c.Context).Result()
}

func (c *Client) ConfigSet(parameter, value string) (string, error) {
	return c.Conn.ConfigSet(c.Context, parameter, value).Result()
}

func (c *Client) ConfigRewrite() (string, error) {
	return c.Conn.ConfigRewrite(c.Context).Result()
}

func (c *Client) DBSize() (int64, error) {
	return c.Conn.DBSize(c.Context).Result()
}

func (c *Client) FlushAll() (string, error) {
	return c.Conn.FlushAll(c.Context).Result()
}

func (c *Client) FlushAllAsync() (string, error) {
	return c.Conn.FlushAllAsync(c.Context).Result()
}

func (c *Client) FlushDB() (string, error) {
	return c.Conn.FlushDB(c.Context).Result()
}

func (c *Client) FlushDBAsync() (string, error) {
	return c.Conn.FlushDBAsync(c.Context).Result()
}

func (c *Client) Info(section ...string) *redis.StringCmd {
	return c.Conn.Info(c.Context, section...)
}

func (c *Client) LastSave() (int64, error) {
	return c.Conn.LastSave(c.Context).Result()
}

func (c *Client) Save() (string, error) {
	return c.Conn.Save(c.Context).Result()
}

func (c *Client) Shutdown() (string, error) {
	return c.Conn.Shutdown(c.Context).Result()
}

func (c *Client) ShutdownSave() (string, error) {
	return c.Conn.ShutdownSave(c.Context).Result()
}

func (c *Client) ShutdownNoSave() (string, error) {
	return c.Conn.ShutdownNoSave(c.Context).Result()
}

func (c *Client) SlaveOf(host, port string) (string, error) {
	return c.Conn.SlaveOf(c.Context, host, port).Result()
}

func (c *Client) SlowLogGet(num int64) ([]redis.SlowLog, error) {
	return c.Conn.SlowLogGet(c.Context, num).Result()
}

func (c *Client) Time() (time.Time, error) {
	return c.Conn.Time(c.Context).Result()
}

func (c *Client) DebugObject(key string) *redis.StringCmd {
	return c.Conn.DebugObject(c.Context, key)
}

func (c *Client) ReadOnly() (string, error) {
	return c.Conn.ReadOnly(c.Context).Result()
}

func (c *Client) ReadWrite() (string, error) {
	return c.Conn.ReadWrite(c.Context).Result()
}

func (c *Client) MemoryUsage(key string, samples ...int) (int64, error) {
	return c.Conn.MemoryUsage(c.Context, key, samples...).Result()
}

func (c *Client) Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	return c.Conn.Eval(c.Context, script, keys, args...)
}

func (c *Client) EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return c.Conn.EvalSha(c.Context, sha1, keys, args...)
}

func (c *Client) ScriptExists(hashes ...string) ([]bool, error) {
	return c.Conn.ScriptExists(c.Context, hashes...).Result()
}

func (c *Client) ScriptFlush() (string, error) {
	return c.Conn.ScriptFlush(c.Context).Result()
}

func (c *Client) ScriptKill() (string, error) {
	return c.Conn.ScriptKill(c.Context).Result()
}

func (c *Client) ScriptLoad(script string) *redis.StringCmd {
	return c.Conn.ScriptLoad(c.Context, script)
}

func (c *Client) Publish(channel string, message interface{}) (int64, error) {
	return c.Conn.Publish(c.Context, channel, message).Result()
}

func (c *Client) PubSubChannels(pattern string) ([]string, error) {
	return c.Conn.PubSubChannels(c.Context, pattern).Result()
}

func (c *Client) PubSubNumSub(channels ...string) (map[string]int64, error) {
	return c.Conn.PubSubNumSub(c.Context, channels...).Result()
}

func (c *Client) PubSubNumPat() (int64, error) {
	return c.Conn.PubSubNumPat(c.Context).Result()
}

func (c *Client) ClusterSlots() ([]redis.ClusterSlot, error) {
	return c.Conn.ClusterSlots(c.Context).Result()
}

func (c *Client) ClusterNodes() *redis.StringCmd {
	return c.Conn.ClusterNodes(c.Context)
}

func (c *Client) ClusterMeet(host, port string) (string, error) {
	return c.Conn.ClusterMeet(c.Context, host, port).Result()
}

func (c *Client) ClusterForget(nodeID string) (string, error) {
	return c.Conn.ClusterForget(c.Context, nodeID).Result()
}

func (c *Client) ClusterReplicate(nodeID string) (string, error) {
	return c.Conn.ClusterReplicate(c.Context, nodeID).Result()
}

func (c *Client) ClusterResetSoft() (string, error) {
	return c.Conn.ClusterResetSoft(c.Context).Result()
}

func (c *Client) ClusterResetHard() (string, error) {
	return c.Conn.ClusterResetHard(c.Context).Result()
}

func (c *Client) ClusterInfo() *redis.StringCmd {
	return c.Conn.ClusterInfo(c.Context)
}

func (c *Client) ClusterKeySlot(key string) (int64, error) {
	return c.Conn.ClusterKeySlot(c.Context, key).Result()
}

func (c *Client) ClusterGetKeysInSlot(slot int, count int) ([]string, error) {
	return c.Conn.ClusterGetKeysInSlot(c.Context, slot, count).Result()
}

func (c *Client) ClusterCountFailureReports(nodeID string) (int64, error) {
	return c.Conn.ClusterCountFailureReports(c.Context, nodeID).Result()
}

func (c *Client) ClusterCountKeysInSlot(slot int) (int64, error) {
	return c.Conn.ClusterCountKeysInSlot(c.Context, slot).Result()
}

func (c *Client) ClusterDelSlots(slots ...int) (string, error) {
	return c.Conn.ClusterDelSlots(c.Context, slots...).Result()
}

func (c *Client) ClusterDelSlotsRange(min, max int) (string, error) {
	return c.Conn.ClusterDelSlotsRange(c.Context, min, max).Result()
}

func (c *Client) ClusterSaveConfig() (string, error) {
	return c.Conn.ClusterSaveConfig(c.Context).Result()
}

func (c *Client) ClusterSlaves(nodeID string) ([]string, error) {
	return c.Conn.ClusterSlaves(c.Context, nodeID).Result()
}

func (c *Client) ClusterFailover() (string, error) {
	return c.Conn.ClusterFailover(c.Context).Result()
}

func (c *Client) ClusterAddSlots(slots ...int) (string, error) {
	return c.Conn.ClusterAddSlots(c.Context, slots...).Result()
}

func (c *Client) ClusterAddSlotsRange(min, max int) (string, error) {
	return c.Conn.ClusterAddSlotsRange(c.Context, min, max).Result()
}

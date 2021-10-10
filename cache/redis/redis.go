package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Client struct {
	Conn 	*redis.Client
	Options *Options
}

type Options struct {
	redis.Options
}

func New(options *Options) (*Client, error) {
	var client = new(Client)
	client.Options = options
	client.Conn = redis.NewClient(&options.Options)
	_, err := client.Ping()
	return client, err
}

func (c *Client) Ping() (string, error) {
	return c.Conn.Ping(context.TODO()).Result()
}

func (c *Client) Command() (map[string]*redis.CommandInfo, error) {
	return c.Conn.Command(context.TODO()).Result()
}

func (c *Client) ClientGetName() (string, error) {
	return c.Conn.ClientGetName(context.TODO()).Result()
}

func (c *Client) Echo(message interface{}) (string, error) {
	return c.Conn.Echo(context.TODO(), message).Result()
}

func (c *Client) Keys(pattern string) ([]string, error) {
	return c.Conn.Keys(context.TODO(), pattern).Result()
}

func (c *Client) Dump(key string) (string, error) {
	return c.Conn.Dump(context.TODO(), key).Result()
}

func (c *Client) Get(key string) (string, error) {
	return c.Conn.Get(context.TODO(), key).Result()
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return c.Conn.Set(context.TODO(), key, value, expiration).Result()
}

func (c *Client) SetEX(key string, value interface{}, expiration time.Duration) (string, error) {
	return c.Conn.SetEX(context.TODO(), key, value, expiration).Result()
}

func (c *Client) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.Conn.SetNX(context.TODO(), key, value, expiration).Result()
}

func (c *Client) SetXX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.Conn.SetXX(context.TODO(), key, value, expiration).Result()
}

func (c *Client) Del(keys ...string) (int64, error) {
	return c.Conn.Del(context.TODO(), keys...).Result()
}

func (c *Client) GetRange(key string, start, end int64) (string, error) {
	return c.Conn.GetRange(context.TODO(), key, start, end).Result()
}

func (c *Client) GetSet(key string, value interface{}) (string, error) {
	return c.Conn.GetSet(context.TODO(), key, value).Result()
}

func (c *Client) GetEx(key string, expiration time.Duration) (string, error) {
	return c.Conn.GetEx(context.TODO(), key, expiration).Result()
}

func (c *Client) GetDel(key string) (string, error) {
	return c.Conn.GetDel(context.TODO(), key).Result()
}

func (c *Client) StrLen(key string) (int64, error) {
	return c.Conn.StrLen(context.TODO(), key).Result()
}

func (c *Client) Incr(key string) (int64, error) {
	return c.Conn.Incr(context.TODO(), key).Result()
}

func (c *Client) Decr(key string) (int64, error) {
	return c.Conn.Decr(context.TODO(), key).Result()
}

func (c *Client) IncrBy(key string, value int64) (int64, error) {
	return c.Conn.IncrBy(context.TODO(), key, value).Result()
}

func (c *Client) DecrBy(key string, decrement int64) (int64, error) {
	return c.Conn.DecrBy(context.TODO(), key, decrement).Result()
}

func (c *Client) IncrByFloat(key string, value float64) (float64, error) {
	return c.Conn.IncrByFloat(context.TODO(), key, value).Result()
}

func (c *Client) Expire(key string, expiration time.Duration) (bool, error) {
	return c.Conn.Expire(context.TODO(), key, expiration).Result()
}

func (c *Client) ExpireAt(key string, tm time.Time) (bool, error) {
	return c.Conn.ExpireAt(context.TODO(), key, tm).Result()
}

func (c *Client) PExpire(key string, expiration time.Duration) (bool, error) {
	return c.Conn.PExpire(context.TODO(), key, expiration).Result()
}

func (c *Client) PExpireAt(key string, tm time.Time) (bool, error) {
	return c.Conn.PExpireAt(context.TODO(), key, tm).Result()
}

func (c *Client) TTL(key string) (time.Duration, error) {
	return c.Conn.TTL(context.TODO(), key).Result()
}

func (c *Client) PTTL(key string) (time.Duration, error) {
	return c.Conn.PTTL(context.TODO(), key).Result()
}

func (c *Client) Exists(keys ...string) (int64, error) {
	return c.Conn.Exists(context.TODO(), keys...).Result()
}

func (c *Client) Unlink(keys ...string) (int64, error) {
	return c.Conn.Unlink(context.TODO(), keys...).Result()
}

func (c *Client) Migrate(host, port, key string, db int, timeout time.Duration) (string, error) {
	return c.Conn.Migrate(context.TODO(), host, port, key, db, timeout).Result()
}

func (c *Client) Move(key string, db int) (bool, error) {
	return c.Conn.Move(context.TODO(), key, db).Result()
}

func (c *Client) ObjectRefCount(key string) (int64, error) {
	return c.Conn.ObjectRefCount(context.TODO(), key).Result()
}

func (c *Client) ObjectEncoding(key string) (string, error) {
	return c.Conn.ObjectEncoding(context.TODO(), key).Result()
}

func (c *Client) ObjectIdleTime(key string) (time.Duration, error) {
	return c.Conn.ObjectIdleTime(context.TODO(), key).Result()
}

func (c *Client) RandomKey() (string, error) {
	return c.Conn.RandomKey(context.TODO()).Result()
}

func (c *Client) Rename(key string, newkey string) (string, error) {
	return c.Conn.Rename(context.TODO(), key, newkey).Result()
}

func (c *Client) RenameNX(key string, newkey string) (bool, error) {
	return c.Conn.RenameNX(context.TODO(), key, newkey).Result()
}

func (c *Client) Type(key string) (string, error) {
	return c.Conn.Type(context.TODO(), key).Result()
}

func (c *Client) Append(key, value string) (int64, error) {
	return c.Conn.Append(context.TODO(), key, value).Result()
}

func (c *Client) MGet(keys ...string) ([]interface{}, error) {
	return c.Conn.MGet(context.TODO(), keys...).Result()
}

func (c *Client) MSet(values ...interface{}) (string, error) {
	return c.Conn.MSet(context.TODO(), values...).Result()
}

func (c *Client) MSetNX(values ...interface{}) (bool, error) {
	return c.Conn.MSetNX(context.TODO(), values...).Result()
}

func (c *Client) GetBit(key string, offset int64) (int64, error) {
	return c.Conn.GetBit(context.TODO(), key, offset).Result()
}

func (c *Client) SetBit(key string, offset int64, value int) (int64, error) {
	return c.Conn.SetBit(context.TODO(), key, offset, value).Result()
}

func (c *Client) BitCount(key string, bitCount *redis.BitCount) (int64, error) {
	return c.Conn.BitCount(context.TODO(), key, bitCount).Result()
}

func (c *Client) BitOpAnd(destKey string, keys ...string) (int64, error) {
	return c.Conn.BitOpAnd(context.TODO(), destKey, keys...).Result()
}

func (c *Client) BitOpOr(destKey string, keys ...string) (int64, error) {
	return c.Conn.BitOpOr(context.TODO(), destKey, keys...).Result()
}

func (c *Client) BitOpXor(destKey string, keys ...string) (int64, error) {
	return c.Conn.BitOpXor(context.TODO(), destKey, keys...).Result()
}

func (c *Client) BitOpNot(destKey string, key string) (int64, error) {
	return c.Conn.BitOpNot(context.TODO(), destKey, key).Result()
}

func (c *Client) BitPos(key string, bit int64, pos ...int64) (int64, error) {
	return c.Conn.BitPos(context.TODO(), key, bit, pos...).Result()
}

func (c *Client) BitField(key string, args ...interface{}) ([]int64, error) {
	return c.Conn.BitField(context.TODO(), key, args...).Result()
}

func (c *Client) SetArgs(key string, value interface{}, args redis.SetArgs) (string, error) {
	return c.Conn.SetArgs(context.TODO(), key, value, args).Result()
}

func (c *Client) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.Conn.Scan(context.TODO(), cursor, match, count).Result()
}

func (c *Client) ScanType(cursor uint64, match string, count int64, keyType string) ([]string, uint64, error) {
	return c.Conn.ScanType(context.TODO(), cursor, match, count, keyType).Result()
}

func (c *Client) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.Conn.SScan(context.TODO(), key, cursor, match, count).Result()
}

func (c *Client) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.Conn.HScan(context.TODO(), key, cursor, match, count).Result()
}

func (c *Client) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.Conn.ZScan(context.TODO(), key, cursor, match, count).Result()
}

func (c *Client) HDel(key string, fields ...string) (int64, error) {
	return c.Conn.HDel(context.TODO(), key, fields...).Result()
}

func (c *Client) HExists(key, field string) (bool, error) {
	return c.Conn.HExists(context.TODO(), key, field).Result()
}

func (c *Client) HGet(key, field string) (string, error) {
	return c.Conn.HGet(context.TODO(), key, field).Result()
}

func (c *Client) HGetAll(key string) (map[string]string, error) {
	return c.Conn.HGetAll(context.TODO(), key).Result()
}

func (c *Client) HIncrBy(key, field string, incr int64) (int64, error) {
	return c.Conn.HIncrBy(context.TODO(), key, field, incr).Result()
}

func (c *Client) HIncrByFloat(key, field string, incr float64) (float64, error) {
	return c.Conn.HIncrByFloat(context.TODO(), key, field, incr).Result()
}

func (c *Client) HKeys(key string) ([]string, error) {
	return c.Conn.HKeys(context.TODO(), key).Result()
}

func (c *Client) HLen(key string) (int64, error) {
	return c.Conn.HLen(context.TODO(), key).Result()
}

func (c *Client) HMGet(key string, fields ...string) ([]interface{}, error) {
	return c.Conn.HMGet(context.TODO(), key, fields...).Result()
}

func (c *Client) HSet(key string, values ...interface{}) (int64, error) {
	return c.Conn.HSet(context.TODO(), key, values...).Result()
}

func (c *Client) HMSet(key string, values ...interface{}) (bool, error) {
	return c.Conn.HMSet(context.TODO(), key, values...).Result()
}

func (c *Client) HSetNX(key, field string, value interface{}) (bool, error) {
	return c.Conn.HSetNX(context.TODO(), field, key, value).Result()
}

func (c *Client) HVals(key string) ([]string, error) {
	return c.Conn.HVals(context.TODO(), key).Result()
}

func (c *Client) HRandField(key string, count int, withValues bool) ([]string, error) {
	return c.Conn.HRandField(context.TODO(), key, count, withValues).Result()
}

func (c *Client) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return c.Conn.BLPop(context.TODO(), timeout, keys...).Result()
}

func (c *Client) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	return c.Conn.BRPop(context.TODO(), timeout, keys...).Result()
}

func (c *Client) BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
	return c.Conn.BRPopLPush(context.TODO(), source, destination, timeout).Result()
}

func (c *Client) LIndex(key string, index int64) (string, error) {
	return c.Conn.LIndex(context.TODO(), key, index).Result()
}

func (c *Client) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	return c.Conn.LInsert(context.TODO(), key, op, pivot, value).Result()
}

func (c *Client) LInsertBefore(key string, pivot, value interface{}) (int64, error) {
	return c.Conn.LInsertBefore(context.TODO(), key, pivot, value).Result()
}

func (c *Client) LInsertAfter(key string, pivot, value interface{}) (int64, error) {
	return c.Conn.LInsertAfter(context.TODO(), key, pivot, value).Result()
}

func (c *Client) LLen(key string) (int64, error) {
	return c.Conn.LLen(context.TODO(), key).Result()
}

func (c *Client) LPop(key string) (string, error) {
	return c.Conn.LPop(context.TODO(), key).Result()
}

func (c *Client) LPopCount(key string, count int) ([]string, error) {
	return c.Conn.LPopCount(context.TODO(), key, count).Result()
}

func (c *Client) LPos(key string, value string, a redis.LPosArgs) (int64, error) {
	return c.Conn.LPos(context.TODO(), key, value, a).Result()
}

func (c *Client) LPosCount(key string, value string, count int64, a redis.LPosArgs) ([]int64, error) {
	return c.Conn.LPosCount(context.TODO(), key, value, count, a).Result()
}

func (c *Client) LPush(key string, values ...interface{}) (int64, error) {
	return c.Conn.LPush(context.TODO(), key, values...).Result()
}

func (c *Client) LPushX(key string, values ...interface{}) (int64, error) {
	return c.Conn.LPushX(context.TODO(), key, values...).Result()
}

func (c *Client) LRange(key string, start, stop int64) ([]string, error) {
	return c.Conn.LRange(context.TODO(), key, start, stop).Result()
}

func (c *Client) LRem(key string, count int64, value interface{}) (int64, error) {
	return c.Conn.LRem(context.TODO(), key, count, value).Result()
}

func (c *Client) LSet(key string, index int64, value interface{}) (string, error) {
	return c.Conn.LSet(context.TODO(), key, index, value).Result()
}

func (c *Client) LTrim(key string, start int64, stop int64) (string, error) {
	return c.Conn.LTrim(context.TODO(), key, start, stop).Result()
}

func (c *Client) RPop(key string) (string, error) {
	return c.Conn.RPop(context.TODO(), key).Result()
}

func (c *Client) RPopCount(key string, count int) ([]string, error) {
	return c.Conn.RPopCount(context.TODO(), key, count).Result()
}

func (c *Client) RPopLPush(source, destination string) (string, error) {
	return c.Conn.RPopLPush(context.TODO(), source, destination).Result()
}

func (c *Client) RPush(key string, values ...interface{}) (int64, error) {
	return c.Conn.RPush(context.TODO(), key, values...).Result()
}

func (c *Client) RPushX(key string, values ...interface{}) (int64, error) {
	return c.Conn.RPushX(context.TODO(), key, values...).Result()
}

func (c *Client) LMove(source, destination, srcpos, destpos string) (string, error) {
	return c.Conn.LMove(context.TODO(), source, destination, srcpos, destpos).Result()
}

func (c *Client) BLMove(source, destination, srcpos, destpos string, timeout time.Duration) (string, error) {
	return c.Conn.BLMove(context.TODO(), source, destination, srcpos, destpos, timeout).Result()
}

func (c *Client) SAdd(key string, members ...interface{}) (int64, error) {
	return c.Conn.SAdd(context.TODO(), key, members...).Result()
}

func (c *Client) SCard(key string) (int64, error) {
	return c.Conn.SCard(context.TODO(), key).Result()
}

func (c *Client) SDiff(keys ...string) ([]string, error) {
	return c.Conn.SDiff(context.TODO(), keys...).Result()
}

func (c *Client) SDiffStore(destination string, keys ...string) (int64, error) {
	return c.Conn.SDiffStore(context.TODO(), destination, keys...).Result()
}

func (c *Client) SInter(keys ...string) ([]string, error) {
	return c.Conn.SInter(context.TODO(), keys...).Result()
}

func (c *Client) SInterStore(destination string, keys ...string) (int64, error) {
	return c.Conn.SInterStore(context.TODO(), destination, keys...).Result()
}

func (c *Client) SIsMember(key string, member interface{}) (bool, error) {
	return c.Conn.SIsMember(context.TODO(), key, member).Result()
}

func (c *Client) SMIsMember(key string, members ...interface{}) ([]bool, error) {
	return c.Conn.SMIsMember(context.TODO(), key, members...).Result()
}

func (c *Client) SMembers(key string) ([]string, error) {
	return c.Conn.SMembers(context.TODO(), key).Result()
}

func (c *Client) SMembersMap(key string) (map[string]struct{}, error) {
	return c.Conn.SMembersMap(context.TODO(), key).Result()
}

func (c *Client) SMove(source, destination string, member interface{}) (bool, error) {
	return c.Conn.SMove(context.TODO(), source, destination, member).Result()
}

func (c *Client) SPop(key string) (string, error) {
	return c.Conn.SPop(context.TODO(), key).Result()
}

func (c *Client) SPopN(key string, count int64) ([]string, error) {
	return c.Conn.SPopN(context.TODO(), key, count).Result()
}

func (c *Client) SRandMember(key string) (string, error) {
	return c.Conn.SRandMember(context.TODO(), key).Result()
}

func (c *Client) SRandMemberN(key string, count int64) ([]string, error) {
	return c.Conn.SRandMemberN(context.TODO(), key, count).Result()
}

func (c *Client) SRem(key string, members ...interface{}) (int64, error) {
	return c.Conn.SRem(context.TODO(), key, members...).Result()
}

func (c *Client) SUnion(keys ...string) ([]string, error) {
	return c.Conn.SUnion(context.TODO(), keys...).Result()
}

func (c *Client) SUnionStore(destination string, keys ...string) (int64, error) {
	return c.Conn.SUnionStore(context.TODO(), destination, keys...).Result()
}

func (c *Client) ZAddArgs(key string, args redis.ZAddArgs) (int64, error) {
	return c.Conn.ZAddArgs(context.TODO(), key, args).Result()
}

func (c *Client) ZAddArgsIncr(key string, args redis.ZAddArgs) (float64, error) {
	return c.Conn.ZAddArgsIncr(context.TODO(), key, args).Result()
}

func (c *Client) ZAdd(key string, members ...*redis.Z) (int64, error) {
	return c.Conn.ZAdd(context.TODO(), key, members...).Result()
}

func (c *Client) ZAddNX(key string, members ...*redis.Z) (int64, error) {
	return c.Conn.ZAddNX(context.TODO(), key, members...).Result()
}

func (c *Client) ZAddXX(key string, members ...*redis.Z) (int64, error) {
	return c.Conn.ZAddXX(context.TODO(), key, members...).Result()
}

func (c *Client) ZCard(key string) (int64, error) {
	return c.Conn.ZCard(context.TODO(), key).Result()
}

func (c *Client) ZCount(key, min, max string) (int64, error) {
	return c.Conn.ZCount(context.TODO(), key, min, max).Result()
}

func (c *Client) ZLexCount(key, min, max string) (int64, error) {
	return c.Conn.ZLexCount(context.TODO(), key, min, max).Result()
}

func (c *Client) ZIncrBy(key string, increment float64, member string) (float64, error) {
	return c.Conn.ZIncrBy(context.TODO(), key, increment, member).Result()
}

func (c *Client) ZInterStore(destination string, store *redis.ZStore) (int64, error) {
	return c.Conn.ZInterStore(context.TODO(), destination, store).Result()
}

func (c *Client) ZInter(store *redis.ZStore) ([]string, error) {
	return c.Conn.ZInter(context.TODO(), store).Result()
}

func (c *Client) ZInterWithScores(store *redis.ZStore) ([]redis.Z, error) {
	return c.Conn.ZInterWithScores(context.TODO(), store).Result()
}

func (c *Client) ZMScore(key string, members ...string) ([]float64, error) {
	return c.Conn.ZMScore(context.TODO(), key, members...).Result()
}

func (c *Client) ZPopMax(key string, count ...int64) ([]redis.Z, error) {
	return c.Conn.ZPopMax(context.TODO(), key, count...).Result()
}

func (c *Client) ZPopMin(key string, count ...int64) ([]redis.Z, error) {
	return c.Conn.ZPopMin(context.TODO(), key, count...).Result()
}

func (c *Client) ZRangeArgs(z redis.ZRangeArgs) ([]string, error) {
	return c.Conn.ZRangeArgs(context.TODO(), z).Result()
}

func (c *Client) ZRangeArgsWithScores(z redis.ZRangeArgs) ([]redis.Z, error) {
	return c.Conn.ZRangeArgsWithScores(context.TODO(), z).Result()
}

func (c *Client) ZRange(key string, start, stop int64) ([]string, error) {
	return c.Conn.ZRange(context.TODO(), key, start, stop).Result()
}

func (c *Client) ZRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return c.Conn.ZRangeWithScores(context.TODO(), key, start, stop).Result()
}

func (c *Client) ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.Conn.ZRangeByScore(context.TODO(), key, opt).Result()
}

func (c *Client) ZRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.Conn.ZRangeByLex(context.TODO(), key, opt).Result()
}

func (c *Client) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return c.Conn.ZRangeByScoreWithScores(context.TODO(), key, opt).Result()
}

func (c *Client) ZRangeStore(dst string, z redis.ZRangeArgs) (int64, error) {
	return c.Conn.ZRangeStore(context.TODO(), dst, z).Result()
}

func (c *Client) ZRank(key, member string) (int64, error) {
	return c.Conn.ZRank(context.TODO(), key, member).Result()
}

func (c *Client) ZRem(key string, members ...interface{}) (int64, error) {
	return c.Conn.ZRem(context.TODO(), key, members...).Result()
}

func (c *Client) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	return c.Conn.ZRemRangeByRank(context.TODO(), key, start, stop).Result()
}

func (c *Client) ZRemRangeByScore(key, min, max string) (int64, error) {
	return c.Conn.ZRemRangeByScore(context.TODO(), key, min, max).Result()
}

func (c *Client) ZRemRangeByLex(key, min, max string) (int64, error) {
	return c.Conn.ZRemRangeByLex(context.TODO(), key, min, max).Result()
}

func (c *Client) ZRevRange(key string, start, stop int64) ([]string, error) {
	return c.Conn.ZRevRange(context.TODO(), key, start, stop).Result()
}

func (c *Client) ZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return c.Conn.ZRevRangeWithScores(context.TODO(), key, start, stop).Result()
}

func (c *Client) ZRevRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.Conn.ZRevRangeByScore(context.TODO(), key, opt).Result()
}

func (c *Client) ZRevRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	return c.Conn.ZRevRangeByLex(context.TODO(), key, opt).Result()
}

func (c *Client) ZRevRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	return c.Conn.ZRevRangeByScoreWithScores(context.TODO(), key, opt).Result()
}

func (c *Client) ZRevRank(key, member string) (int64, error) {
	return c.Conn.ZRevRank(context.TODO(), key, member).Result()
}

func (c *Client) ZScore(key, member string) (float64, error) {
	return c.Conn.ZScore(context.TODO(), key, member).Result()
}

func (c *Client) ZUnion(store redis.ZStore) ([]string, error) {
	return c.Conn.ZUnion(context.TODO(), store).Result()
}

func (c *Client) ZUnionWithScores(store redis.ZStore) ([]redis.Z, error) {
	return c.Conn.ZUnionWithScores(context.TODO(), store).Result()
}

func (c *Client) ZUnionStore(dest string, store *redis.ZStore) (int64, error) {
	return c.Conn.ZUnionStore(context.TODO(), dest, store).Result()
}

func (c *Client) ZRandMember(key string, count int, withScores bool) ([]string, error) {
	return c.Conn.ZRandMember(context.TODO(), key, count, withScores).Result()
}

func (c *Client) ZDiff(keys ...string) ([]string, error) {
	return c.Conn.ZDiff(context.TODO(), keys...).Result()
}

func (c *Client) ZDiffWithScores(keys ...string) ([]redis.Z, error) {
	return c.Conn.ZDiffWithScores(context.TODO(), keys...).Result()
}

func (c *Client) ZDiffStore(destination string, keys ...string) (int64, error) {
	return c.Conn.ZDiffStore(context.TODO(), destination, keys...).Result()
}

func (c *Client) PFAdd(key string, els ...interface{}) (int64, error) {
	return c.Conn.PFAdd(context.TODO(), key, els...).Result()
}

func (c *Client) PFCount(keys ...string) (int64, error) {
	return c.Conn.PFCount(context.TODO(), keys...).Result()
}

func (c *Client) PFMerge(dest string, keys ...string) (string, error) {
	return c.Conn.PFMerge(context.TODO(), dest, keys...).Result()
}

func (c *Client) BgRewriteAOF() (string, error) {
	return c.Conn.BgRewriteAOF(context.TODO()).Result()
}

func (c *Client) BgSave() (string, error) {
	return c.Conn.BgSave(context.TODO()).Result()
}

func (c *Client) GeoAdd(key string, geoLocation ...*redis.GeoLocation) (int64, error) {
	return c.Conn.GeoAdd(context.TODO(), key, geoLocation...).Result()
}

func (c *Client) GeoRadius(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) ([]redis.GeoLocation, error) {
	return c.Conn.GeoRadius(context.TODO(), key, longitude, latitude, query).Result()
}

func (c *Client) GeoRadiusStore(key string, longitude, latitude float64, query *redis.GeoRadiusQuery) (int64, error) {
	return c.Conn.GeoRadiusStore(context.TODO(), key, longitude, latitude, query).Result()
}

func (c *Client) GeoRadiusByMember(key, member string, query *redis.GeoRadiusQuery) ([]redis.GeoLocation, error) {
	return c.Conn.GeoRadiusByMember(context.TODO(), key, member, query).Result()
}

func (c *Client) GeoRadiusByMemberStore(key, member string, query *redis.GeoRadiusQuery) (int64, error) {
	return c.Conn.GeoRadiusByMemberStore(context.TODO(), key, member, query).Result()
}

func (c *Client) GeoSearch(key string, q *redis.GeoSearchQuery) ([]string, error) {
	return c.Conn.GeoSearch(context.TODO(), key, q).Result()
}

func (c *Client) GeoSearchLocation(key string, q *redis.GeoSearchLocationQuery) ([]redis.GeoLocation, error) {
	return c.Conn.GeoSearchLocation(context.TODO(), key, q).Result()
}

func (c *Client) GeoSearchStore(key, store string, q *redis.GeoSearchStoreQuery) (int64, error) {
	return c.Conn.GeoSearchStore(context.TODO(), key, store, q).Result()
}

func (c *Client) GeoDist(key string, member1, member2, unit string) (float64, error) {
	return c.Conn.GeoDist(context.TODO(), key, member1, member2, unit).Result()
}

func (c *Client) GeoHash(key string, members ...string) ([]string, error) {
	return c.Conn.GeoHash(context.TODO(), key, members...).Result()
}

func (c *Client) GeoPos(key string, members ...string) ([]*redis.GeoPos, error) {
	return c.Conn.GeoPos(context.TODO(), key, members...).Result()
}

func (c *Client) Watch(fn func(tx *redis.Tx) error, keys ...string) error {
	return c.Conn.Watch(context.TODO(), fn, keys...)
}

func (c *Client) Wait(numSlaves int, timeout time.Duration) (int64, error) {
	return c.Conn.Wait(context.TODO(), numSlaves, timeout).Result()
}

// 网络优化器, 通过 fn 函数缓冲一堆命令并一次性将它们发送到服务器执行, 好处是 节省了每个命令的网络往返时间（RTT）
func (c *Client) Pipeline(fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error) {
	pipeliner := c.Conn.Pipeline()
	fn(pipeliner)
	return pipeliner.Exec(context.TODO())
}

// 事务 - 类似 Pipeline, 但是它内部会使用 MULTI/EXEC 包裹排队的命令
func (c *Client) TxPipeline(fn func(pipe redis.Pipeliner)) ([]redis.Cmder, error) {
	txPipeline := c.Conn.TxPipeline()
	fn(txPipeline)
	return txPipeline.Exec(context.TODO())
}


func (c *Client) ClientKill(ipPort string) (string, error) {
	return c.Conn.ClientKill(context.TODO(), ipPort).Result()
}

func (c *Client) ClientKillByFilter(keys ...string) (int64, error) {
	return c.Conn.ClientKillByFilter(context.TODO(), keys...).Result()
}

func (c *Client) ClientList() (string, error) {
	return c.Conn.ClientList(context.TODO()).Result()
}

func (c *Client) ClientPause(dur time.Duration) (bool, error) {
	return c.Conn.ClientPause(context.TODO(), dur).Result()
}

func (c *Client) ClientID() (int64, error) {
	return c.Conn.ClientID(context.TODO()).Result()
}

func (c *Client) ClientUnblock(id int64) (int64, error) {
	return c.Conn.ClientUnblock(context.TODO(), id).Result()
}

func (c *Client) ClientUnblockWithError(id int64) (int64, error) {
	return c.Conn.ClientUnblockWithError(context.TODO(), id).Result()
}

func (c *Client) ConfigGet(parameter string) ([]interface{}, error) {
	return c.Conn.ConfigGet(context.TODO(), parameter).Result()
}

func (c *Client) ConfigResetStat() (string, error) {
	return c.Conn.ConfigResetStat(context.TODO()).Result()
}

func (c *Client) ConfigSet(parameter, value string) (string, error) {
	return c.Conn.ConfigSet(context.TODO(), parameter, value).Result()
}

func (c *Client) ConfigRewrite() (string, error) {
	return c.Conn.ConfigRewrite(context.TODO()).Result()
}

func (c *Client) DBSize() (int64, error) {
	return c.Conn.DBSize(context.TODO()).Result()
}

func (c *Client) FlushAll() (string, error) {
	return c.Conn.FlushAll(context.TODO()).Result()
}

func (c *Client) FlushAllAsync() (string, error) {
	return c.Conn.FlushAllAsync(context.TODO()).Result()
}

func (c *Client) FlushDB() (string, error) {
	return c.Conn.FlushDB(context.TODO()).Result()
}

func (c *Client) FlushDBAsync() (string, error) {
	return c.Conn.FlushDBAsync(context.TODO()).Result()
}

func (c *Client) Info(section ...string) (string, error) {
	return c.Conn.Info(context.TODO(), section...).Result()
}

func (c *Client) LastSave() (int64, error) {
	return c.Conn.LastSave(context.TODO()).Result()
}

func (c *Client) Save() (string, error) {
	return c.Conn.Save(context.TODO()).Result()
}

func (c *Client) Shutdown() (string, error) {
	return c.Conn.Shutdown(context.TODO()).Result()
}

func (c *Client) ShutdownSave() (string, error) {
	return c.Conn.ShutdownSave(context.TODO()).Result()
}

func (c *Client) ShutdownNoSave() (string, error) {
	return c.Conn.ShutdownNoSave(context.TODO()).Result()
}

func (c *Client) SlaveOf(host, port string) (string, error) {
	return c.Conn.SlaveOf(context.TODO(), host, port).Result()
}

func (c *Client) SlowLogGet(num int64) ([]redis.SlowLog, error) {
	return c.Conn.SlowLogGet(context.TODO(), num).Result()
}

func (c *Client) Time() (time.Time, error) {
	return c.Conn.Time(context.TODO()).Result()
}

func (c *Client) DebugObject(key string) (string, error) {
	return c.Conn.DebugObject(context.TODO(), key).Result()
}

func (c *Client) ReadOnly() (string, error) {
	return c.Conn.ReadOnly(context.TODO()).Result()
}

func (c *Client) ReadWrite() (string, error) {
	return c.Conn.ReadWrite(context.TODO()).Result()
}

func (c *Client) MemoryUsage(key string, samples ...int) (int64, error) {
	return c.Conn.MemoryUsage(context.TODO(), key, samples...).Result()
}

func (c *Client) Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	return c.Conn.Eval(context.TODO(), script, keys, args...).Result()
}

func (c *Client) EvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return c.Conn.EvalSha(context.TODO(), sha1, keys, args...).Result()
}

func (c *Client) ScriptExists(hashes ...string) ([]bool, error) {
	return c.Conn.ScriptExists(context.TODO(), hashes...).Result()
}

func (c *Client) ScriptFlush() (string, error) {
	return c.Conn.ScriptFlush(context.TODO()).Result()
}

func (c *Client) ScriptKill() (string, error) {
	return c.Conn.ScriptKill(context.TODO()).Result()
}

func (c *Client) ScriptLoad(script string) (string, error) {
	return c.Conn.ScriptLoad(context.TODO(), script).Result()
}

func (c *Client) Publish(channel string, message interface{}) (int64, error) {
	return c.Conn.Publish(context.TODO(), channel, message).Result()
}

func (c *Client) PubSubChannels(pattern string) ([]string, error) {
	return c.Conn.PubSubChannels(context.TODO(), pattern).Result()
}

func (c *Client) PubSubNumSub(channels ...string) (map[string]int64, error) {
	return c.Conn.PubSubNumSub(context.TODO(), channels...).Result()
}

func (c *Client) PubSubNumPat() (int64, error) {
	return c.Conn.PubSubNumPat(context.TODO()).Result()
}

func (c *Client) ClusterSlots() ([]redis.ClusterSlot, error) {
	return c.Conn.ClusterSlots(context.TODO()).Result()
}

func (c *Client) ClusterNodes() (string, error) {
	return c.Conn.ClusterNodes(context.TODO()).Result()
}

func (c *Client) ClusterMeet(host, port string) (string, error) {
	return c.Conn.ClusterMeet(context.TODO(), host, port).Result()
}

func (c *Client) ClusterForget(nodeID string) (string, error) {
	return c.Conn.ClusterForget(context.TODO(), nodeID).Result()
}

func (c *Client) ClusterReplicate(nodeID string) (string, error) {
	return c.Conn.ClusterReplicate(context.TODO(), nodeID).Result()
}

func (c *Client) ClusterResetSoft() (string, error) {
	return c.Conn.ClusterResetSoft(context.TODO()).Result()
}

func (c *Client) ClusterResetHard() (string, error) {
	return c.Conn.ClusterResetHard(context.TODO()).Result()
}

func (c *Client) ClusterInfo() (string, error) {
	return c.Conn.ClusterInfo(context.TODO()).Result()
}

func (c *Client) ClusterKeySlot(key string) (int64, error) {
	return c.Conn.ClusterKeySlot(context.TODO(), key).Result()
}

func (c *Client) ClusterGetKeysInSlot(slot int, count int) ([]string, error) {
	return c.Conn.ClusterGetKeysInSlot(context.TODO(), slot, count).Result()
}

func (c *Client) ClusterCountFailureReports(nodeID string) (int64, error) {
	return c.Conn.ClusterCountFailureReports(context.TODO(), nodeID).Result()
}

func (c *Client) ClusterCountKeysInSlot(slot int) (int64, error) {
	return c.Conn.ClusterCountKeysInSlot(context.TODO(), slot).Result()
}

func (c *Client) ClusterDelSlots(slots ...int) (string, error) {
	return c.Conn.ClusterDelSlots(context.TODO(), slots...).Result()
}

func (c *Client) ClusterDelSlotsRange(min, max int) (string, error) {
	return c.Conn.ClusterDelSlotsRange(context.TODO(), min, max).Result()
}

func (c *Client) ClusterSaveConfig() (string, error) {
	return c.Conn.ClusterSaveConfig(context.TODO()).Result()
}

func (c *Client) ClusterSlaves(nodeID string) ([]string, error) {
	return c.Conn.ClusterSlaves(context.TODO(), nodeID).Result()
}

func (c *Client) ClusterFailover() (string, error) {
	return c.Conn.ClusterFailover(context.TODO()).Result()
}

func (c *Client) ClusterAddSlots(slots ...int) (string, error) {
	return c.Conn.ClusterAddSlots(context.TODO(), slots...).Result()
}

func (c *Client) ClusterAddSlotsRange(min, max int) (string, error) {
	return c.Conn.ClusterAddSlotsRange(context.TODO(), min, max).Result()
}

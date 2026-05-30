package redis

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

// setupMiniRedis 创建一个 miniredis 实例用于测试
func setupMiniRedis(t *testing.T) *miniredis.Miniredis {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("failed to create miniredis: %v", err)
	}
	return mr
}

// testClient 使用 miniredis 创建测试客户端
func testClient(t *testing.T) Client {
	mr := setupMiniRedis(t)
	t.Cleanup(func() { mr.Close() })

	opts := &Options{
		Options: redis.Options{
			Network: "tcp",
			Addr:    mr.Addr(),
		},
	}
	client, err := NewClient(context.Background(), opts)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client
}

func TestNewClient(t *testing.T) {
	client := testClient(t)

	// 测试 Ping
	ctx := context.Background()
	result := client.Ping(ctx)
	if result.Err() != nil {
		t.Errorf("Ping failed: %v", result.Err())
	}
}

func TestSetAndGet(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// 测试 Set
	setResult := client.Set(ctx, "test_key", "test_value", time.Hour)
	if setResult.Err() != nil {
		t.Errorf("Set failed: %v", setResult.Err())
	}

	// 测试 Get
	getResult := client.Get(ctx, "test_key")
	if getResult.Err() != nil {
		t.Errorf("Get failed: %v", getResult.Err())
	}
	if getResult.Val() != "test_value" {
		t.Errorf("Get value mismatch: expected 'test_value', got '%s'", getResult.Val())
	}
}

func TestKeys(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// 设置一些键
	client.Set(ctx, "key1", "value1", time.Hour)
	client.Set(ctx, "key2", "value2", time.Hour)

	// 测试 Keys
	result := client.Keys(ctx, "*")
	if result.Err() != nil {
		t.Errorf("Keys failed: %v", result.Err())
	}
	if len(result.Val()) < 2 {
		t.Errorf("Keys should return at least 2 keys, got %d", len(result.Val()))
	}
}

func TestDel(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// 设置键
	client.Set(ctx, "del_key", "del_value", time.Hour)

	// 测试 Del
	result := client.Del(ctx, "del_key")
	if result.Err() != nil {
		t.Errorf("Del failed: %v", result.Err())
	}

	// 确认已删除
	getResult := client.Get(ctx, "del_key")
	if getResult.Err() == nil {
		t.Errorf("Key should be deleted")
	}
}

func TestSetNX(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// 测试 SetNX - 键不存在时应该成功
	result := client.SetNX(ctx, "nx_key", "nx_value", time.Hour)
	if result.Err() != nil {
		t.Errorf("SetNX failed: %v", result.Err())
	}
	if !result.Val() {
		t.Errorf("SetNX should return true for new key")
	}

	// 再次 SetNX - 键已存在时应该失败
	result2 := client.SetNX(ctx, "nx_key", "nx_value2", time.Hour)
	if result2.Val() {
		t.Errorf("SetNX should return false for existing key")
	}
}

func TestIncr(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// 测试 Incr
	client.Set(ctx, "counter", "10", time.Hour)
	result := client.Incr(ctx, "counter")
	if result.Err() != nil {
		t.Errorf("Incr failed: %v", result.Err())
	}
	if result.Val() != 11 {
		t.Errorf("Incr value mismatch: expected 11, got %d", result.Val())
	}
}

func TestExpire(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// 设置键
	client.Set(ctx, "expire_key", "expire_value", time.Hour)

	// 测试 Expire
	result := client.Expire(ctx, "expire_key", time.Minute*10)
	if result.Err() != nil {
		t.Errorf("Expire failed: %v", result.Err())
	}
	if !result.Val() {
		t.Errorf("Expire should return true")
	}
}

func TestTTL(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// 设置键并设置过期时间
	client.Set(ctx, "ttl_key", "ttl_value", time.Hour)

	// 测试 TTL
	result := client.TTL(ctx, "ttl_key")
	if result.Err() != nil {
		t.Errorf("TTL failed: %v", result.Err())
	}
}

func TestExists(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// 设置键
	client.Set(ctx, "exists_key", "exists_value", time.Hour)

	// 测试 Exists - 存在
	result := client.Exists(ctx, "exists_key")
	if result.Err() != nil {
		t.Errorf("Exists failed: %v", result.Err())
	}
	if result.Val() != 1 {
		t.Errorf("Exists should return 1 for existing key")
	}

	// 测试 Exists - 不存在
	result2 := client.Exists(ctx, "non_exists_key")
	if result2.Val() != 0 {
		t.Errorf("Exists should return 0 for non-existing key")
	}
}

func TestPipeline(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	// 测试 Pipeline
	cmds, err := client.TxPipeline(ctx, func(pipe redis.Pipeliner) {
		pipe.Set(ctx, "pipe_key", "pipe_value", time.Hour)
		pipe.Get(ctx, "pipe_key")
	})
	if err != nil {
		t.Errorf("Pipeline failed: %v", err)
	}
	t.Logf("Pipeline commands: %v", cmds)
}

func TestClose(t *testing.T) {
	client := testClient(t)

	// 测试 Close
	err := client.Close()
	if err != nil {
		t.Errorf("Close failed: %v", err)
	}
}

func TestRaw(t *testing.T) {
	client := testClient(t)

	// 测试 Raw 返回原生客户端
	raw := client.Raw()
	if raw == nil {
		t.Errorf("Raw should return non-nil client")
	}
}
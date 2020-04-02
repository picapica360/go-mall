// Package redis for cache provider
package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"

	"go-mall/lib/cache"
)

var (
	// DefaultKey the collection name of redis for cache adapter.
	DefaultKey = "mall_redis"
)

// Cache is redis cache adapter.
type Cache struct {
	p        *redis.Pool
	conninfo string
	dbNum    int
	key      string
	password string
	maxIdle  int // the pool has a limit of connections that are idle.
}

func init() {
	cache.Register("redis", NewRedisCache)
}

// NewRedisCache create a new redis cache with default collection name.
func NewRedisCache() cache.Cache {
	return &Cache{key: DefaultKey}
}

// args[0] must be the key name
func (c *Cache) do(cmd string, args ...interface{}) (interface{}, error) {
	if len(args) < 1 {
		return nil, errors.New("miss required arguments")
	}
	args[0] = c.associate(args[0])
	conn := c.p.Get()
	defer conn.Close()

	return conn.Do(cmd, args...)
}

// format key, like 'prifix:key'
func (c *Cache) associate(originKey interface{}) string {
	return fmt.Sprintf("%s:%s", c.key, originKey)
}

// Get get cache from redis
func (c *Cache) Get(key string) interface{} {
	if v, err := c.do("GET", key); err == nil {
		return v
	}
	return nil
}

// GetMulti get cache from redis
func (c *Cache) GetMulti(keys []string) []interface{} {
	conn := c.p.Get()
	defer conn.Close()

	var args []interface{}
	for _, key := range keys {
		args = append(args, c.associate(key))
	}
	values, err := redis.Values(conn.Do("MGET", args...))
	if err != nil {
		return nil
	}

	return values
}

// Put put cache to redis
func (c *Cache) Put(key string, val interface{}, timeout time.Duration) error {
	_, err := c.do("SETEX", key, int64(timeout/time.Second), val)
	return err
}

// Delete delete cache in redis
func (c *Cache) Delete(key string) error {
	_, err := c.do("DEL", key)
	return err
}

// Increment increment counter in redis.
func (c *Cache) Increment(key string) error {
	_, err := redis.Bool(c.do("INCRBY", key, 1))
	return err
}

// Decrement decrement counter in redis.
func (c *Cache) Decrement(key string) error {
	_, err := redis.Bool(c.do("INCRBY", key, -1))
	return err
}

// IsExist check whether cache in redis
func (c *Cache) IsExist(key string) bool {
	v, err := redis.Bool(c.do("EXISTS", key))
	if err != nil {
		return false
	}

	return v
}

// ClearAll clean all cache in redis
func (c *Cache) ClearAll() error {
	conn := c.p.Get()
	defer conn.Close()

	keys, err := redis.Strings(c.do("KEYS", c.key+":*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		if _, err := conn.Do("DEL", key); err != nil {
			return err
		}
	}

	return err
}

// StartAndGC start gc routine based on config string settings.
// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
// the cache item in redis are stored forever, so no gc operation.
func (c *Cache) StartAndGC(config string) error {
	var conf map[string]string
	json.Unmarshal([]byte(config), &conf)

	if _, ok := conf["key"]; !ok {
		conf["key"] = DefaultKey
	}
	if _, ok := conf["conn"]; !ok {
		return errors.New("config has no conn key")
	}

	// Format redis://<password>@<host>:<port>
	conf["conn"] = strings.Replace(conf["conn"], "redis://", "", 1)
	if i := strings.Index(conf["conn"], "@"); i > -1 {
		conf["password"] = conf["conn"][0:i]
		conf["conn"] = conf["conn"][i+1:]
	}

	if _, ok := conf["dbNum"]; !ok {
		conf["dbNum"] = "0"
	}
	if _, ok := conf["password"]; !ok {
		conf["password"] = ""
	}
	if _, ok := conf["maxIdle"]; !ok {
		conf["maxIdle"] = "3"
	}

	c.key = conf["key"]
	c.conninfo = conf["conn"]
	c.dbNum, _ = strconv.Atoi(conf["dbNum"])
	c.password = conf["password"]
	c.maxIdle, _ = strconv.Atoi(conf["maxIdle"])

	c.connectInit()

	// verify the redis whether could connect to the redis server.
	conn := c.p.Get()
	defer conn.Close()

	return conn.Err()
}

// connect to redis.
func (c *Cache) connectInit() {
	dialFunc := func() (conn redis.Conn, err error) {
		conn, err = redis.Dial("tcp", c.conninfo)
		if err != nil {
			return nil, err
		}

		if len(c.password) != 0 {
			if _, err = conn.Do("AUTH", c.password); err != nil {
				conn.Close()
				return nil, err
			}
		}

		_, err = conn.Do("SELECT", c.dbNum)
		if err != nil {
			conn.Close()
			return nil, err
		}

		return
	}

	c.p = &redis.Pool{
		MaxIdle:     c.maxIdle,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}

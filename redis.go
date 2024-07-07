package rds

import (
	"context"
	"github.com/archine/ioc"
	"github.com/archine/redis-starter/conf"
	"github.com/redis/go-redis/v9"
	"github.com/spaolacci/murmur3"
	"github.com/spf13/viper"
	"log"
	"sort"
)

var ConfigHook func(opt *redis.Options) // 配置钩子，用于修改配置

type Redis struct {
	clients []*redis.Client
}

func (r *Redis) CreateBean() ioc.Bean {
	v := ioc.GetBeanByName("viper.Viper")
	if v == nil {
		log.Fatalf("faild to create bean redis, the config reader is nil")
	}
	reader := v.(*viper.Viper)
	opts := make([]*conf.Options, 0)
	if err := reader.UnmarshalKey("redis", &opts); err != nil {
		log.Fatalf("Error create bean with name redis, %s", err.Error())
	}
	sort.Slice(opts, func(i, j int) bool {
		return opts[i].Addr > opts[j].Addr
	})

	var instance Redis

	for _, opt := range opts {
		options, err := conf.ConvertToOfficialOptions(opt)
		if err != nil {
			log.Fatalf("Error create bean with name redis, %s", err.Error())
		}
		if ConfigHook != nil {
			ConfigHook(options)
		}
		redisConn := redis.NewClient(options)
		ping := redisConn.Ping(context.Background())
		if ping.Err() != nil {
			log.Fatalf("faild to create bean redis, connect error: %s", ping.Err().Error())
		}
		instance.clients = append(instance.clients, redisConn)
	}

	ConfigHook = nil
	return &instance
}

// GetClient 获取客户端
func (r *Redis) GetClient() *redis.Client {
	return r.clients[0]
}

// GetClientByHash 获取redis客户端
func (r *Redis) GetClientByHash(key string) *redis.Client {
	h32 := murmur3.New32()
	_, _ = h32.Write([]byte(key))
	hashVal := h32.Sum32()
	redisInx := int(hashVal) % len(r.clients)
	return r.clients[redisInx]
}

package conf

import (
	"github.com/archine/gin-plus/v3/beans"
	"github.com/redis/go-redis/v9"
	"time"
)

// Options Redis 配置选项。
type Options struct {
	// Redis 服务器的地址，格式为 "host:port"。
	Addr string `mapstructure:"addr"`

	// 使用的用户名
	Username string `mapstructure:"username"`

	// 连接密码
	Password string `mapstructure:"password"`

	// 数据库编号。默认值为 0。
	DB int `mapstructure:"db"`

	// 放弃前的最大重试次数。默认是 3 次重试；-1 表示禁用重试。
	MaxRetries int `mapstructure:"max_retries"`

	// 建立新连接的拨号超时时间。默认是 5 秒。
	DialTimeout time.Duration `mapstructure:"dial_timeout"`

	// 读取操作的超时时间。支持的值：0 - 默认超时（3 秒），-1 - 无超时（无限期阻塞），-2 - 禁用 SetReadDeadline 调用。
	ReadTimeout time.Duration `mapstructure:"read_timeout"`

	// 写入操作的超时时间。
	// 支持的值：0 - 默认超时（3 秒），-1 - 无超时（无限期阻塞），-2 - 禁用 SetWriteDeadline 调用。
	WriteTimeout time.Duration `mapstructure:"write_timeout"`

	// 连接池的类型。
	// true 表示 FIFO（先进先出）池，false 表示 LIFO（后进先出）池。
	PoolFIFO bool `mapstructure:"pool_fifo"`

	// 连接池的基本连接数。默认值为每个可用 CPU 10 个连接。
	PoolSize int `mapstructure:"pool_size"`

	// 如果所有连接都忙，客户端等待连接的时间。默认是 ReadTimeout + 1 秒。
	PoolTimeout time.Duration `mapstructure:"pool_timeout"`

	// 连接池中保持的最小空闲连接数。默认值为 0，表示没有限制。
	MinIdleConns int `mapstructure:"min_idle_conns"`

	// 连接池中保持的最大空闲连接数。默认值为 0，表示没有限制。
	MaxIdleConns int `mapstructure:"max_idle_conns"`

	// 连接池中分配的最大连接数。当为 0 时，连接池中连接的数量没有限制。
	MaxActiveConns int `mapstructure:"max_active_conns"`

	// 连接的最大空闲时间，应小于服务器的超时。默认是 30 分钟；-1 禁用空闲超时检查。
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`

	// 连接的最大重用时间。默认是无限制。
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

// ConvertToOfficialOptions 转换为官方的配置选项。
func ConvertToOfficialOptions(myOpts *Options) (*redis.Options, error) {
	var opts redis.Options
	err := beans.CopyProperties(myOpts, &opts)
	return &opts, err
}

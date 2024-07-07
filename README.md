![](https://img.shields.io/badge/version-v1.0.0-green.svg) &nbsp; ![](https://img.shields.io/badge/version-go1.21-green.svg) &nbsp;  ![](https://img.shields.io/badge/builder-success-green.svg) &nbsp;

> Redis 自动装配，只需简单配置即可使用

## 一、安装

### 1. Go get

```shell
go get github.com/archine/redis-starter@v1.0.0
```

### 2、Go Mod

```shell
# Add the following line to your go.mod file
github.com/archine/redis-starter v1.0.0

# Execute the following command in your project directory
go mod tidy
```

## 二、使用

### 1. 配置

| 配置项                | 	描述                                                                 |
|--------------------|---------------------------------------------------------------------|
| addr               | 	Redis 服务器的地址，格式为 "host"。                                           |
| username           | 	使用的用户名。                                                            |
| password           | 	连接密码。                                                              |
| db                 | 	数据库编号。默认值为 0。                                                      |
| max_retries        | 	放弃前的最大重试次数。默认是 3 次重试；-1 表示禁用重试。                                    |       
| dial_timeout       | 	建立新连接的拨号超时时间。默认是 5 秒。                                              |
| read_timeout       | 	读取操作的超时时间。支持的值：0 - 默认超时（3 秒），-1 - 无超时，-2 - 禁用 SetReadDeadline 调用。  | 
| write_timeout      | 	写入操作的超时时间。支持的值：0 - 默认超时（3 秒），-1 - 无超时，-2 - 禁用 SetWriteDeadline 调用。 |
| pool_fifo          | 	连接池的类型。true 表示 FIFO（先进先出）池，false 表示 LIFO（后进先出）池。                   |     
| pool_size          | 	连接池的基本连接数。默认值为每个可用 CPU 10 个连接。                                     |
| pool_timeout       | 	如果所有连接都忙，客户端等待连接的时间。默认是 ReadTimeout + 1 秒。                         |   
| min_idle_conns     | 	连接池中保持的最小空闲连接数。默认值为 0。                                             |          
| max_idle_conns     | 	连接池中保持的最大空闲连接数。默认值为 0。                                             |              
| max_active_conns   | 	连接池中分配的最大连接数。当为 0 时，连接池中连接的数量没有限制。                                 |                        
| conn_max_idle_time | 	连接的最大空闲时间。默认是 30 分钟；-1 禁用空闲超时检查。                                   |                     
| conn_max_lifetime  | 	连接的最大重用时间。默认是无限制。                                                  |            

**示例**
    
```yaml
redis:
  - addr: "localhost:6379"
    username:
    password:
    db: 0
    max_retries: 3
    dial_timeout: 5s
    read_timeout: 3s
    write_timeout: 3s
    pool_fifo: true
    pool_size: 10
    pool_timeout: 4s
    min_idle_conns: 0
    max_idle_conns: 0
    max_active_conns: 0
    conn_max_idle_time: 30m
    conn_max_lifetime: 0
```

### 2. 使用案例

```go
package model

import (
    "context"
    "github.com/archine/ioc"
    starter "github.com/archine/redis-starter"
    "github.com/redis/go-redis/v9"
)

type RedisMapper struct {
    *starter.Redis // 注入 Redis Bean
}

func (r *RedisMapper) CreateBean() ioc.Bean {
    return &RedisMapper{}
}

// 示例方法，用于与 Redis 交互
func (r *RedisMapper) SetValue(key, value string) error {
    err := r.Redis.GetClient().Set(context.Background(), "", "").Err()
    if err != nil {
        return err
    }
    return nil
}

```


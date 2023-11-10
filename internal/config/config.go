package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	MySQLConf struct {
		DataSource string
	}
	RedisConf redis.RedisConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}

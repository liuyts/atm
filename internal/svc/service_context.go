package svc

import (
	"ATM/internal/config"
	"ATM/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	Tx          sqlx.Session
	UserModel   model.UserModel
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Tx:          sqlx.NewMysql(c.MySQLConf.DataSource),
		UserModel:   model.NewUserModel(sqlx.NewMysql(c.MySQLConf.DataSource)),
		RedisClient: redis.MustNewRedis(c.RedisConf),
	}
}

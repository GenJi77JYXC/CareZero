package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"www.genji.xin/backend/CareZero/model"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseCfg model.DbConfig
	Rds         redis.RedisConf
}

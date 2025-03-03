package svc

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	"www.genji.xin/backend/CareZero/authServer/authservice"
	"www.genji.xin/backend/CareZero/model"
	"www.genji.xin/backend/CareZero/productServer/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Rds     *redis.Redis
	AuthRpc authservice.AuthService
	Auth    *casbin.Enforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DatabaseCfg.User, c.DatabaseCfg.PassWord, c.DatabaseCfg.Host, c.DatabaseCfg.Port, c.DatabaseCfg.Schema)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable:       true, // 使用单数表名，即表名称后面没有s
			IdentifierMaxLength: 64,
		},
		DryRun:                                   false, // 当=true时，生成 SQL 但不执行，可以用于准备或测试生成的 SQL
		PrepareStmt:                              true,  // PreparedStmt 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
		AllowGlobalUpdate:                        false, // 不允许在 没有指定任何条件 的情况下对整个表的数据执行更新或删除操作
		DisableAutomaticPing:                     false, // 在完成初始化后，GORM 会自动 ping 数据库以检查数据库的可用性, 若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: false, // 在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.User{}, &model.Product{})
	if err != nil {
		return nil
	}
	fmt.Println("productServer --  mysql connect success, currentDatabase is ", db.Migrator().CurrentDatabase())
	// 初始化Redis
	rdsConf := redis.RedisConf{
		Host: c.Rds.Host,
		Type: c.Rds.Type,
		//Pass:        c.Redis.Pass,
		//Tls:         c.Redis.Tls,
		//NonBlock:    c.Redis.NonBlock,
		PingTimeout: time.Second,
	}
	rds := redis.MustNewRedis(rdsConf)
	fmt.Println("productServer -- redis 连接成功")

	enforcer, err := casbin.NewEnforcer("C:\\Users\\16129\\go\\src\\www.genji.xin\\backend\\CareZero\\try\\model.conf", "C:\\Users\\16129\\go\\src\\www.genji.xin\\backend\\CareZero\\try\\policy.csv")
	if err != nil {
		panic(err)
	}

	//初始化casbin 的适配器
	//adapter, err := gormadapter.NewAdapterByDB(db)
	//if err != nil {
	//	panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	//}
	//加载模型，生成casbin执行器
	//enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	//if err != nil {
	//	panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	//}
	// 从数据库加载policy
	//err = enforcer.LoadPolicy()
	//if err != nil {
	//	log.Fatalf("failed to load policy: %v", err)
	//	return nil
	//}

	fmt.Println("productServer -- 鉴权组件Casbin注册成功")
	fmt.Println("productServer配置初始化成功")

	return &ServiceContext{
		Config:  c,
		DB:      db,
		Rds:     rds,
		AuthRpc: authservice.NewAuthService(zrpc.MustNewClient(c.AuthRpcConf)),
		Auth:    enforcer,
	}
}

package servers

import (
	"bito_group/config"
	"fmt"

	"bito_group/internal/user"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"

	//  mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbDriver = "mysql"
	dbURLFmt = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

type Repos struct {
	UserRepo user.UserRepo
}

func NewDB(cfg *config.SugaredConfig) *gorm.DB {
	dbURL := fmt.Sprintf(dbURLFmt, cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Database)

	db, err := gorm.Open(dbDriver, dbURL)
	if err != nil {
		panic(fmt.Sprintf("gorm open dbURL=%v, err=%v", dbURL, err))
	}

	db.LogMode(cfg.Log.Env == "dev")

	return db
}

func NewCache(cfg *config.SugaredConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       0,
	})
}

func NewRepos(cfg *config.SugaredConfig) *Repos {
	// 持久化类型的 repo
	// db := NewDB(cfg)
	// userRepo := user.NewMysqlUserRepo(db)

	userRepo := user.NewMysqlUserRepo(nil)
	// 如果有其他操作db的, 可以一起夾帶出去

	return &Repos{
		UserRepo: userRepo,
	}
}

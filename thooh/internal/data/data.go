package data

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"

	"thooh/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewReadRepo, NewCategoryRepo, NewArticleRepo)

// Data .
type Data struct {
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger klog.Logger) (*Data, func(), error) {
	log := klog.NewHelper(logger)
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		rdb: rdb,
	}
	return d, func() {
		log.Info("closing the data resources")
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

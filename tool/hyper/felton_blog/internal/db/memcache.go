package db

import (
	"github.com/griffin702/ginana/library/cache/memcache"
	"github.com/griffin702/ginana/library/conf/paladin"

	"hyper-go/tool\hyper/felton_blog/internal/config"
)

func NewMC(cfg *config.Config) (mc memcache.Memcache, err error) {
	key := "memcache.toml"
	if err = paladin.Get(key).UnmarshalTOML(cfg); err != nil {
		return
	}
	mc = memcache.New(cfg.Memcache)
	return
}

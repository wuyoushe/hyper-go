package service

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/griffin702/ginana/library/cache/memcache"
	"github.com/griffin702/ginana/library/database"
	"github.com/griffin702/service/tools"
	"github.com/jinzhu/gorm"

	"hyper-go/tool\hyper/felton_blog/internal/config"
)

type Service interface {
	Close()
	SetEnforcer(ef *casbin.SyncedEnforcer) (err error)
	GetEFRoles(ctx context.Context) (roles []*database.EFRolePolicy, err error)
	GetEFUsers(ctx context.Context) (users []*database.EFUseRole, err error)
}

func New(cfg *config.Config, db *gorm.DB, mc memcache.Memcache, hm HelperMap) (s Service, err error) {
	s = &service{
		cfg:  cfg,
		db:   db,
		mc:   mc,
		hm:   hm,
		tool: tools.Tools,
	}
	return
}

type service struct {
	cfg  *config.Config
	db   *gorm.DB
	ef   *casbin.SyncedEnforcer
	mc   memcache.Memcache
	hm   HelperMap
	tool *tools.Tool
}

func (s *service) Close() {
	_ = s.db.Close()
}

// Close close the resource.
func (s *service) SetEnforcer(ef *casbin.SyncedEnforcer) (err error) {
	if !s.cfg.Casbin.Enable {
		return
	}
	if s.tool.PtrIsNil(ef) {
		return fmt.Errorf("enforcer is nil")
	}
	s.ef = ef
	return
}

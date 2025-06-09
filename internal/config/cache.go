package config

import (
	"time"

	"github.com/spf13/viper"
)

var cacheCfg *CacheCfg

// CacheCfg holds list of cache keys ...
type CacheCfg struct {
	Prefix     string
	TTLDefault time.Duration
}

// loadCacheCfg loads Keys for cache configuration
func loadCacheCfg() {
	cacheCfg = &CacheCfg{
		Prefix:     viper.GetString("cache.prefix"),
		TTLDefault: viper.GetDuration("cache.ttl.default") * time.Second,
	}
}

// Cache returns the cache keys configuration
func Cache() *CacheCfg {
	return cacheCfg
}

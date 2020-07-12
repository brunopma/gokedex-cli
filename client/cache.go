package client

import "time"

//Settings is the struct that implements cache settings options
type Settings struct {
	CustomExpire time.Duration
	UseCache     bool
}

//defaultSettings are settings not meant to be changed.
type defaultSettings struct {
	MaxExpire time.Duration
	MinExpire time.Duration
}

//CacheSettings are user settings for cache expiration.
var CacheSettings = Settings{
	CustomExpire: 0,
	UseCache:     true,
}

var defaultCacheSettings = defaultSettings{
	MaxExpire: 10 * time.Minute,
	MinExpire: 5 * time.Minute,
}

// ClearCache clears all cached data.
func ClearCache() {
	c.Flush()
}

// setCache adds new item to local cache.
func setCache(endpoint string, body []byte) {
	if CacheSettings.CustomExpire != 0 {
		c.Set(endpoint, body, CacheSettings.CustomExpire*time.Minute)
	} else {
		c.SetDefault(endpoint, body)
	}
}

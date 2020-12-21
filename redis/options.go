package redis

import (
	"time"

	r "github.com/go-redis/redis/v8"
)

// OptionsFunc represents an configuration function for Redis.
type OptionsFunc func(*r.UniversalOptions)

// WithDB configures the Redis db num.
func WithDB(db int) OptionsFunc {
	return func(o *r.UniversalOptions) {
		o.DB = db
	}
}

// WithUsername configures the Redis username.
func WithUsername(username string) OptionsFunc {
	return func(o *r.UniversalOptions) {
		o.Username = username
	}
}

// WithPassword configures the Redis password.
func WithPassword(password string) OptionsFunc {
	return func(o *r.UniversalOptions) {
		o.Password = password
	}
}

// WithMasterName configures the Redis Master Name.
func WithMasterName(masterName string) OptionsFunc {
	return func(o *r.UniversalOptions) {
		o.MasterName = masterName
	}
}

// WithPoolSize configures the Redis pool size.
func WithPoolSize(size int) OptionsFunc {
	return func(o *r.UniversalOptions) {
		o.PoolSize = size
	}
}

// WithPoolTimeout configures the Redis pool timeout.
func WithPoolTimeout(timeout time.Duration) OptionsFunc {
	return func(o *r.UniversalOptions) {
		o.PoolTimeout = timeout
	}
}

// WithDialTimeout configures the Redis dail timeout.
func WithDialTimeout(timeout time.Duration) OptionsFunc {
	return func(o *r.UniversalOptions) {
		o.DialTimeout = timeout
	}
}

// WithReadTimeout configures the Redis read timeout.
func WithReadTimeout(timeout time.Duration) OptionsFunc {
	return func(o *r.UniversalOptions) {
		o.ReadTimeout = timeout
	}
}

// WithWriteTimeout configures the Redis write timeout.
func WithWriteTimeout(timeout time.Duration) OptionsFunc {
	return func(o *r.UniversalOptions) {
		o.WriteTimeout = timeout
	}
}

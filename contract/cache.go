package contract

import "time"

type Cache interface {
	Get(route string) (data interface{}, err error)
	Set(route string, data interface{}, expiration time.Duration) (err error)
	// Invalidate(route string) (err error)
}

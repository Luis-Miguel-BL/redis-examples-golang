package contract

import "time"

type Lock interface {
	Get(key string) (hasLock bool)
	Set(key string, lockExpiration time.Duration) (err error)
}

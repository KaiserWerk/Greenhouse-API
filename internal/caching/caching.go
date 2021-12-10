package caching

import "time"

var (
	lastAccess time.Time
)

func SetLastAccess() {
	lastAccess = time.Now()
}
func GetLastAccess() time.Time {
	return lastAccess
}

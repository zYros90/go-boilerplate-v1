package data

import "fmt"

const (
	cacheKeyPrefix = "cache"
)

func usrKey(usrName string) string {
	return fmt.Sprintf("%s:%s", cacheKeyPrefix, usrName)
}

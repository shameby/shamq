package global

import (
	"sync"
)

var (
	ChanMappingMutex *sync.Mutex
)

func init() {
	ChanMappingMutex = new(sync.Mutex)
}
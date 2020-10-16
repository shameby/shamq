package global

import (
	"sync"

	"shamq/types"
)

type ChanMap struct {
	L *sync.Mutex
	M map[string]chan *types.Msg
}

var (
	ChanMapping ChanMap
)

func init() {
	ChanMapping = ChanMap{
		L: new(sync.Mutex),
		M: make(map[string]chan *types.Msg, 1000),
	}
}

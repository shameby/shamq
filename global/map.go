package global

import (
	"shamq/types"
)

var (
	ChanMapping map[string]chan *types.Msg
)

func init() {
	ChanMapping = make(map[string]chan *types.Msg)
}

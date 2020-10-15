package conn

import (
	"net"

	"shamq/types"
	"shamq/global"
)

func handleReq(c *net.TCPConn, req *types.TcpReq) {
	global.ChanMappingMutex.Lock()
	_, exist := global.ChanMapping[req.Q]
	if !exist {
		global.ChanMapping[req.Q] = make(chan *types.Msg, 1000)
	}
	global.ChanMappingMutex.Unlock()

	switch req.ReqType {
	case types.Pub:
		handlePubReq(c, req)
	case types.Listen:
		handleListenReq(c, req)
	}
}

func handlePubReq(c *net.TCPConn, req *types.TcpReq) {
	global.ChanMapping[req.Q] <- req.Msg

	c.Write(types.SuccessResp("send msg done").ToByte())
}

func handleListenReq(c *net.TCPConn, req *types.TcpReq) {
	c.Write(types.SuccessResp("listing on " + req.Q).ToByte())
	for msg := range global.ChanMapping[req.Q] {
		go c.Write(types.SuccessResp(msg.Data).ToByte())
	}
}

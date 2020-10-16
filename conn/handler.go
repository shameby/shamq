package conn

import (
	"net"

	"shamq/types"
	"shamq/global"
)

func handleReq(c *net.TCPConn, req *types.TcpReq) {
	syncQueueInit(req.Q)
	switch req.ReqType {
	case types.Pub:
		handlePubReq(c, req)
	case types.Listen:
		handleListenReq(c, req)
	}
}

func syncQueueInit(q string) {
	global.ChanMapping.L.Lock()
	defer global.ChanMapping.L.Unlock()
	_, exist := global.ChanMapping.M[q]
	if !exist {
		global.ChanMapping.M[q] = make(chan *types.Msg, 1000)
	}
}

func handlePubReq(c *net.TCPConn, req *types.TcpReq) {
	global.ChanMapping.M[req.Q] <- req.Msg

	c.Write(types.SuccessResp("send msg done").ToByte())
}

func handleListenReq(c *net.TCPConn, req *types.TcpReq) {
	c.Write(types.SuccessResp("listing on " + req.Q).ToByte())
	for msg := range global.ChanMapping.M[req.Q] {
		go c.Write(types.SuccessResp(msg.Data).ToByte())
	}
}

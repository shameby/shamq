package conn

import (
	"io"
	"net"
	"bufio"
	"encoding/json"

	"shamq/logger"
	"shamq/config"
	"shamq/types"
)

func Listen() {
	var (
		tcpAddr     *net.TCPAddr
		tcpListener *net.TCPListener
	)

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:"+config.Conf.Common.TcpPort)
	tcpListener, _ = net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()

	logger.Info("Server ready to conn, Listening on " + config.Conf.Common.TcpPort)

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		logger.Info("A client connected :" + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(c *net.TCPConn) {
	defer func() {
		logger.Error(" Disconnected : " + c.RemoteAddr().String())
		c.Close()
	}()

	reader := bufio.NewReader(c)
	for {
		var (
			req    string
			err    error
			reqObj = &types.TcpReq{}
		)
		if req, err = reader.ReadString('\n'); err != nil || err == io.EOF {
			break
		}

		req = req[:len(req)-1]
		logger.Info(req)

		if err = json.Unmarshal([]byte(req), &reqObj); err != nil {
			c.Write(types.StandardResp{types.ParamsValidateErr, types.ParamsValidateErr.Error(), req}.ToByte())
			break
		}

		go handleReq(c, reqObj)
	}
}

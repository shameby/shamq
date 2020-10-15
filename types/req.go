package types

const (
	Pub    int8 = iota + 1
	Listen
)

type TcpReq struct {
	ReqType int8   `json:"req_type"`
	Q       string `json:"q"`
	Msg     *Msg   `json:"msg"`
}

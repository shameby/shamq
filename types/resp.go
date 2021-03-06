package types

import "encoding/json"

type StandardResp struct {
	Code ErrCode `json:"code"`
	Msg  string  `json:"msg"`
	Data string  `json:"data"`
}

func SuccessResp(data string) *StandardResp {
	return &StandardResp{
		Code: NoErr,
		Msg:  NoErr.Error(),
		Data: data,
	}
}

func (s StandardResp) ToByte() []byte {
	var (
		b []byte
	)

	b, _ = json.Marshal(s)

	return []byte(string(b) + "\n")
}

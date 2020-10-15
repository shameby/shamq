package types

type ErrCode int64

const (
	NoErr             ErrCode = iota
	ParamsValidateErr
)

var errMapping = map[ErrCode]string{
	NoErr:             "success",
	ParamsValidateErr: "params validate err",
}

func (e ErrCode) Error() string {
	return errMapping[e]
}

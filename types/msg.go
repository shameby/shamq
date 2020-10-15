package types

type Msg struct {
	Data string `json:"data"`
	Opt  Option `json:"opt"`
}

type Option struct {
	IsSync    bool `json:"is_sync"`
	IsPersist bool `json:"is_persist"`
}

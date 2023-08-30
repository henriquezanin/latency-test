package models

type Error struct {
	InternalErr      error
	InternalErrLevel int
	HttpErr          string
	HttpStatusCode   int
}

type Params struct {
	ServiceChainURL string
}

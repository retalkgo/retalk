package entity

type Server struct {
	Base
	IsInit bool
	ApiKey string
	Lang   string
}

type CookedServer struct {
	IsInit bool
	Lang   string
}

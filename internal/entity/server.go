package entity

type Server struct {
	Base
	IsInit bool
	ApiKey string
}

type CookedServer struct {
	IsInit bool
}

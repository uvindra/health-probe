package common

type Service interface {
	Start(port int)
	Stop()
}

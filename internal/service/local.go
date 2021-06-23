package service

type LocalService interface {
	Invite()
	Respond()
	Close()
}

type localService struct {
	LocalService
}

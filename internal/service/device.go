package service

type DeviceService interface {
	Add()
	Send()
	Remove()
	Close()
}

type deviceService struct {
	DeviceService
}

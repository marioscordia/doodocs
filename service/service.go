package service


type Service struct {
	File
}

func NewService() *Service {
	return &Service{
		NewFileService(),
	}
}


package services

var (
	ItemService ItemsServiceInterface = &itemService{}
)

type itemService struct {}

type ItemsServiceInterface interface {
	CetakAngka()
	CetakString()
}

func (s *itemService) CetakAngka() {

}

func (s *itemService) CetakString()  {

}
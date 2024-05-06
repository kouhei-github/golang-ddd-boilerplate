package services

type IHelloWorldRepository interface {
	GetHello(id int) (string, error)
	GetWorld() string
}

type HelloWorldService struct {
	repositoryHelloWorld IHelloWorldRepository
}

func NewHelloWorldService(
	repositoryHelloWorld IHelloWorldRepository,
) *HelloWorldService {
	return &HelloWorldService{
		repositoryHelloWorld: repositoryHelloWorld,
	}
}

func (c *HelloWorldService) HelloMessage(tid int) string {
	msg, _ := c.repositoryHelloWorld.GetHello(tid)
	return msg
}

func (c *HelloWorldService) WorldMessage(id int, userID int) string {
	return c.repositoryHelloWorld.GetWorld()
}

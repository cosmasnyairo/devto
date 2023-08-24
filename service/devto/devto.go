package devtoservice

import "context"

type devtoMethods interface {
	GetArticles(ctx context.Context, baseurl string, username string) ([]Article, error)
}

type devtoService struct {
	devtoMethods devtoMethods
}

type devto struct {
	Username string    `json:"username"`
	Articles []Article `json:"articles"`
}

func NewdevtoService(devtoMethods devtoMethods) *devtoService {
	return &devtoService{
		devtoMethods: devtoMethods,
	}
}

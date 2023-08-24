package devto

import (
	"context"
	"fmt"

	"github.com/cosmasnyairo/devto/internal/devto"
	devtoservice "github.com/cosmasnyairo/devto/service/devto"
)

func Run() error {

	devto := devtoservice.NewdevtoService(&devto.DevtoRequestBody{})
	articles, err := devto.GetArticles(context.Background(), "https://dev.to/api/articles?username=", "cosmasnyairo")
	if err != nil {
		return err
	}
	fmt.Println(articles)
	return nil
}

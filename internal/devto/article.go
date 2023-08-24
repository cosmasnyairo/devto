package devto

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	devtoservice "github.com/cosmasnyairo/devto/service/devto"
)

type ArticleRequestBody struct {
}

func (d *DevtoRequestBody) GetArticles(ctx context.Context, baseurl string, username string) ([]devtoservice.Article, error) {
	url := fmt.Sprintf("%v%v", baseurl, username)
	req, err := http.NewRequest("GET", url, nil)
	articles := []devtoservice.Article{}
	if err != nil {
		return []devtoservice.Article{}, err
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return []devtoservice.Article{}, err
	}
	defer resp.Body.Close()

	response, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal(response, &articles); err != nil {
		return []devtoservice.Article{}, err
	}

	return articles, nil

}

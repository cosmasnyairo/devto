package devtoservice

import "context"

var EmptyArticle = Article{}

type Article struct {
	Id                     int      `json:"id,omitempty"`
	Title                  string   `json:"title,omitempty"`
	Url                    string   `json:"url,omitempty"`
	Published_timestamp    string   `json:"published_timestamp,omitempty"`
	Public_reactions_count int      `json:"public_reactions_count,omitempty"`
	Social_image           string   `json:"social_image,omitempty"`
	Tag_list               []string `json:"tag_list,omitempty"`
}

func (m *devtoService) GetArticles(ctx context.Context, baseurl string, username string) ([]Article, error) {
	// Log here
	articles, err := m.devtoMethods.GetArticles(ctx, baseurl, username)
	if err != nil {
		return []Article{EmptyArticle}, err
	}
	return articles, nil
}

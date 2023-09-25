package dto

import (
	"myblog.backend/model"
	"time"
)

type ArticleResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Title        string    `json:"title"`
	Desc         string    `json:"desc"`
	Content      string    `json:"content"`
	Img          string    `json:"img"`
	CommentCount int       `json:"comment_count"`
	ReadCount    int       `json:"read_count"`
	Category     Category  `json:"category"`
	Author       Author    `json:"author"`
}

type Author struct {
	UserID   uint   `json:"user_id"`
	FullName string `json:"full_name"`
}

type Category struct {
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func ArticleToResponse(article *model.Article) ArticleResponse {
	return ArticleResponse{
		ID:           article.ID,
		CreatedAt:    article.CreatedAt,
		UpdatedAt:    article.UpdatedAt,
		Title:        article.Title,
		Desc:         article.Desc,
		Content:      article.Content,
		Img:          article.Img,
		CommentCount: article.CommentCount,
		ReadCount:    article.ReadCount,
		Category: Category{
			CategoryID:   article.CategoryID,
			CategoryName: article.Category.Name,
		},
		Author: Author{
			UserID:   article.UserID,
			FullName: article.User.FullName,
		},
	}
}

func ArticleSliceToResponse(articles []model.Article) []ArticleResponse {
	var responses []ArticleResponse
	for _, article := range articles {
		response := ArticleToResponse(&article)
		responses = append(responses, response)
	}
	return responses
}

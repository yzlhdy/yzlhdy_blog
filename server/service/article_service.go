package service

import (
	"yzlhdy_blog/dto"
	"yzlhdy_blog/entity"
	"yzlhdy_blog/repository"

	"github.com/mashingan/smapping"
)

type ArticleService interface {
	CreateArticle(article dto.ArticleCreateDto) entity.Article
	DeleteArticle(articleId int) entity.Article
	UpdateArticle(id int, article dto.ArticleUpdateDto) entity.Article
	AllArticle(page int, limit int) ([]entity.Article, int64)
}

type articleService struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) ArticleService {
	return &articleService{articleRepo: articleRepo}
}

// 新增文章
func (s *articleService) CreateArticle(article dto.ArticleCreateDto) entity.Article {
	articleDto := entity.Article{}
	err := smapping.FillStruct(&articleDto, smapping.MapFields(&article))
	if err != nil {
		panic(err)
	}
	result := s.articleRepo.CreateArticle(articleDto)
	return result
}

// 删除文章
func (s *articleService) DeleteArticle(articleId int) entity.Article {
	return s.articleRepo.DeleteArticle(articleId)
}

// 更新文章
func (s *articleService) UpdateArticle(id int, article dto.ArticleUpdateDto) entity.Article {

	articleDto := entity.Article{}
	err := smapping.FillStruct(&articleDto, smapping.MapFields(&article))
	if err != nil {
		panic(err)
	}
	result := s.articleRepo.UpdateArticle(id, articleDto)
	return result
}

// 查询所有文章
func (s *articleService) AllArticle(page int, limit int) ([]entity.Article, int64) {
	return s.articleRepo.AllArticle(page, limit)
}

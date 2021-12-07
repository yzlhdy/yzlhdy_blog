package repository

import (
	"yzlhdy_blog/entity"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article entity.Article) entity.Article
	DeleteArticle(articleId int) entity.Article
	UpdateArticle(id int, article entity.Article) entity.Article
	AllArticle(page int, limit int) ([]entity.Article, int64)
}

type articleRepository struct {
	connection *gorm.DB
}

func NewArticleRepository(connection *gorm.DB) ArticleRepository {
	return &articleRepository{
		connection: connection,
	}
}

func (r *articleRepository) CreateArticle(article entity.Article) entity.Article {
	r.connection.Create(&article)
	return article
}

func (r *articleRepository) DeleteArticle(articleId int) entity.Article {
	var article entity.Article
	r.connection.First(&article, articleId)
	r.connection.Delete(&article)
	return article
}

func (r *articleRepository) UpdateArticle(id int, article entity.Article) entity.Article {
	articleData := entity.Article{}
	r.connection.Model(&article).Where("id = ?", id).Updates(article)
	return articleData

}

func (r *articleRepository) AllArticle(page int, limit int) ([]entity.Article, int64) {
	var articles []entity.Article
	var count int64
	r.connection.Preload("Classification").Limit(limit).Offset((page - 1) * limit).Find(&articles)
	r.connection.Model(&entity.Article{}).Count(&count)
	return articles, count
}

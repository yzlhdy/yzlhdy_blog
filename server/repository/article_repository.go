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
	FndArticle(id int) entity.Article
}

type articleRepository struct {
	connection *gorm.DB
}

// 实例化
func NewArticleRepository(connection *gorm.DB) ArticleRepository {
	return &articleRepository{
		connection: connection,
	}
}

// 创建文章
func (r *articleRepository) CreateArticle(article entity.Article) entity.Article {
	r.connection.Create(&article)
	return article
}

// 删除文章
func (r *articleRepository) DeleteArticle(articleId int) entity.Article {
	var article entity.Article
	r.connection.First(&article, articleId)
	r.connection.Delete(&article)
	return article
}

// 更新文章
func (r *articleRepository) UpdateArticle(id int, article entity.Article) entity.Article {
	articleData := entity.Article{}
	r.connection.Model(&article).Where("id = ?", id).Updates(article)
	return articleData

}

// 分页查询
func (r *articleRepository) AllArticle(page int, limit int) ([]entity.Article, int64) {
	var articles []entity.Article
	var count int64
	r.connection.Preload("Classification").Offset((page - 1) * limit).Limit(limit).Find(&articles).Count(&count)
	return articles, count
}

// 查询文章
func (r *articleRepository) FndArticle(id int) entity.Article {
	var article entity.Article
	r.connection.First(&article, id)
	return article
}

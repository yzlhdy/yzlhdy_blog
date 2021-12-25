package repository

import (
	"yzlhdy_blog/entity"

	"gorm.io/gorm"
)

type RecategoryRepository interface {
	GetAll(page int, limit int) ([]entity.ResourceCategory, int64)
	GetByID(id int) entity.ResourceCategory
	Create(recategory entity.ResourceCategory) entity.ResourceCategory
	Update(id int, recategory entity.ResourceCategory) entity.ResourceCategory
	Delete(id int) entity.ResourceCategory
}

type recategoryRepository struct {
	connection *gorm.DB
}

func NewRecategoryRepository(connection *gorm.DB) RecategoryRepository {
	return &recategoryRepository{
		connection: connection,
	}
}

func (db *recategoryRepository) GetAll(page int, limit int) ([]entity.ResourceCategory, int64) {
	var recategory []entity.ResourceCategory
	var total int64
	db.connection.Offset((page - 1) * limit).Limit(limit).Find(&recategory).Count(&total)
	return recategory, total

}
func (db *recategoryRepository) GetByID(id int) entity.ResourceCategory {
	var recategory entity.ResourceCategory
	db.connection.Where("id = ?", id).First(&recategory)
	return recategory
}
func (db *recategoryRepository) Create(recategory entity.ResourceCategory) entity.ResourceCategory {
	db.connection.Save(&recategory)
	return recategory
}
func (db *recategoryRepository) Update(id int, recategory entity.ResourceCategory) entity.ResourceCategory {
	db.connection.Model(&recategory).Where("id = ?", id).Updates(recategory)
	return recategory
}
func (db *recategoryRepository) Delete(id int) entity.ResourceCategory {
	var recategory entity.ResourceCategory
	db.connection.Where("id = ?", id).Delete(&recategory)
	return recategory
}

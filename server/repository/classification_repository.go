package repository

import (
	"fmt"
	"yzlhdy_blog/entity"

	"gorm.io/gorm"
)

type classificationRepository struct {
	connection *gorm.DB
}

func NewClassificationRepository(connection *gorm.DB) ClassificationRepository {
	return &classificationRepository{
		connection: connection,
	}
}

type ClassificationRepository interface {
	InsertClassification(imitation entity.Classification) entity.Classification
	FindClassificationById(id int) entity.Classification
	UpdateClassification(id int, imitation entity.Classification) entity.Classification
	DeleteClassification(id int) entity.Classification
	AllClassification(page int, limit int) ([]entity.Classification, int64)
}

func (db *classificationRepository) InsertClassification(imitation entity.Classification) entity.Classification {
	db.connection.Save(&imitation)
	return imitation
}

func (db *classificationRepository) FindClassificationById(id int) entity.Classification {
	classification := entity.Classification{}
	db.connection.Where("id = ?", id).First(&classification)
	return classification
}

func (db *classificationRepository) UpdateClassification(id int, imitation entity.Classification) entity.Classification {
	fiction := entity.Classification{}
	db.connection.Model(&fiction).Where("id = ?", id).Updates(imitation)
	return fiction
}

func (db *classificationRepository) DeleteClassification(id int) entity.Classification {
	fiction := entity.Classification{}
	db.connection.Where("id = ?", id).Delete(&fiction)
	return fiction
}

func (db *classificationRepository) AllClassification(page int, limit int) ([]entity.Classification, int64) {
	classifications := []entity.Classification{}
	var total int64
	db.connection.Offset((page - 1) * limit).Limit(limit).Find(&classifications).Count(&total)
	fmt.Println(total)

	return classifications, total
}

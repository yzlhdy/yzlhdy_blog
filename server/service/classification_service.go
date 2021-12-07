package service

import (
	"fmt"
	"log"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/entity"
	"yzlhdy_blog/repository"

	"github.com/mashingan/smapping"
)

type ClassificationService interface {
	InsertClassification(imitation dto.CreateClassifDto) entity.Classification
	FindClassificationById(id int) entity.Classification
	UpdateClassification(id int, imitation dto.UpdateClassifDto) entity.Classification
	DeleteClassification(id int) entity.Classification
	AllClassification(page int, limit int) ([]entity.Classification, int64)
}

type classificationService struct {
	classRep repository.ClassificationRepository
}

func NewClassificationService(classRep repository.ClassificationRepository) ClassificationService {
	return &classificationService{
		classRep: classRep,
	}
}

func (service *classificationService) InsertClassification(imitation dto.CreateClassifDto) entity.Classification {
	classDto := entity.Classification{}
	err := smapping.FillStruct(&classDto, smapping.MapFields(&imitation))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	classCreate := service.classRep.InsertClassification(classDto)
	return classCreate

}

func (service *classificationService) FindClassificationById(id int) entity.Classification {
	return service.classRep.FindClassificationById(id)
}

func (service *classificationService) UpdateClassification(id int, imitation dto.UpdateClassifDto) entity.Classification {

	classify := entity.Classification{}
	err := smapping.FillStruct(&classify, smapping.MapFields(&imitation))
	if err != nil {
		fmt.Println(err)
	}
	return service.classRep.UpdateClassification(id, classify)
}

func (service *classificationService) DeleteClassification(id int) entity.Classification {
	return service.classRep.DeleteClassification(id)
}

func (service *classificationService) AllClassification(page int, limit int) ([]entity.Classification, int64) {
	return service.classRep.AllClassification(page, limit)
}

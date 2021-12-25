package service

import (
	"log"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/entity"
	"yzlhdy_blog/repository"

	"github.com/mashingan/smapping"
)

type RecategoryService interface {
	GetAll(page int, limit int) ([]entity.ResourceCategory, int64)
	GetByID(id int) entity.ResourceCategory
	Create(recategory dto.CreateReCategory) entity.ResourceCategory
	Update(id int, recategory dto.UpdateReCategory) entity.ResourceCategory
	Delete(id int) entity.ResourceCategory
}

type recategoryService struct {
	repository repository.RecategoryRepository
}

func NewRecategoryService(repo repository.RecategoryRepository) RecategoryService {
	return &recategoryService{
		repository: repo,
	}
}

func (re *recategoryService) GetAll(page int, limit int) ([]entity.ResourceCategory, int64) {
	return re.repository.GetAll(page, limit)
}
func (re *recategoryService) GetByID(id int) entity.ResourceCategory {
	return re.repository.GetByID(id)
}

func (re *recategoryService) Create(recategory dto.CreateReCategory) entity.ResourceCategory {
	recategoryCreate := entity.ResourceCategory{}
	err := smapping.FillStruct(&recategoryCreate, smapping.MapFields(&recategory))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	return re.repository.Create(recategoryCreate)
}

func (re *recategoryService) Update(id int, recategory dto.UpdateReCategory) entity.ResourceCategory {
	recategoryCreate := entity.ResourceCategory{}
	err := smapping.FillStruct(&recategoryCreate, smapping.MapFields(&recategory))
	if err != nil {
		panic(err)
	}
	return re.repository.Update(id, recategoryCreate)
}

func (re *recategoryService) Delete(id int) entity.ResourceCategory {
	return re.repository.Delete(id)
}

package service

import (
	"log"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/entity"
	"yzlhdy_blog/repository"

	"github.com/mashingan/smapping"
)

type ResourcesService interface {
	GetResources(page int, limit int) ([]entity.Resources, int64)
	GetResource(id int) entity.Resources
	CreateResource(resource dto.CreateResourcesDto) entity.Resources
	UpdateResource(id int, resource dto.UpdateResourcesDto) entity.Resources
	DeleteResource(id int) entity.Resources
}

type resourcesService struct {
	resourcesRepository repository.ResourcesRepository
}

func NewResourcesService(resourcesRepository repository.ResourcesRepository) ResourcesService {
	return &resourcesService{resourcesRepository: resourcesRepository}
}

// GetResources 获取资源列表
func (r *resourcesService) GetResources(page int, limit int) ([]entity.Resources, int64) {
	return r.resourcesRepository.GetResources(page, limit)
}

// GetResource 获取资源
func (r *resourcesService) GetResource(id int) entity.Resources {
	return r.resourcesRepository.GetResource(id)
}

// CreateResource 创建资源
func (r *resourcesService) CreateResource(resource dto.CreateResourcesDto) entity.Resources {
	var resources entity.Resources
	err := smapping.FillStruct(&resources, smapping.MapFields(&resource))
	if err != nil {
		log.Fatal(err)
	}
	return r.resourcesRepository.CreateResource(resources)
}

// UpdateResource 更新资源
func (r *resourcesService) UpdateResource(id int, resource dto.UpdateResourcesDto) entity.Resources {
	var resources entity.Resources
	err := smapping.FillStruct(&resources, smapping.MapFields(&resource))
	if err != nil {
		log.Fatal(err)
	}
	return r.resourcesRepository.UpdateResource(id, resources)
}

// DeleteResource 删除资源
func (r *resourcesService) DeleteResource(id int) entity.Resources {
	return r.resourcesRepository.DeleteResource(id)
}

package repository

import (
	"yzlhdy_blog/entity"

	"gorm.io/gorm"
)

type ResourcesRepository interface {
	GetResources(page int, limit int) ([]entity.Resources, int64)
	GetResource(id int) entity.Resources
	CreateResource(resource entity.Resources) entity.Resources
	UpdateResource(id int, resource entity.Resources) entity.Resources
	DeleteResource(id int) entity.Resources
}

type resourcesRepository struct {
	db *gorm.DB
}

func NewResourcesRepository(db *gorm.DB) ResourcesRepository {
	return &resourcesRepository{db: db}
}

// GetResources 获取资源列表
func (r *resourcesRepository) GetResources(page int, limit int) ([]entity.Resources, int64) {
	var resources []entity.Resources
	var count int64
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	r.db.Model(&entity.Resources{}).Count(&count)
	r.db.Limit(limit).Offset((page - 1) * limit).Find(&resources)
	return resources, count
}

// GetResource 获取资源
func (r *resourcesRepository) GetResource(id int) entity.Resources {
	var resource entity.Resources
	r.db.Where("id = ?", id).First(&resource)
	return resource
}

// CreateResource 创建资源
func (r *resourcesRepository) CreateResource(resource entity.Resources) entity.Resources {
	r.db.Create(&resource)
	return resource
}

// UpdateResource 更新资源
func (r *resourcesRepository) UpdateResource(id int, resource entity.Resources) entity.Resources {
	r.db.Model(&entity.Resources{}).Where("id = ?", id).Updates(resource)
	return resource
}

// DeleteResource 删除资源
func (r *resourcesRepository) DeleteResource(id int) entity.Resources {
	var resource entity.Resources
	r.db.Where("id = ?", id).Delete(&resource)
	return resource
}

package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/helper"
	"yzlhdy_blog/service"

	"github.com/gin-gonic/gin"
)

type ResourcesController interface {
	GetResources(ctx *gin.Context)
	GetResource(ctx *gin.Context)
	CreateResource(ctx *gin.Context)
	UpdateResource(ctx *gin.Context)
	DeleteResource(ctx *gin.Context)
}

type resourcesController struct {
	resourcesService service.ResourcesService
}

// NewResourcesController 创建资源控制器
func NewResourcesController(resourcesService service.ResourcesService) ResourcesController {
	return &resourcesController{resourcesService: resourcesService}
}

// GetResources 获取资源列表
func (s *resourcesController) GetResources(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	resources, total := s.resourcesService.GetResources(page, limit)
	response := helper.BuildResponseAll(200, "success", resources, total, nil)
	ctx.JSON(http.StatusOK, response)
}

// GetResource 获取资源
func (s *resourcesController) GetResource(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result := s.resourcesService.GetResource(id)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

// CreateResource 创建资源
func (s *resourcesController) CreateResource(ctx *gin.Context) {
	var createResources dto.CreateResourcesDto
	validate, err := helper.BindAndValid(ctx, &createResources)
	fmt.Println(createResources)
	if !validate {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusOK, response)
		return
	}
	result := s.resourcesService.CreateResource(createResources)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

// UpdateResource 更新资源
func (s *resourcesController) UpdateResource(ctx *gin.Context) {
	var updateResources dto.UpdateResourcesDto
	id, _ := strconv.Atoi(ctx.Param("id"))
	valiedate, err := helper.BindAndValid(ctx, &updateResources)
	if !valiedate {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}
	result := s.resourcesService.UpdateResource(id, updateResources)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)

}

// DeleteResource 删除资源
func (s *resourcesController) DeleteResource(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err.Error())
		ctx.JSON(http.StatusOK, response)
		return
	}
	result := s.resourcesService.DeleteResource(id)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

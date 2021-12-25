package controller

import (
	"net/http"
	"strconv"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/helper"
	"yzlhdy_blog/service"

	"github.com/gin-gonic/gin"
)

type ResourceController interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type resourceController struct {
	service service.RecategoryService
}

func NewResourceController(service service.RecategoryService) ResourceController {
	return &resourceController{
		service: service,
	}
}

// get All
func (se *resourceController) GetAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	result, total := se.service.GetAll(page, limit)
	response := helper.BuildResponseAll(200, "success", result, total, nil)
	ctx.JSON(http.StatusOK, response)

}

// get by id
func (se *resourceController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	result := se.service.GetByID(id)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

func (se *resourceController) Create(ctx *gin.Context) {
	var resourDto dto.CreateReCategory
	validate, err := helper.BindAndValid(ctx, &resourDto)
	if !validate {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusOK, response)
		return
	}

	result := se.service.Create(resourDto)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

func (se *resourceController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse(400, "Id是必传的", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusOK, response)
		return
	}
	result := se.service.Delete(id)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

func (se *resourceController) Update(ctx *gin.Context) {
	var resourDto dto.UpdateReCategory
	id, _ := strconv.Atoi(ctx.Param("id"))
	validate, err := helper.BindAndValid(ctx, &resourDto)
	if !validate {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusOK, response)
		return
	}
	result := se.service.Update(id, resourDto)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

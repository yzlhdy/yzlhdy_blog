package controller

import (
	"net/http"
	"strconv"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/helper"
	"yzlhdy_blog/service"

	"github.com/gin-gonic/gin"
)

type ClassificationController interface {
	InsertClassification(ctx *gin.Context)
	FindClassificationById(ctx *gin.Context)
	UpdateClassification(ctx *gin.Context)
	DeleteClassification(ctx *gin.Context)
	AllClassification(ctx *gin.Context)
}

type classificationController struct {
	classiService service.ClassificationService
}

func NewClassificationController(classiService service.ClassificationService) ClassificationController {
	return &classificationController{
		classiService: classiService,
	}
}

// @Summary 添加分类
func (c *classificationController) InsertClassification(ctx *gin.Context) {
	var createDto dto.CreateClassifDto
	ctx.ShouldBind(&createDto)
	valide, err := helper.BindAndValid(ctx, &createDto)
	if !valide {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(response.Status, response)
		return
	}
	result := c.classiService.InsertClassification(createDto)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(response.Status, response)
}

// @Summary 根据id查询分类
func (c *classificationController) FindClassificationById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(response.Status, response)
		return
	}
	result := c.classiService.FindClassificationById(id)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(response.Status, response)
}

// @Summary 更新分类
func (c *classificationController) UpdateClassification(ctx *gin.Context) {
	var updateDto dto.UpdateClassifDto
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(response.Status, response)
		return

	}
	ctx.ShouldBind(&updateDto)
	validate, err := helper.BindAndValid(ctx, &updateDto)
	if !validate {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	result := c.classiService.UpdateClassification(id, updateDto)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(http.StatusOK, response)
}

// @Summary 删除分类
func (c *classificationController) DeleteClassification(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(response.Status, response)
		return
	}
	result := c.classiService.DeleteClassification(id)
	response := helper.BuildResponse(200, "success", result)
	ctx.JSON(response.Status, response)
}

// @Summary 查询所有分类
func (c *classificationController) AllClassification(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	result, total := c.classiService.AllClassification(page, limit)
	response := helper.BuildResponseAll(200, "success", result, total, nil)
	ctx.JSON(http.StatusOK, response)
}

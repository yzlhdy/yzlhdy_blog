package controller

import (
	"net/http"
	"strconv"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/helper"
	"yzlhdy_blog/service"

	"github.com/gin-gonic/gin"
)

type ArticleController interface {
	CreateArticle(ctx *gin.Context)
	DeleteArticle(ctx *gin.Context)
	UpdateArticle(ctx *gin.Context)
	AllArticle(ctx *gin.Context)
}

type articleController struct {
	articleService service.ArticleService
}

func NewArticleController(articleService service.ArticleService) ArticleController {
	return &articleController{articleService: articleService}
}

func (c *articleController) CreateArticle(ctx *gin.Context) {
	var articleDto dto.ArticleCreateDto
	ctx.ShouldBind(&articleDto)
	validate, err := helper.BindAndValid(ctx, &articleDto)
	if !validate {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := c.articleService.CreateArticle(articleDto)
	response := helper.BuildResponse(200, "success", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *articleController) DeleteArticle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := c.articleService.DeleteArticle(id)
	response := helper.BuildResponse(200, "success", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *articleController) UpdateArticle(ctx *gin.Context) {
	var updateArticleDto dto.ArticleUpdateDto
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	ctx.ShouldBind(&updateArticleDto)
	validate, err := helper.BindAndValid(ctx, &updateArticleDto)
	if !validate {
		response := helper.BuildErrorResponse(400, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := c.articleService.UpdateArticle(id, updateArticleDto)
	response := helper.BuildResponse(200, "success", res)
	ctx.JSON(http.StatusOK, response)

}

func (c *articleController) AllArticle(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Param("page"))
	limit, _ := strconv.Atoi(ctx.Param("limit"))
	res, total := c.articleService.AllArticle(page, limit)
	response := helper.BuildResponseAll(200, "success", res, total, nil)
	ctx.JSON(http.StatusOK, response)

}

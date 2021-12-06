package controller

import (
	"net/http"
	"strconv"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/helper"
	"yzlhdy_blog/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JwtService
}

func NewUserController(userService service.UserService, jwtService service.JwtService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) FindUser(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	res := c.userService.FindUser(page, limit)
	response := helper.BuildResponse(200, "success", res)
	ctx.JSON(http.StatusOK, response)

}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var userDto dto.UserUpdateDto
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildResponse(400, "参数错误", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return

	}
	ctx.ShouldBind(&userDto)
	valide, err := helper.BindAndValid(ctx, &userDto)
	if !valide {
		response := helper.BuildResponse(400, "参数错误", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := c.userService.UpdateUser(id, userDto)
	response := helper.BuildResponse(200, "success", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response := helper.BuildResponse(400, "参数错误", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := c.userService.DeleteUser(int(id))
	response := helper.BuildResponse(200, "success", res)
	ctx.JSON(http.StatusOK, response)
}

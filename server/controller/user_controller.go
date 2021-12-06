package controller

import (
	"strconv"
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
	c.userService.FindUser(page, limit)
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	panic("not implemented") // TODO: Implement
}

package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/entity"
	"yzlhdy_blog/helper"
	"yzlhdy_blog/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JwtService
}

func NewAuthController(authService service.AuthService, jwtService service.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginForm dto.LoginDTO

	ctx.ShouldBind(&loginForm)
	valid, err := helper.BindAndValid(ctx, &loginForm)
	fmt.Println(valid)
	if !valid {
		response := helper.BuildErrorResponse(403, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := c.authService.VerifyCredential(loginForm.Email, loginForm.Password)
	if v, ok := res.(entity.User); ok {
		genToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = genToken
		fmt.Println(genToken)
		response := helper.BuildResponse(200, "登录成功", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse(401, "用户名或密码错误", helper.EmptyObj{}, err)
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerForm dto.UserCreateDto

	ctx.ShouldBind(&registerForm)
	valid, err := helper.BindAndValid(ctx, &registerForm)
	if !valid {
		response := helper.BuildErrorResponse(401, "参数错误", helper.EmptyObj{}, err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerForm.Email) {
		response := helper.BuildErrorResponse(401, "邮箱已被注册", helper.EmptyObj{}, nil)
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		result := c.authService.CreateUser(registerForm)
		token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(result.ID), 10))
		result.Token = token
		response := helper.BuildResponse(200, "注册成功", result)
		ctx.JSON(http.StatusOK, response)

	}
}

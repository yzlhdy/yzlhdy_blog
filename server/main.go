package main

import (
	"os"
	"yzlhdy_blog/config"
	"yzlhdy_blog/controller"
	"yzlhdy_blog/middleware"
	"yzlhdy_blog/repository"
	"yzlhdy_blog/service"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetUpDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userService    service.UserService       = service.NewUserService(userRepository)
	jwtService     service.JwtService        = service.NewJwtService()
	authService    service.AuthService       = service.NewAuthService(userRepository)

	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	serviceHost := os.Getenv("SERVICE_HOST")
	defer config.CloseDatabase(db)
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(middleware.Translations())
	route.Use(middleware.Loggers())
	authRoutes := route.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	userRoutes := route.Group("api/user")
	{
		userRoutes.GET("/", userController.FindUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
	}
	route.Run(":" + serviceHost)
}

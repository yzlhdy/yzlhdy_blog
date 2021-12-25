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
	db              *gorm.DB                            = config.SetUpDatabaseConnection()
	userRepository  repository.UserRepository           = repository.NewUserRepository(db)
	userService     service.UserService                 = service.NewUserService(userRepository)
	jwtService      service.JwtService                  = service.NewJwtService()
	authService     service.AuthService                 = service.NewAuthService(userRepository)
	classRepository repository.ClassificationRepository = repository.NewClassificationRepository(db)
	classisService  service.ClassificationService       = service.NewClassificationService(classRepository)
	// article
	articleRepo    repository.ArticleRepository = repository.NewArticleRepository(db)
	articleService service.ArticleService       = service.NewArticleService(articleRepo)

	// controller
	authController  controller.AuthController           = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController           = controller.NewUserController(userService, jwtService)
	classController controller.ClassificationController = controller.NewClassificationController(classisService)

	// article
	articleController controller.ArticleController = controller.NewArticleController(articleService)
	uploadController  controller.UploadController  = controller.NewUploadController()
	//  资源分类
	resourceRepository repository.RecategoryRepository = repository.NewRecategoryRepository(db)
	resourceService    service.RecategoryService       = service.NewRecategoryService(resourceRepository)
	resoureController  controller.ResourceController   = controller.NewResourceController(resourceService)
	// 资源
	resourcesRepository repository.ResourcesRepository = repository.NewResourcesRepository(db)
	resourcesService    service.ResourcesService       = service.NewResourcesService(resourcesRepository)
	resourcesController controller.ResourcesController = controller.NewResourcesController(resourcesService)
)

func main() {
	serviceHost := os.Getenv("SERVICE_HOST")
	defer config.CloseDatabase(db)
	route := gin.New()
	route.Use(middleware.Cors())
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(middleware.Translations())
	// route.Use(middleware.Loggers())
	route.MaxMultipartMemory = 8 << 20 // 8 MiB
	authRoutes := route.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := route.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/", userController.FindUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
		userRoutes.PUT("", userController.UpdateUser)
	}

	// 分类
	classRoutes := route.Group("api/classification")
	{
		classRoutes.GET("/:id", classController.FindClassificationById)
		classRoutes.GET("/list", classController.AllClassification)
		classRoutes.POST("", classController.InsertClassification)
		classRoutes.PUT("/:id", classController.UpdateClassification)
		classRoutes.DELETE("/:id", classController.DeleteClassification)
	}
	// 文章
	articleRoutes := route.Group("api/article")
	{
		articleRoutes.GET("", articleController.AllArticle)
		articleRoutes.POST("", articleController.CreateArticle)
		articleRoutes.PUT("/:id", articleController.UpdateArticle)
		articleRoutes.DELETE("/:id", articleController.DeleteArticle)
		articleRoutes.GET("/info/:id", articleController.AllArticleByClassification)
	}
	// 上传
	uploadRoutes := route.Group("api/upload")
	{
		uploadRoutes.POST("/image", uploadController.Upload)
	}
	// 资源分类
	resourceRoutes := route.Group("api/resource")
	{
		resourceRoutes.GET("/list", resoureController.GetAll)
		resourceRoutes.POST("", resoureController.Create)
		resourceRoutes.PUT("/:id", resoureController.Update)
		resourceRoutes.DELETE("/:id", resoureController.Delete)
		resourceRoutes.GET("/:id", resoureController.GetByID)
	}
	// 资源列表
	resourcesRoutes := route.Group("api/resources")
	{
		resourcesRoutes.GET("/list", resourcesController.GetResources)
	}

	route.Run(":" + serviceHost)
}

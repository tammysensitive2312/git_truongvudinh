package app

import (
	"git_truongvudinh/go_web/internal/handler"
	"git_truongvudinh/go_web/internal/repositories"
	"git_truongvudinh/go_web/internal/usecases"
	dao "git_truongvudinh/go_web/pkg"
	"github.com/gin-gonic/gin"
	"log"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()
	router.Use(loggingMiddleware())
	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", userHandler.CreateNewUser)
		userGroup.GET("/get/:id", userHandler.GetUserById)
	}
	userProjectGroup := router.Group("/user_projects")
	{
		userProjectGroup.POST("/create/")
		userProjectGroup.GET("/:id/")
	}
	return router
}

// loggingMiddleware là middleware để log các yêu cầu
func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("call to endpoints user: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

func NewGinEngine() *gin.Engine {
	engine := gin.Default()

	dsn := "root:truong@tcp(localhost:3306)/go_web_example?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := dao.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := usecases.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	engine = SetupRouter(userHandler)
	return engine
}

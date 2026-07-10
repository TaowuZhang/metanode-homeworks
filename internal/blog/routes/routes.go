package routes

import (
	"time"

	"metanode-go-backend-homeworks/internal/blog/controllers"
	"metanode-go-backend-homeworks/internal/blog/middleware"
	"metanode-go-backend-homeworks/internal/blog/repository"
	"metanode-go-backend-homeworks/internal/blog/services"
	"metanode-go-backend-homeworks/internal/blog/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, jwtSecret string, jwtTTLSeconds int64) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	repos := repository.NewRepositories(db)
	svcs := services.NewServices(repos, jwtSecret, timeDuration(jwtTTLSeconds))
	authCtl := controllers.NewAuthController(svcs.Auth)
	postCtl := controllers.NewPostController(svcs.Posts)
	commentCtl := controllers.NewCommentController(svcs.Comments)

	router.GET("/health", func(c *gin.Context) { utils.OK(c, gin.H{"status": "ok"}) })

	api := router.Group("/api/v1")
	api.POST("/auth/register", authCtl.Register)
	api.POST("/auth/login", authCtl.Login)
	api.GET("/posts", postCtl.List)
	api.GET("/posts/:id", postCtl.Get)
	api.GET("/posts/:id/comments", commentCtl.List)

	protected := api.Group("")
	protected.Use(middleware.Auth(jwtSecret))
	protected.GET("/profile", authCtl.Profile)
	protected.POST("/posts", postCtl.Create)
	protected.PUT("/posts/:id", postCtl.Update)
	protected.DELETE("/posts/:id", postCtl.Delete)
	protected.POST("/posts/:id/comments", commentCtl.Create)

	return router
}

func timeDuration(seconds int64) time.Duration {
	if seconds <= 0 {
		seconds = 24 * 3600
	}
	return time.Duration(seconds) * time.Second
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nrad-K/blog-server/controllers"
	"github.com/nrad-K/blog-server/services"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	service := services.NewAppService(db)
	aController := controllers.NewArticleController(service)
	cController := controllers.NewCommentController(service)

	router := gin.Default()
	router.NoRoute(controllers.NoRouteHandler)
	articleRouter := router.Group("/articles")
	{
		articleRouter.POST("/", aController.PostArticleHandler)
		articleRouter.PUT("/", aController.PutArticleHandler)
		articleRouter.GET("/list", aController.ArticleListHandler)
		articleRouter.GET("/:id", aController.ArticleDetailHandler)
		articleRouter.DELETE("/:id", aController.DeleteArticleHandler)
		articleRouter.POST("/like", aController.PostLikeHandler)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.POST("/", cController.PostCommentHandler)
	}

	return router
}

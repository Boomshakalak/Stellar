package routes

import (
	"github.com/gin-gonic/gin"
	"stellar/controllers"
	"stellar/middleware"
)

type PostRouteController struct {
	postController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController {
	return PostRouteController{postController}
}

func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {

	router := rg.Group("posts")
	router.Use(middleware.DeserializeUser())
	router.POST("", pc.postController.CreatePost)
	router.GET("", pc.postController.FindPosts)
	router.PUT("/:postId", pc.postController.UpdatePost)
	router.GET("/:postId", pc.postController.FindPostById)
	router.DELETE("/:postId", pc.postController.DeletePost)
}

package router

import (
	"gin-postgres/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	// User router
	app.GET("/users", controller.GetAllUser)
	app.POST("/users", controller.CreateUser)
	app.GET("/users/:id", controller.GetUserById) // /users/123
	app.PUT("/users/:id", controller.UpdateUser)
	app.DELETE("/users/:id", controller.DeleteUser)

	// Post router
	app.GET("/users/:id/posts", controller.GetPostsOfUser)
	app.POST("/users/:id/posts", controller.CreatePost)
	app.GET("/users/:id/posts/:postId", controller.GetPostDetail)
	app.PUT("/users/:id/posts/:postId", controller.UpdatePost)
	app.DELETE("/users/:id/posts/:postId", controller.DeletePost)
}

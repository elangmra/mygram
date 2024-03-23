package router

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartDB() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.UserRegister)
		userGroup.POST("/login", controllers.UserLogin)
		userGroup.PUT("/edit/:id", middlewares.Authentication(), controllers.UserUpdate)
		userGroup.DELETE("/delete/:id", middlewares.Authentication(), controllers.UserDelete)
	}

	photoGroup := r.Group("/photos")
	{
		photoGroup.Use(middlewares.Authentication())
		photoGroup.POST("/post", controllers.CreatePhoto)
		photoGroup.GET("/get", controllers.GetPhoto)
		photoGroup.PUT("/edit/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoGroup.DELETE("/delete/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentGroup := r.Group("/comments")
	{
		commentGroup.Use(middlewares.Authentication())
		commentGroup.POST("/post", controllers.CreateComment)
		commentGroup.GET("/get", controllers.GetComment)
		commentGroup.PUT("/edit/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentGroup.DELETE("/delete/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialmediasGroup := r.Group("/socialmedias")
	{
		socialmediasGroup.Use(middlewares.Authentication())
		socialmediasGroup.POST("/post", controllers.CreateSocialMedia)
		socialmediasGroup.GET("/get", controllers.GetSocialMedia)
		socialmediasGroup.PUT("/edit/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialmediasGroup.DELETE("/delete/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}

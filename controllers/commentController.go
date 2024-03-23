package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userId

	if err := db.Debug().Create(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})

}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))
	comments := []models.Comment{}

	if err := db.Debug().Preload("User").Preload("Photo").Where("user_id = ?", userId).Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	var response []gin.H
	for _, comment := range comments {
		commentData := gin.H{
			"id":         comment.ID,
			"message":    comment.Message,
			"photo_id":   comment.PhotoID,
			"user_id":    comment.UserID,
			"created_at": comment.CreatedAt,
			"updated_at": comment.UpdatedAt,
		}
		userData := gin.H{
			"id":       comment.User.ID,
			"email":    comment.User.Email,
			"username": comment.User.Username,
		}
		photoData := gin.H{
			"id":        comment.Photo.ID,
			"title":     comment.Photo.Title,
			"caption":   comment.Photo.Caption,
			"photo_url": comment.Photo.PhotoURL,
			"user_id":   comment.Photo.UserID,
		}
		commentData["User"] = userData
		commentData["Photo"] = photoData
		response = append(response, commentData)
	}
	c.JSON(http.StatusOK, response)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()

	commentId := c.Param("commentId")

	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	db.First(&Comment, commentId)

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userId

	if err := db.Model(&Comment).Where("id = ?", commentId).Updates(&Comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"updated_at": Comment.UpdatedAt,
	})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentId := c.Param("commentId")
	Comment := models.Comment{}
	db.First(&Comment, commentId)
	db.Delete(&Comment)
	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}

package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {

	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userId

	if err := db.Debug().Create(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})

}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))
	photos := []models.Photo{}

	if err := db.Debug().Preload("User").Where("user_id = ?", userId).Find(&photos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	var response []gin.H
	for _, photo := range photos {
		photoData := gin.H{
			"id":         photo.ID,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoURL,
			"user_id":    photo.UserID,
			"created_at": photo.CreatedAt,
			"updated_at": photo.UpdatedAt,
		}
		userData := gin.H{
			"email":    photo.User.Email,
			"username": photo.User.Username,
		}
		photoData["User"] = userData
		response = append(response, photoData)
	}
	c.JSON(http.StatusOK, response)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()

	photoId := c.Param("photoId")

	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	db.First(&Photo, photoId)

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userId

	if err := db.Model(&Photo).Where("id = ?", photoId).Updates(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoURL,
		"user_id":    Photo.UserID,
		"updated_at": Photo.UpdatedAt,
	})

}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	photoId := c.Param("photoId")
	Photo := models.Photo{}
	db.First(&Photo, photoId)
	db.Delete(&Photo)
	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}

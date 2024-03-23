package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	contentType := helpers.GetContentType(c)

	sosmed := models.SocialMedia{}

	if contentType == appJson {
		c.ShouldBindJSON(&sosmed)
	} else {
		c.ShouldBind(&sosmed)
	}

	sosmed.UserID = userId

	if err := db.Debug().Create(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":               sosmed.ID,
		"name":             sosmed.Name,
		"social_media_url": sosmed.SocialMediaURL,
		"user_id":          sosmed.UserID,
		"created_at":       sosmed.CreatedAt,
	})

}

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()

	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	sosmeds := []models.SocialMedia{}

	if err := db.Debug().Preload("User").Where("user_id = ?", userId).Find(&sosmeds).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	var response gin.H
	var data []gin.H

	for _, sosmeds := range sosmeds {
		sosmedsData := gin.H{
			"id":               sosmeds.ID,
			"name":             sosmeds.Name,
			"social_media_url": sosmeds.SocialMediaURL,
			"user_id":          sosmeds.UserID,
			"created_at":       sosmeds.CreatedAt,
			"updated_at":       sosmeds.UpdatedAt,
		}
		userData := gin.H{
			"id":       sosmeds.User.ID,
			"email":    sosmeds.User.Email,
			"username": sosmeds.User.Username,
		}
		sosmedsData["User"] = userData
		data = append(data, sosmedsData)

	}

	response = gin.H{
		"social_medias": data,
	}

	c.JSON(http.StatusOK, response)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()

	sosmedId := c.Param("socialMediaId")

	userId := uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))
	contentType := helpers.GetContentType(c)

	sosmed := models.SocialMedia{}
	db.First(&sosmed, sosmedId)

	if contentType == appJson {
		c.ShouldBindJSON(&sosmed)
	} else {
		c.ShouldBind(&sosmed)
	}

	sosmed.UserID = userId

	if err := db.Model(&sosmed).Where("id = ?", sosmedId).Updates(&sosmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":               sosmed.ID,
		"name":             sosmed.Name,
		"social_media_url": sosmed.SocialMediaURL,
		"user_id":          sosmed.UserID,
		"updated_at":       sosmed.UpdatedAt,
	})

}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	sosmedId := c.Param("socialMediaId")
	sosmed := models.SocialMedia{}
	db.First(&sosmed, sosmedId)
	db.Delete(&sosmed)
	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}

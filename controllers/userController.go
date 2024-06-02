package controllers

import (
    "net/http"
    "myapi/database"
    "myapi/models"
    "github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
    var users []models.User

    if result := database.DB.Find(&users); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": users})
}


func FindUser(c *gin.Context) {
    var user models.User
    if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user := models.User{Name: input.Name, Email: input.Email}
    database.DB.Create(&user)
    c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
    var user models.User
    if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }

    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Model(&user).Updates(input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
    var user models.User
    if err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
        return
    }
    database.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"data": true})
}

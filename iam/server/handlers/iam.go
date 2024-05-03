package handlers

import (
	"KubeInsight/iam/server/params"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (iam *IamHandler) Login(c *gin.Context) {
	var user params.UserLogin
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login Failed:"})
		return
	}

	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userName is NULL"})
		return
	} else if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is NULL"})
		return
	}

	err := iam.iamService.Login(user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully Login"})
}

func (iam *IamHandler) Auth(c *gin.Context) {
	var user params.UserLogin
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Failed:"})
		return
	}
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userName is NULL"})
		return
	}

	if err := iam.iamService.Auth(user.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error:": "token is error"}) //TODO 修复这一块的error
		log.Println("Error:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully Auth"})
}

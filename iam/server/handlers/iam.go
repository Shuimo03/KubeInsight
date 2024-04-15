package handlers

import (
	"KubeInsight/iam/server/params"
	"github.com/gin-gonic/gin"
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

	if err := iam.iamService.Login(user.Name, user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully Login"})
}

package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *Router) LoginHandler(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userID, role, err := r.userAuth.AuthenticateUser(
		c.Request.Context(),
		input.Username,
		input.Password,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token, err := r.authService.GenerateToken(userID, role, time.Hour*24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

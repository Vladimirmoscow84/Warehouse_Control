package handlers

import (
	"net/http"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/auth"
	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
	"github.com/gin-gonic/gin"
)

func (r *Router) RegisterHandler(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email"`
		RoleID   int    `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	user := &model.User{
		Username:     input.Username,
		PasswordHash: hashedPassword,
		Email:        input.Email,
		RoleID:       input.RoleID,
	}

	id, err := r.userCreator.CreateUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

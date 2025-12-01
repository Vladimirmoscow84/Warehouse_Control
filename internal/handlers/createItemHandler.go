package handlers

import (
	"net/http"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
	"github.com/gin-gonic/gin"
)

func (r *Router) CreateItemHandler(c *gin.Context) {

	valUserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user did not found in context"})
		return
	}
	userID, ok := valUserID.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id type"})
		return
	}
	var item model.Item
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := r.itemCreator.CreateItem(c.Request.Context(), &item, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})

}

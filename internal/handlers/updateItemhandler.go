package handlers

import (
	"net/http"
	"strconv"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
	"github.com/gin-gonic/gin"
)

func (r *Router) UpdateItemHandler(c *gin.Context) {
	valUserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found"})
		return
	}

	userID, ok := valUserID.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id value"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	var item model.Item
	err = c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	item.ID = id

	err = r.itemsUpdater.UpdateItem(c.Request.Context(), &item, userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "item updated"})
}

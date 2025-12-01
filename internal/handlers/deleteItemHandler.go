package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Router) DeleteItemHandler(c *gin.Context) {
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
	err = r.itemsDeleter.DeleteItem(c.Request.Context(), id, userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "item deleted"})
}

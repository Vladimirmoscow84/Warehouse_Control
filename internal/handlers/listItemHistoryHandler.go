package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Router) ListItemHistoryHandler(c *gin.Context) {
	idStr := c.Param("id")
	itemID, err := strconv.Atoi(idStr)
	if err != nil || itemID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	history, err := r.itemsGetter.ListItemHistory(c.Request.Context(), itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, history)
}

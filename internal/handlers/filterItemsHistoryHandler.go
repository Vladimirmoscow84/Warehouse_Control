package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (r *Router) FilterItemsHistoryHandler(c *gin.Context) {

	idStr := c.Param("id")
	itemID, err := strconv.Atoi(idStr)
	if err != nil || itemID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	//опциональные query-параметры
	var userIDPtr *int
	if userIDStr := c.Query("user_id"); userIDStr != "" {
		userID, err := strconv.Atoi(userIDStr)
		if err != nil || userID <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
			return
		}
		userIDPtr = &userID
	}

	var actionTypePtr *string
	if actionType := c.Query("action_type"); actionType != "" {
		actionTypePtr = &actionType
	}

	var fromPtr, toPtr *time.Time
	if fromStr := c.Query("from"); fromStr != "" {
		fromTime, err := time.Parse(time.RFC3339, fromStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid from date"})
			return
		}
		fromPtr = &fromTime
	}
	if toStr := c.Query("to"); toStr != "" {
		toTime, err := time.Parse(time.RFC3339, toStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid to date"})
			return
		}
		toPtr = &toTime
	}

	history, err := r.itemsGetter.FilterItemHistory(c.Request.Context(), itemID, userIDPtr, actionTypePtr, fromPtr, toPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, history)
}

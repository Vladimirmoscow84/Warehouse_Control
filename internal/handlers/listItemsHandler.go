package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) ListItemsHandler(c *gin.Context) {
	items, err := r.itemsGetter.ListItems(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

package handlers

import (
	"context"
	"time"

	"github.com/Vladimirmoscow84/Warehouse_Control/internal/auth"
	"github.com/Vladimirmoscow84/Warehouse_Control/internal/middleware"
	"github.com/Vladimirmoscow84/Warehouse_Control/internal/model"
	"github.com/Vladimirmoscow84/Warehouse_Control/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/wb-go/wbf/ginext"
)

type itemCreator interface {
	CreateItem(ctx context.Context, i *model.Item, userID int) (int, error)
}

type itemsGetter interface {
	GetItem(ctx context.Context, id int) (*model.Item, error)
	ListItems(ctx context.Context) ([]*model.Item, error)
	ListItemHistory(ctx context.Context, itemID int) ([]*model.ItemHistory, error)
	FilterItemHistory(ctx context.Context, itemID int, userID *int, actionType *string, from, to *time.Time) ([]*model.ItemHistory, error)
}

type itemsUpdater interface {
	UpdateItem(ctx context.Context, i *model.Item, userID int) error
}

type itemsDeleter interface {
	DeleteItem(ctx context.Context, id, userID int) error
}

type authService interface {
	CheckToken(token string) (*auth.UserClaims, error)
}

type Router struct {
	Router       *ginext.Engine
	itemCreator  itemCreator
	itemsGetter  itemsGetter
	itemsUpdater itemsUpdater
	itemsDeleter itemsDeleter
	authService  authService
}

func New(router *ginext.Engine, iCreator itemCreator, iGetter itemsGetter, iUpdater itemsUpdater, iDeleter itemsDeleter, aService authService) *Router {
	return &Router{
		Router:       router,
		itemCreator:  iCreator,
		itemsGetter:  iGetter,
		itemsUpdater: iUpdater,
		itemsDeleter: iDeleter,
		authService:  aService,
	}
}

func (r *Router) Routes(jwtSecret string) {
	authMiddleware := middleware.AuthMiddleware(auth.New(jwtSecret))

	adminMiddleware := middleware.RequreRoles(service.RoleAdmin)
	managerMiddleware := middleware.RequreRoles(service.RoleAdmin, service.RoleManager)

	r.Router.POST("/items", authMiddleware, managerMiddleware, r.CreateItemHandler)
	r.Router.GET("/items/:id", authMiddleware, r.GetItemHandler)
	r.Router.GET("/items", authMiddleware, r.ListItemsHandler)
	r.Router.PUT("/items/:id", authMiddleware, managerMiddleware, r.UpdateItemHandler)
	r.Router.DELETE("/items/:id", authMiddleware, adminMiddleware, r.DeleteItemHandler)

	r.Router.GET("/items/:id/history", authMiddleware, r.ListItemHistoryHandler)
	r.Router.GET("/items/:id/history/filter", authMiddleware, r.FilterItemsHistoryHandler)

	r.Router.GET("/", func(c *gin.Context) { c.File("./web/index.html") })
	r.Router.Static("/static", "./web")

}

package routes

import (
	"api-center/controllers"
	"api-center/middleware"
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(r *gin.Engine, db *gorm.DB) {
	h := &controllers.Handler{DB: db.Debug()}
	// Create a group for api/v1
	v1 := r.Group("/api/v1")
	v1.POST("/users", h.CreateUser)
}

func SetupRouter(ctx context.Context, db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	UserRoute(r, db)

	// r.POST("/users", controllers.CreateUser)
	// r.GET("/users/:id", controllers.GetUser)
	return r
}

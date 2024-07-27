package routes

import (
	"api-center/controllers"
	"api-center/middleware"
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(r *gin.Engine, db *gorm.DB) {
	h := &controllers.UserHandler{DB: db.Debug()}
	// Create a group for api/v1
	v1 := r.Group("/api/v1")
	v1.POST("/users", h.CreateUser)
	v1.POST("/login", h.Login)
}

func TeamsRoute(r *gin.Engine, db *gorm.DB) {
	h := &controllers.TeamHandler{DB: db.Debug()}
	// Create a group for api/v1
	v1 := r.Group("/api/v1")
	v1.GET("/teams", h.GetListTeams)
}

func ProjectRoute(r *gin.Engine, db *gorm.DB) {
	h := &controllers.ProjectHandler{DB: db.Debug()}
	// Create a group for api/v1
	v1 := r.Group("/api/v1")
	v1.GET("/projects", h.GetListProjects)
	v1.POST("/project", h.CreateProject)
	v1.PUT("/project", h.UpdateProject)
	v1.DELETE("/project", h.DeleteProject)
	v1.POST("/project/favorite", h.AddProjectFavorite)
}

func SetupRouter(ctx context.Context, db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Recovery())

	UserRoute(r, db)
	TeamsRoute(r, db)
	ProjectRoute(r, db)

	return r
}

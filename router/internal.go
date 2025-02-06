package router

import (
	"context"

	controller "example.com/m/internal/controllers"
	"example.com/m/internal/repository"
	service "example.com/m/internal/services"
	postgres "example.com/m/pkg/db"

	"github.com/gin-gonic/gin"

	"github.com/omniful/go_commons/http"
)

func Initialize(ctx context.Context, s *http.Server) (err error) {

	// Setup WMS Routes
	wmsV1 := s.Engine.Group("/api/v1")

	newRepository := repository.NewRepository(postgres.GetCluster().DbCluster)
	newService := service.NewService(newRepository)
	controller := controller.NewController(newService)
	wmsV1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "mst"})
	})

	wmsV1.GET("/hubs/:hub_id", controller.GetHub) 
	return

}

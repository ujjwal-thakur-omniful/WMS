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
	hubcontroller := controller.NewController(newService)

	skuRepo := repository.NewSKURepository(postgres.GetCluster().DbCluster)
	skuService := service.NewSKUService(skuRepo)
	skuController := controller.NewSkuController(skuService)
	wmsV1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "mst"})
	})

	wmsV1.GET("/hubs/:hub_id", hubcontroller.GetHub)
	wmsV1.POST("/hubs", hubcontroller.CreateHub)

	wmsV1.GET("/sku/:sku_id", skuController.GetSKU)
	 wmsV1.GET("/sku/:tenant_id/:seller_id/:sku_id", skuController.GetSKUByTenantIDAndSellerID)
	wmsV1.POST("/sku", skuController.CreateSKU)
	return

}

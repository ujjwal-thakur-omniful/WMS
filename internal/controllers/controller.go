package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"example.com/m/internal/domain"
	"example.com/m/internal/request"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	hubService domain.HubService
}

var ctrl *Controller
var ctrlOnce sync.Once

func NewController(svc domain.HubService) *Controller {
	ctrlOnce.Do(func() {
		ctrl = &Controller{
			hubService: svc,
		}
	})
	return ctrl
}
type SkuController struct {
	skuService domain.SkuService
}

func NewSkuController(svc domain.SkuService) *SkuController {
	return &SkuController{
		skuService: svc,
	}
}

func (skuController *SkuController) GetSKU(c *gin.Context) {
	skuID := c.Param("sku_id")
	skuIDUint, err := strconv.ParseUint(skuID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch SKU id"})
		return
	}

	skuDetails, cusErr := skuController.skuService.GetSku(c, skuIDUint)
	if cusErr.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		return
	}
	fmt.Println("skuDetails", skuDetails)

	c.JSON(http.StatusOK, gin.H{"sku": skuDetails})
}


func (skuController *SkuController) CreateSKU(c *gin.Context) {
	var newSKU request.Sku // Assuming request.SKU is defined similarly to request.Hub
	if err := c.ShouldBindJSON(&newSKU); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdSKU, cusErr := skuController.skuService.CreateSku(c, newSKU)

	if cusErr.Exists() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create SKU"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"sku": createdSKU})
}
func (skuController *SkuController) GetSKUByTenantIDAndSellerID(c *gin.Context) {
	tenantID := c.Param("tenant_id")
	sellerID := c.Param("seller_id")
	skuID := c.Param("sku_id")

	tenantIDUint, err := strconv.ParseUint(tenantID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tenant id"})
		return
	}

	sellerIDUint, err := strconv.ParseUint(sellerID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch seller id"})
		return
	}

	skuIDUint, err := strconv.ParseUint(skuID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch SKU id"})
		return
	}

	skuDetails, cusErr := skuController.skuService.GetSKUByTenantIDAndSellerID(c, tenantIDUint, sellerIDUint, skuIDUint)
	if cusErr.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
		return
	}
	fmt.Println("skuDetails", skuDetails)

	c.JSON(http.StatusOK, gin.H{"sku": skuDetails})
}

func (hubController *Controller) GetHub(c *gin.Context) {
	hubID := c.Param("hub_id")
	hubIDUint, err := strconv.ParseUint(hubID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hub id"})
		return
	}

	hubDetails, cusErr := hubController.hubService.GetHubDetails(c, hubIDUint)
	if cusErr.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
		return
	}
	fmt.Println("hubDetails", hubDetails)

	c.JSON(http.StatusOK, gin.H{"hub": hubDetails})

}
func (hubController *Controller) CreateHub(c *gin.Context) {
	var newHub request.Hub
	if err := c.ShouldBindJSON(&newHub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdHub, cusErr := hubController.hubService.CreateHub(c, newHub)
	if cusErr.Exists() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hub"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"hub": createdHub})
}


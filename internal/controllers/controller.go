package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"example.com/m/internal/domain"
	request "example.com/m/internal/requests"
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

type InventoryController struct {
	inventoryService domain.InventoryService
}

func NewInventoryController(svc domain.InventoryService) *InventoryController {
	return &InventoryController{
		inventoryService: svc,
	}
}

func (inventoryController *InventoryController) GetInventoryDetails(c *gin.Context) {
	sellerID := c.Param("seller_id")
	hubID := c.Param("hub_id")

	sellerIDUint, err := strconv.ParseUint(sellerID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch seller id"})
		return
	}

	hubIDUint, err := strconv.ParseUint(hubID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hub id"})
		return
	}

	inventoryDetails, cusErr := inventoryController.inventoryService.GetInventoryDetails(c, sellerIDUint, hubIDUint)
	if cusErr.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"inventory": inventoryDetails})
}

func (inventoryController *InventoryController) UpdateInventory(c *gin.Context) {
	inventoryID := c.Param("inventory_id")
	skuID := c.Param("sku_id")

	inventoryIDUint, err := strconv.ParseUint(inventoryID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inventory id"})
		return
	}

	skuIDUint, err := strconv.ParseUint(skuID, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch SKU id"})
		return
	}

	// Assuming the request body contains the updated inventory details
	var updatedInventory request.Inventory
	if err := c.ShouldBindJSON(&updatedInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedInventoryResponse, cusErr := inventoryController.inventoryService.UpdateInventory(c, inventoryIDUint, skuIDUint)
	if cusErr.Exists() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inventory": updatedInventoryResponse})
}

func (inventoryController *InventoryController) CreateInventory(c *gin.Context) {
	var newInventory request.Inventory
	if err := c.ShouldBindJSON(&newInventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdInventory, cusErr := inventoryController.inventoryService.CreateInventory(c, newInventory)
	if cusErr.Exists() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"inventory": createdInventory})
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

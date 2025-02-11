// package validate

// import (
// 	"context"
// 	"fmt"
// 	"github.com/omniful/go_commons/error"
// 	"example.com/m/internal/repository"
// )

// func ValidateHubAndSKU(db *repository.Repository) func(ctx context.Context) error {
// 	return func(ctx context.Context) error {
// 		// Logic to validate hub and SKU
// 		// This is a placeholder for actual validation logic
// 		fmt.Println("Validating hub and SKU...")
// 		// Assume validation passes for now
// 		return nil
// 	}
// }

//	func ValidateAndUpdateInventory(db *repository.Repository) func(ctx context.Context) error {
//		return func(ctx context.Context) error {
//			// Logic to validate and update inventory
//			// This is a placeholder for actual validation and update logic
//			fmt.Println("Validating and updating inventory...")
//			// Assume validation and update passes for now
//			return nil
//		}
//	}
package validate

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	model "example.com/m/internal/models"
	responses "example.com/m/internal/response"
	"github.com/gin-gonic/gin"

	"github.com/omniful/go_commons/db/sql/postgres"
	"gorm.io/gorm"
)

type ValidateInventoryRequest struct {
	SKUID           string `json:"sku_id"`
	QuantityOrdered int    `json:"quantity_ordered"`
	HubID           string `json:"hub_id"`
}

type HubInventory struct {
	ID                    uint   `gorm:"primary_key"`
	SKUID                 string `gorm:"column:sku_id"`
	HubID                 string `gorm:"column:hub_id"`
	QuantityOfEachProduct int    `gorm:"column:quantity_of_each_product"`
}

func ValidateAndUpdateInventory(db *postgres.DbCluster) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		fmt.Println("Validate And Update Inventory fxn called inside WMS")

		var req ValidateInventoryRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
			return
		}

		var inventory HubInventory
		result := db.GetMasterDB(ctx).Where("sku_id = ? AND hub_id = ?", req.SKUID, req.HubID).First(&inventory)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
			return
		}

		if req.QuantityOrdered > inventory.QuantityOfEachProduct {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient inventory"})
			return
		}

		// Reduce inventory quantity
		if err := inventory.ReduceQuantity(db.GetMasterDB(ctx), req.QuantityOrdered); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Inventory validation and update successful"})
	}
}

func (inventory *HubInventory) ReduceQuantity(db *gorm.DB, quantity int) error {
	// Decrease the quantity
	inventory.QuantityOfEachProduct -= quantity

	// Update the record in the database
	if err := db.Save(inventory).Error; err != nil {
		return err
	}

	log.Println("Inventory updated ")

	return nil
}
func ValidateHubAndSKU(db *postgres.DbCluster) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("ValidateHubAndSKU function called")

		var request model.ValidateOrderRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid request payload",
				"details": err.Error(),
			})
			return
		}

		log.Println("parsed request from ValidateHubAndSKU function is: ", request)

		// Now we can access the values using:
		// request.SKUID
		// request.HubID

		var wg sync.WaitGroup
		respChan := make(chan model.ValidationResponse, 2)


		// Start hub validation goroutine
		wg.Add(1)
		go func() {
			defer wg.Done()
			var hub responses.Hub
			result := db.GetMasterDB(ctx).First(&hub, request.HubID)


			if result.Error != nil {
				respChan <- model.ValidationResponse{
					IsValid: false,
					Error:   fmt.Errorf("hub validation failed: %v", result.Error),
				}
				return
			}
			respChan <- model.ValidationResponse{IsValid: true}
		}()

		// Start SKU validation goroutine
		wg.Add(1)
		go func() {
			defer wg.Done()
			var sku responses.Sku
			result := db.GetMasterDB(ctx).First(&sku, request.SKUID)
			if result.Error != nil {
				respChan <- model.ValidationResponse{
					IsValid: false,
					Error:   fmt.Errorf("SKU validation failed: %v", result.Error),
				}
				return
			}
			respChan <- model.ValidationResponse{IsValid: true}
		}()

		// Wait for both goroutines to complete
		go func() {
			wg.Wait()
			close(respChan)
		}()

		// Process results
		var finalResponse model.ValidationResponse
		finalResponse.IsValid = true


		for resp := range respChan {
			if !resp.IsValid {
				finalResponse.IsValid = false
				finalResponse.Error = resp.Error
				break
			}
		}

		// Send response based on validation result
		if finalResponse.IsValid {
			ctx.JSON(http.StatusOK, gin.H{"message": "Validation successful"})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": finalResponse.Error.Error()})
		}
	}
}
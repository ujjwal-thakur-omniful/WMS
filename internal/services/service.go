package service

import (
	"context"
	"fmt"

	"example.com/m/internal/domain"
	request "example.com/m/internal/requests"

	responses "example.com/m/internal/response"
	error2 "github.com/omniful/go_commons/error"
)

type Service struct {
	repository domain.HubRepository
}

func NewService(repo domain.HubRepository) *Service {
	return &Service{
		repository: repo,
	}
}

type SkuService struct {
	repository domain.SkuRepository
}

func NewSKUService(repo domain.SkuRepository) *SkuService {
	return &SkuService{
		repository: repo,
	}
}

type InventoryService struct {
	repository domain.InventoryRepository
}

func NewInventoryService(repo domain.InventoryRepository) *InventoryService {
	return &InventoryService{
		repository: repo,
	}
}

func (s *InventoryService) GetInventoryDetails(ctx context.Context, sellerID, hubID uint64) (responses.Inventory, error2.CustomError) {
	inventory, cusErr := s.repository.GetInventoryDetails(ctx, sellerID, hubID)
	if cusErr.Exists() {
		return responses.Inventory{}, cusErr
	}
	fmt.Println("Inventory details", inventory)
	return inventory, cusErr

}

func (s *InventoryService) UpdateInventory(ctx context.Context, inventoryID, skuID uint64) (responses.Inventory, error2.CustomError) {
	// Logic to update inventory
	inventory, cusErr := s.repository.UpdateInventory(ctx, inventoryID, skuID)
	if cusErr.Exists() {

		return responses.Inventory{}, cusErr
	}
	fmt.Println("Updated inventory", inventory)
	return inventory, cusErr
}

func (s *InventoryService) CreateInventory(ctx context.Context, inventory request.Inventory) (responses.Inventory, error2.CustomError) {
	createdInventory, cusErr := s.repository.CreateInventory(ctx, inventory)
	if cusErr.Exists() {
		return responses.Inventory{}, cusErr
	}
	fmt.Println("Created inventory", createdInventory)
	return createdInventory, cusErr
}

func (s *SkuService) GetSku(ctx context.Context, skuID uint64) (responses.Sku, error2.CustomError) {
	sku, cusErr := s.repository.GetSku(ctx, skuID)
	if cusErr.Exists() {
		return responses.Sku{}, cusErr
	}
	fmt.Println("sku", sku)
	return sku, cusErr
}

func (s *SkuService) CreateSku(ctx context.Context, sku request.Sku) (responses.Sku, error2.CustomError) {
	skuResponse, cusErr := s.repository.CreateSku(ctx, sku)
	if cusErr.Exists() {
		return responses.Sku{}, cusErr
	}
	fmt.Println("Created SKU", skuResponse)
	return skuResponse, cusErr
}
func (s *SkuService) GetSKUByTenantIDAndSellerID(ctx context.Context, tenantID, sellerID, skuID uint64) (responses.Sku, error2.CustomError) {
	sku, cusErr := s.repository.GetSKUByTenantIDAndSellerID(ctx, tenantID, sellerID, skuID)
	if cusErr.Exists() {
		return responses.Sku{}, cusErr
	}
	fmt.Println("sku by tenant and seller", sku)
	return sku, cusErr
}

func (s *Service) GetHubDetails(ctx context.Context, hubID uint64) (responses.Hub, error2.CustomError) {
	// Fetch hub details from the database
	hub, cusErr := s.repository.GetHub(ctx, hubID)
	if cusErr.Exists() {
		return responses.Hub{}, cusErr

	}
	fmt.Println("hub", hub)
	return hub, cusErr
}
func (s *Service) CreateHub(ctx context.Context, hub request.Hub) (responses.Hub, error2.CustomError) {
	// Create a new hub in the database
	hubResponse, cusErr := s.repository.CreateHub(ctx, hub)
	if cusErr.Exists() {
		return responses.Hub{}, cusErr
	}
	fmt.Println("Created hub", hubResponse)
	return hubResponse, cusErr
}

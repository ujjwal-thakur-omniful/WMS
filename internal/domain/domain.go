package domain

import (
	"context"

	request "example.com/m/internal/requests"
	responses "example.com/m/internal/response"
	error2 "github.com/omniful/go_commons/error"
)




type HubService interface {
	GetHubDetails(c context.Context,  hub_id uint64) (responses.Hub, error2.CustomError)
	CreateHub(ctx context.Context, hub request.Hub) (responses.Hub,  error2.CustomError)

}
type HubRepository interface {
	GetHub(c context.Context,  hub_id uint64) (responses.Hub, error2.CustomError)
	CreateHub(ctx context.Context, hub request.Hub) (responses.Hub,  error2.CustomError)
}
type SkuService interface {
	GetSku(c context.Context,  sku_id uint64) (responses.Sku, error2.CustomError)
	CreateSku(ctx context.Context, sku request.Sku) (responses.Sku,  error2.CustomError)
	GetSKUByTenantIDAndSellerID(c context.Context, tenant_id, seller_id uint64, sku_id uint64) (responses.Sku, error2.CustomError)

}
type SkuRepository interface {
	GetSku(c context.Context,  sku_id uint64) (responses.Sku, error2.CustomError)
	CreateSku(ctx context.Context, sku request.Sku) (responses.Sku,  error2.CustomError)
	GetSKUByTenantIDAndSellerID(c context.Context, tenant_id, seller_id uint64, sku_id uint64) (responses.Sku, error2.CustomError)

}
type InventoryService interface {
	GetInventoryDetails(c context.Context, seller_id uint64, hub_id uint64) (responses.Inventory, error2.CustomError)
	UpdateInventory(ctx context.Context, inventory_id uint64, sku_id uint64) (responses.Inventory,  error2.CustomError)


	CreateInventory(ctx context.Context, inventory request.Inventory) (responses.Inventory,  error2.CustomError)
}
type InventoryRepository interface {
	GetInventoryDetails(c context.Context, seller_id uint64, hub_id uint64) (responses.Inventory, error2.CustomError)
	UpdateInventory(ctx context.Context, inventory_id uint64, sku_id uint64) (responses.Inventory,  error2.CustomError)
	CreateInventory(ctx context.Context, inventory request.Inventory) (responses.Inventory,  error2.CustomError)
}
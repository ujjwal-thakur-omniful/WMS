package domain

import (
	"context"

	responses "example.com/m/internal/response"
	error2 "github.com/omniful/go_commons/error"
	request "example.com/m/internal/request"
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
	GetHubDetails(c context.Context,  hub_id uint64) (responses.Hub, error2.CustomError)
	CreateHub(ctx context.Context, hub request.Hub) (responses.Hub,  error2.CustomError)

}
type InventoryRepository interface {
	GetHub(c context.Context,  hub_id uint64) (responses.Hub, error2.CustomError)
	CreateHub(ctx context.Context, hub request.Hub) (responses.Hub,  error2.CustomError)
}
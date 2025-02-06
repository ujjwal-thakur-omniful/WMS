package repository

import (
	"context"
	"fmt"

	"example.com/m/internal/request"
	responses "example.com/m/internal/response"
	"github.com/omniful/go_commons/db/sql/postgres"
	error2 "github.com/omniful/go_commons/error"

	"sync"
)

type Repository struct {
	db *postgres.DbCluster
}

var repo *Repository
var repoOnce sync.Once

func NewRepository(db *postgres.DbCluster) *Repository {
	repoOnce.Do(func() {
		repo = &Repository{
			db: db,
		}
	})

	return repo
}
type SKURepository struct {
	db *postgres.DbCluster
}

func NewSKURepository(db *postgres.DbCluster) *SKURepository {
	return &SKURepository{
		db: db,
	}
}

func (r *SKURepository) GetSku(c context.Context, sku_id uint64) (responses.Sku, error2.CustomError) {
	var sku responses.Sku

	// Perform the query to get the SKU details by ID
	if err := r.db.GetMasterDB(c).Where("id = ?", sku_id).First(&sku).Error; err != nil {
		return responses.Sku{}, error2.CustomError{}
	}

	fmt.Println("sku", sku)

	return sku, error2.CustomError{}
}

func (r *SKURepository) CreateSku(ctx context.Context, sku request.Sku) (responses.Sku, error2.CustomError) {
	var skuResponse responses.Sku

	// Perform the query to create a new SKU
	if err := r.db.GetMasterDB(ctx).Create(&sku).Scan(&skuResponse).Error; err != nil {
		return responses.Sku{}, error2.CustomError{}
	}

	fmt.Println("Created SKU", skuResponse)

	return skuResponse, error2.CustomError{}
}
func (r *SKURepository) GetSKUByTenantIDAndSellerID(c context.Context, tenant_id, seller_id, sku_id uint64) (responses.Sku, error2.CustomError) {
	var sku responses.Sku

	// Perform the query to get the SKU details by tenant ID, seller ID, and SKU ID
	if err := r.db.GetMasterDB(c).Where("tenant_id = ? AND seller_id = ? AND id = ?", tenant_id, seller_id, sku_id).First(&sku).Error; err != nil {
		return responses.Sku{}, error2.CustomError{}
	}

	fmt.Println("sku by tenant and seller", sku)

	return sku, error2.CustomError{}
}



func (r *Repository) GetHub(c context.Context, hub_id uint64) (responses.Hub, error2.CustomError) {

	var hub responses.Hub

	// Perform the query to get the hub details by ID
	if err := r.db.GetMasterDB(c).Where("id = ?", hub_id).First(&hub).Error; err != nil {

		return responses.Hub{}, error2.CustomError{}

	}

	// Handle other types of errors (e.g., database connection issues)

	fmt.Println("hub", hub)

	return hub, error2.CustomError{}

}
func (r *Repository) CreateHub(ctx context.Context, hub request.Hub) (responses.Hub, error2.CustomError) {
	var hubResponse responses.Hub

	// Perform the query to create a new hub
	if err := r.db.GetMasterDB(ctx).Create(&hub).Scan(&hubResponse).Error; err != nil {
		return responses.Hub{}, error2.CustomError{}
	}

	fmt.Println("Created hub", hubResponse)

	return hubResponse, error2.CustomError{}
}

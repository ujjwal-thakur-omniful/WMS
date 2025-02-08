package responses

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type (
	Hub struct {
		ID        int64     `json:"id"`
		Name      string     `json:"name"`
		TenantID  int64      `json:"tenant_id"`
		Location  string     `json:"location"`
		CreatedAt time.Time   `json:"created_at"`
		CreatedBy int64      `json:"created_by"`
		UpdatedAt time.Time  `json:"updated_at"`
		UpdatedBy int64      `json:"updated_by"`
		DeletedAt null.Time `json:"deleted_at,omitempty"`
	}
	Sku struct {
		ID          int64           `gorm:"primaryKey;autoIncrement" json:"id"`
		SellerID    int64           `gorm:"not null" json:"seller_id"`     // Foreign key to the seller
		Attributes  map[string]interface{} `gorm:"type:jsonb" json:"attributes"` // JSONB field for storing attributes
		PPU         float64         `gorm:"not null" json:"ppu"`           // Price Per Unit
		CreatedAt   int64           `gorm:"autoCreateTime" json:"created_at"`
		UpdatedAt   int64           `gorm:"autoUpdateTime" json:"updated_at"`
	}
	Inventory struct {
		ID          int64     `json:"id"`
		SkuID       int64     `json:"sku_id"`     // Foreign key to the seller
		HubID       int64     `json:"hub_id"`        // Foreign key to the hub
		Quantity    int64     `json:"quantity"`      // Quantity of the inventory item
		CreatedAt   time.Time `json:"created_at"`    // Timestamp for when the inventory was created
		UpdatedAt   time.Time `json:"updated_at"`    // Timestamp for when the inventory was last updated
		DeletedAt   null.Time `json:"deleted_at,omitempty"` // Optional field for soft deletion
	}
)
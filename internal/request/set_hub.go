package request

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
)
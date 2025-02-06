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
)
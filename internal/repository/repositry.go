package repository

import (
	"context"
	"fmt"

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

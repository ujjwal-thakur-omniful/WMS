package domain

import (
	"context"

	responses "example.com/m/internal/response"
	error2 "github.com/omniful/go_commons/error"
)






type HubService interface {
	GetHubDetails(c context.Context,  hub_id uint64) (responses.Hub, error2.CustomError)

}
type HubRepository interface {
	GetHub(c context.Context,  hub_id uint64) (responses.Hub, error2.CustomError)

}
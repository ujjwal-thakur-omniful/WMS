package models

type ValidationResponse struct {
	IsValid bool
	Error   error
}
type ValidateOrderRequest struct {
	SKUID string `json:"sku_id"`
	HubID string `json:"hub_id"`
}

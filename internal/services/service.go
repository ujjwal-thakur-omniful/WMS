package service

import (
	"context"
	"fmt"

	"example.com/m/internal/domain"
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

func (s *Service) GetHubDetails(ctx context.Context, hubID uint64) (responses.Hub, error2.CustomError) {
	// Fetch hub details from the database
	hub, cusErr := s.repository.GetHub(ctx, hubID)
	if cusErr.Exists() {
		return responses.Hub{}, cusErr

	}
	fmt.Println("hub", hub)
	return hub, cusErr
}

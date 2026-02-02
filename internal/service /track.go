package service

import (
	"context"
	"time"

	"github.com/Sushka-s-makom/parcel-tracker/internal/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, track domain.Track) error
}

type TrackService struct {
	repo Repository
}

func NewTrackService(repo Repository) *TrackService {
	return &TrackService{repo: repo}
}

func (s *TrackService) CreateTrack(ctx context.Context, number, carrier string) (domain.Track, error) {
	track := domain.Track{
		ID:           uuid.New(),
		TrackNumber:  number,
		Carrier:      carrier,
		Status:       domain.StatusNew,
		LastUpdateAt: time.Now().UTC(),
		CreatedAt:    time.Now().UTC(),
	}

	err := s.repo.Create(ctx, track)
	return track, err
}

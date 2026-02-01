package repository

import (
	"context"

	"github.com/Sushka-s-makom/parcel-tracker/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TrackRepository struct {
	db *pgxpool.Pool
}

func NewTrackRepository(db *pgxpool.Pool) *TrackRepository {
	return &TrackRepository{db: db}
}

func (r *TrackRepository) Create(ctx context.Context, track domain.Track) error {
	query := `INSERT INTO tracks (id, track_number, carrier, status, last_update_at, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, track.ID, track.TrackNumber, track.Carrier, track.Status, track.LastUpdateAt, track.CreatedAt)
	return err
}

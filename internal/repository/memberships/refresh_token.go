package memberships

import (
	"context"

	"github.com/rifaiAhmed/fastcampus/internal/model/memberships"
	"github.com/rs/zerolog/log"
)

func (r *repository) InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, expired_token, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.ExpiredToken, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		log.Error().Err(err).Msg("failed to create refresh token")
		return err
	}

	return nil
}

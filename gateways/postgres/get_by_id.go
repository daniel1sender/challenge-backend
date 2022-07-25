package postgres

import (
	"context"

	"github.com/daniel1sender/alura-flix/domain"
	"github.com/daniel1sender/alura-flix/domain/entity"
	"github.com/jackc/pgx/v4"
)

const getByIDQuery = `SELECT id, title, description, url 
FROM videos
WHERE id = $1`

func (fs VideoStorage) GetByID(ctx context.Context, id string) (entity.Video, error) {
	video := entity.Video{}
	err := fs.Conn.QueryRow(ctx, getByIDQuery, id).Scan(&video.ID, &video.Title, &video.Description, &video.URL)
	if err == pgx.ErrNoRows {
		return entity.Video{}, domain.ErrNoVideoFound
	} else if err != nil {
		return entity.Video{}, err
	}

	return video, nil
}

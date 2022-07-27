package postgres

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain/videos"
	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
	"github.com/jackc/pgx/v4"
)

const getByIDQuery = `SELECT id, title, description, url 
FROM videos
WHERE id = $1`

func (fs VideoStorage) GetByID(ctx context.Context, id string) (entities.Video, error) {
	video := entities.Video{}
	err := fs.Conn.QueryRow(ctx, getByIDQuery, id).Scan(&video.ID, &video.Title, &video.Description, &video.URL)
	if err == pgx.ErrNoRows {
		return entities.Video{}, videos.ErrNoVideoFound
	} else if err != nil {
		return entities.Video{}, err
	}

	return video, nil
}

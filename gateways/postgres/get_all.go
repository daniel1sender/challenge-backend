package postgres

import (
	"context"

	"github.com/daniel1sender/alura-flix/domain"
	"github.com/daniel1sender/alura-flix/domain/entity"
)

const getAllQuery = `SELECT id, title, description, url FROM videos`

func (fs VideoStorage) GetAll(ctx context.Context) ([]entity.Video, error) {
	var video entity.Video
	var allVideo []entity.Video
	rows, err := fs.Conn.Query(ctx, getAllQuery)
	if err != nil {
		return []entity.Video{}, err
	}

	for rows.Next() {
		err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.URL)
		if err != nil {
			return []entity.Video{}, err
		}
		allVideo = append(allVideo, video)
	}

	if len(allVideo) == 0 {
		return []entity.Video{}, domain.ErrNoVideoFound
	}

	return allVideo, nil
}

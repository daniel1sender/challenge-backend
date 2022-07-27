package postgres

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain"
	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

const getAllQuery = `SELECT id, title, description, url FROM videos`

func (fs VideoStorage) GetAll(ctx context.Context) ([]entities.Video, error) {
	var video entities.Video
	var allVideo []entities.Video
	rows, err := fs.Conn.Query(ctx, getAllQuery)
	if err != nil {
		return []entities.Video{}, err
	}

	for rows.Next() {
		err := rows.Scan(&video.ID, &video.Title, &video.Description, &video.URL)
		if err != nil {
			return []entities.Video{}, err
		}
		allVideo = append(allVideo, video)
	}

	if len(allVideo) == 0 {
		return []entities.Video{}, domain.ErrNoVideoFound
	}

	return allVideo, nil
}

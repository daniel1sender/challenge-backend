package videos

import (
	"context"
	"fmt"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

const insertQuery = `INSERT INTO videos(
	id,
	title,
	description,
	url
	) VALUES (
	$1,
	$2,
	$3,
	$4)`

func (fs VideoStorage) Insert(ctx context.Context, video entities.Video) error {
	if _, err := fs.Conn.Exec(ctx, insertQuery, video.ID, video.Title, video.Description, video.URL); err != nil {
		return fmt.Errorf("unable to insert the video due to: %w", err)
	}
	return nil
}

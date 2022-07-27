package videos

import (
	"context"
	"fmt"
)

const deleteByIDQuery = `DELETE FROM videos WHERE id=$1`

func (fs VideoStorage) DeleteById(ctx context.Context, id string) error {
	_, err := fs.Conn.Exec(ctx, deleteByIDQuery, id)
	if err != nil {
		return fmt.Errorf("could not delete the video due to: %s", err)
	}
	return nil
}

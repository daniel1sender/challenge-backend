package videos

import (
	"context"
	"fmt"
)

const updateDescriptionQuery = `UPDATE videos
	SET 
	description=$1
	WHERE
	id=$2`

func (fs VideoStorage) UpdateDescriptionByID(ctx context.Context, id, description string) error {
	_, err := fs.Conn.Exec(ctx, updateDescriptionQuery, description, id)
	if err != nil {
		return fmt.Errorf("unable to update the description due to: %s", err)
	}
	return nil
}

package postgres

import (
	"context"
	"fmt"
)

const updateTitleQuery = `UPDATE videos 
	SET 
	title=$1
	WHERE 
	id=$2`

func (fs VideoStorage) UpdateTitleByID(ctx context.Context, id, title string) error {
	_, err := fs.Conn.Exec(ctx, updateTitleQuery, title, id)
	if err != nil {
		return fmt.Errorf("unable to update the title due to: %s", err)
	}
	return nil
}

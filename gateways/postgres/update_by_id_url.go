package postgres

import (
	"context"
	"fmt"
)

const updateURLQuery = `UPDATE videos 
	SET 
	url=$1
	WHERE 
	id=$2`

func (fs VideoStorage) UpdateURLByID(ctx context.Context, id, url string) error {
	_, err := fs.Conn.Exec(ctx, updateURLQuery, url, id)
	if err != nil {
		return fmt.Errorf("unable to update the url due to: %s", err)
	}
	return nil
}

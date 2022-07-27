package categories

import (
	"context"
	"fmt"
)

const deleteByIDQuery = `DELETE FROM categories WHERE id=$1`

func (cs CategoryStorage) DeleteById(ctx context.Context, id string) error {
	_, err := cs.Conn.Exec(ctx, deleteByIDQuery, id)
	if err != nil {
		return fmt.Errorf("could not delete the category due to: %s", err)
	}
	return nil
}
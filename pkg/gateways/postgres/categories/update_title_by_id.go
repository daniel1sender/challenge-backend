package categories

import (
	"context"
	"fmt"
)

const updateTitleQuery = `UPDATE categories
	SET 
	title=$1
	WHERE
	id=$2`

func (cs CategoryStorage) UpdateTitleByID(ctx context.Context, id, title string) error {
	_, err := cs.Conn.Exec(ctx, updateTitleQuery, title, id)
	if err != nil {
		return fmt.Errorf("unable to update the title due to: %s", err)
	}
	return nil
}

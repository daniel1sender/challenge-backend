package categories

import (
	"context"
	"fmt"
)

const updateColorQuery = `UPDATE categories
	SET 
	color=$1
	WHERE
	id=$2`

func (cs CategoryStorage) UpdateColorByID(ctx context.Context, id, color string) error {
	_, err := cs.Conn.Exec(ctx, updateColorQuery, color, id)
	if err != nil {
		return fmt.Errorf("unable to update the color due to: %s", err)
	}
	return nil
}

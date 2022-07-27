package categories

import (
	"context"
	"fmt"

	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

const insertQuery = `INSERT INTO categories(
	id,
	title,
	color,
	) VALUES (
	$1,
	$2,
	$3)`

func (cs CategoryStorage) Insert(ctx context.Context, category entities.Category) error {
	if _, err := cs.Conn.Exec(ctx, insertQuery, category.ID, category.Title, category.Color); err != nil {
		return fmt.Errorf("unable to insert the category due to: %w", err)
	}
	return nil
}

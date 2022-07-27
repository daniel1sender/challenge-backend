package categories

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain/categories"
	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
	"github.com/jackc/pgx/v4"
)

const getByIDQuery = `SELECT id, title, color
FROM categories
WHERE id = $1`

func (cs CategoryStorage) GetByID(ctx context.Context, id string) (entities.Category, error) {
	category := entities.Category{}
	err := cs.Conn.QueryRow(ctx, getByIDQuery, id).Scan(&category.ID, &category.Title, &category.Color)
	if err == pgx.ErrNoRows {
		return entities.Category{}, categories.ErrNoCategoryFound
	} else if err != nil {
		return entities.Category{}, err
	}

	return category, nil
}

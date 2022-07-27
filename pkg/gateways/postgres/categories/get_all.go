package categories

import (
	"context"

	"github.com/daniel1sender/alura-flix/pkg/domain/categories"
	"github.com/daniel1sender/alura-flix/pkg/domain/entities"
)

const getAllQuery = `SELECT id, title, color FROM categories`

func (fs CategoryStorage) GetAll(ctx context.Context) ([]entities.Category, error) {
	var category entities.Category
	var allCategories []entities.Category
	rows, err := fs.Conn.Query(ctx, getAllQuery)
	if err != nil {
		return []entities.Category{}, err
	}

	for rows.Next() {
		err := rows.Scan(&category.ID, &category.Title, &category.Color)
		if err != nil {
			return []entities.Category{}, err
		}
		allCategories = append(allCategories, category)
	}

	if len(allCategories) == 0 {
		return []entities.Category{}, categories.ErrNoCategoryFound
	}

	return allCategories, nil
}

package usecases

import "context"

func (u CategoryUseCase) UpdateColorByID(ctx context.Context, id, color string) error {
	return u.storage.UpdateColorByID(ctx, id, color)
}

package usecases

import "context"

func (u CategoryUseCase) UpdateTitleByID(ctx context.Context, id, title string) error {
	return u.storage.UpdateTitleByID(ctx, id, title)
}

package usecases

import "context"

func (u CategoryUseCase) DeleteById(ctx context.Context, id string) error {
	return u.storage.DeleteById(ctx, id)
}

package usecases

import (
	"context"
	"fmt"
)

func (u VideoUseCase) UpdateByID(ctx context.Context, id, title, description, url string) error {
	video, err := u.storage.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("could not update the video on registers due to: %w", err)
	}

	if len(title) != 0 && title != video.Title {
		err := u.storage.UpdateTitleByID(ctx, id, title)
		if err != nil {
			return fmt.Errorf("could not update the video on registers due to: %w", err)
		}
	}

	if len(description) != 0 && description != video.Description {
		err := u.storage.UpdateDescriptionByID(ctx, id, description)
		if err != nil {
			return fmt.Errorf("could not update the video on registers due to: %w", err)
		}
	}

	if len(url) != 0 && url != video.URL {
		err := u.storage.UpdateURLByID(ctx, id, url)
		if err != nil {
			return fmt.Errorf("could not update the video on registers due to: %w", err)
		}
	}

	return nil
}

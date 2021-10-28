package rate

import (
	"context"
	"malctl/client"
)

func RateAnime(id, rating int) error {
	ctx := context.Background()
	client.Client.RateAnime(ctx, id, rating)
	return nil
}

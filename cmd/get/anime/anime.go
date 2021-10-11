package anime

import (
	"context"
	"log"
	"malctl/client"
	"strconv"
)

func GetAnimeInfo(args []string) {
	ctx := context.Background()
	animeID, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	client.Client.AnimeDetails(ctx, animeID)
}

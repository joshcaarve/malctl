package anime

import (
	"context"
	"log"
	"malctl/client"
	"strconv"
)

func GetAnimeInfo(animeIDArg string) {
	ctx := context.Background()
	animeID, err := strconv.Atoi(animeIDArg)
	if err != nil {
		log.Fatal(err)
	}
	client.Client.AnimeDetails(ctx, animeID)
}

func GetEnglishName(id string) string {
	ctx := context.Background()
	animeID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	return client.Client.AnimeIDEnglish(ctx, animeID)
}

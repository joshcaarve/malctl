package user

import (
	"context"
	"fmt"
	"malctl/client"

	"github.com/nstratos/go-myanimelist/mal"
)

func GetUserInfo(args []string) {
	if args[0] == "info" {
		ctx := context.Background()
		client.Client.UserMyInfo(ctx)
	} else if args[0] == "anime" {

		if args[1] == "watch" {
			ctx := context.Background()
			var status mal.AnimeStatus = "watching"
			client.Client.UserAnimeList(ctx, status)
		} else if args[1] == "complete" {
			ctx := context.Background()
			var status mal.AnimeStatus = "completed"
			client.Client.UserAnimeList(ctx, status)
		} else if args[1] == "hold" {
			ctx := context.Background()
			var status mal.AnimeStatus = "on_hold"
			client.Client.UserAnimeList(ctx, status)
		} else if args[1] == "plan" {
			ctx := context.Background()
			var status mal.AnimeStatus = "plan_to_watch"
			client.Client.UserAnimeList(ctx, status)
		} else {
			fmt.Printf("[ERROR] Unrecognized: %s\n", args[1])
		}
	}
}

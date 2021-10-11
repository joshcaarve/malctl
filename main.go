package main

import (
	"context"
	"log"
	"malctl/auth"
	"malctl/client"

	"github.com/nstratos/go-myanimelist/mal"
)

func main() {
	tokenClient, err := auth.CreateTokenClient()
	if err != nil {
		log.Fatal(err)
	}
	c := client.DemoClient{
		Client: mal.NewClient(tokenClient),
	}
	ctx := context.Background()
	err = c.Showcase(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

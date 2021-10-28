package main

import (
	"malctl/client"
	"malctl/cmd"
)

func main() {
	client.Init()
	cmd.Execute()
}

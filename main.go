package main

import (
	"encoding/json"
	"fmt"
	"malctl/animehash"
	"malctl/client"
	"malctl/cmd"
	"malctl/xmlparse"
	"os"
)

const (
	url = "https://myanimelist.net/sitemap/anime-000.xml"
)

func xmlToString(urlSetXML xmlparse.UrlSetXML) []string {
	var animeUrls []string
	for _, urlXML := range urlSetXML.Url {
		animeUrls = append(animeUrls, urlXML.Loc)
	}
	return animeUrls
}

func main() {
	urlSetXML := xmlparse.GetUrlSetXML(url)
	animeUrls := xmlToString(urlSetXML)
	animeHash := animehash.GetAnimeHashList(animeUrls)
	if animeHash != nil {
		animeHashJSON, _ := json.Marshal(animeHash)
		fmt.Println(string(animeHashJSON))
	}

	os.Exit(0)
	client.Init()
	cmd.Execute()
}

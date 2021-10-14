package seed

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"malctl/animehash"
	"malctl/db"
	"malctl/xmlparse"
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

func SeedAnimeHash(custom string) {
	urlSetXML := xmlparse.GetUrlSetXML(url)
	animeUrls := xmlToString(urlSetXML)
	animeHash := animehash.GetAnimeHashList(animeUrls)
	file, _ := json.MarshalIndent(animeHash, "", " ")

	err := ioutil.WriteFile(db.DbPath, file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

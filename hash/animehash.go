package animehash

import (
	"malctl/cmd/get/anime"
	"malctl/util"
	"strings"
)

const (
	urlConst = "https://myanimelist.net/anime/"
)

type AnimeName struct {
	Name        string
	EnglishName string
}

type AnimeIDHash map[string]AnimeName

func GetAnimeHashList(urls []string) AnimeIDHash {
	animeIDHash := make(AnimeIDHash)
	for _, url := range urls {
		var animeName AnimeName
		urlShort := strings.ReplaceAll(url, urlConst, "")
		animeID := util.StringBefore(urlShort, "/")
		animeName.EnglishName = anime.GetEnglishName(animeID)
		animeName.Name = cleanName(util.StringAfter(urlShort, "/"))
		animeIDHash[animeID] = animeName
	}
	return animeIDHash
}

func cleanName(animeName string) string {
	cleanAnimeName := strings.ReplaceAll(animeName, "__", ": ")
	cleanAnimeName = strings.ReplaceAll(cleanAnimeName, "_", " ")
	return cleanAnimeName
}

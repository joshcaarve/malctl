package animehash

import (
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
		animeName.Name = util.StringAfter(urlShort, "/")
		animeIDHash[animeID] = animeName
	}
	return animeIDHash
}

package guess

import (
	"fmt"
	"log"
	"malctl/util"
	"strings"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func GetIDFromGuess(guess string) string {
	search := fmt.Sprintf("anime mal %s", guess)
	var id string
	searchResults, err := googlesearch.Search(nil, search)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, result := range searchResults {
			if strings.Contains(result.URL, "myanimelist.net") {
				id = util.StringAfter(result.URL, "https://myanimelist.net/anime/")
				id = util.StringBefore(id, "/")
				break
			}
		}
	}
	return id
}

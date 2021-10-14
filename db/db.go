package db

import (
	"encoding/json"
	"log"
	"malctl/animehash"
	"os"
)

const (
	DbPath = "db/db.json"
)

func GetAnimeNameID(animeID string) string {
	clientFile, err := os.ReadFile(DbPath)
	if err != nil {
		log.Fatal(err)
	}
	var animeIDHash animehash.AnimeIDHash

	err = json.Unmarshal(clientFile, &animeIDHash)
	if err != nil {
		log.Fatal(err)
	}
	return animeIDHash[animeID].Name
}

func GetAnimeIDFromName(name string) string {
	clientFile, err := os.ReadFile(DbPath)
	if err != nil {
		log.Fatal(err)
	}
	var animeIDHash animehash.AnimeIDHash

	err = json.Unmarshal(clientFile, &animeIDHash)
	if err != nil {
		log.Fatal(err)
	}
	for id, animeName := range animeIDHash {
		if animeName.Name == name {
			return id
		}
	}
	return ""
}

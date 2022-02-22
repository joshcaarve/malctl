# malctl
malctl is a command line interface to interact with a user account myanimelist.com.

This was made purely for fun.

# Building
```sh
$ go build -o malctl
```

# Example runs

## Get user anime (complete, watching, hold, plan)
Will print the id, episodes, rating, and name.

I cutoff the output in this README
```sh
$ ./malctl get user anime complete
Caching refreshed oauth2 token...
34213   12/ 12  7 Getsuyoubi no Tawawa
40602   12/ 12  8 7 Seeds 2nd Season
38735   12/ 12  8 7 Seeds
33047   10/ 10  5 Fate/Extra: Last Encore
37999   12/ 12  8 Kaguya-sama: Love is War
40938   13/ 13  4 Higehiro: After Being Rejected, I Shaved and Took in a High School Runaway
37450   13/ 13  6 Rascal Does Not Dream of Bunny Girl Senpai
...
```

## Get all shows the user is watching
```sh
$ ./malctl get user anime watch
 9379    0/ 12  0 Ground Control to Psychoelectric Girl
  481   52/224  8 Yu-Gi-Oh!
14813   12/ 13  0 My Teen Romantic Comedy SNAFU
 2001    2/ 27  0 Gurren Lagann
   80    3/ 43  0 Mobile Suit Gundam
```

## Seed your local machine to hash Anime IDs to Names
Basically pulls an `xml` from mal and parses it into a makeshift database to store some info on local to make the malctl commands more flexible.

I also implemented a `guess` function which fills in the gaps of the seeding.

This will create a `db.json` inside `malctl/db`
See the README there for more info.

This command takes a little bit to push processing.
```sh
$ ./malctl seed
Processing seed...
Seed complete
```


## Get information on a specific anime from ID
```sh
./malctl get anime id 30                       
Neon Genesis Evangelion
ID: 30
English: Neon Genesis Evangelion
Type: TV
Episodes: 26
Premiered: Fall 1995
Studios: Gainax Tatsunoko Production
Source: Original
Genres: Action Avant Garde Drama Mecha Psychological Sci-Fi
Duration: 24 min. per ep.
```

## Get information on specific anime from name
NOTE: name only works if you have seeded a local DB.
```sh
$ ./malctl get anime name "Neon Genesis Evangelion"
Neon Genesis Evangelion
ID: 30
English: Neon Genesis Evangelion
Type: TV
Episodes: 26
Premiered: Fall 1995
Studios: Gainax Tatsunoko Production
Source: Original
Genres: Action Avant Garde Drama Mecha Psychological Sci-Fi
Duration: 24 min. per ep.
```

## Getting anime information via a guess
If you don't know the ID or didn't seed the db and don't know the exact name, you can guess by having malctl perform google search.

This feature was added for fun.
```sh
$ ./malctl get anime guess "NGE"
Neon Genesis Evangelion
ID: 30
English: Neon Genesis Evangelion
Type: TV
Episodes: 26
Premiered: Fall 1995
Studios: Gainax Tatsunoko Production
Source: Original
Genres: Action Avant Garde Drama Mecha Psychological Sci-Fi
Duration: 24 min. per ep.
```


## Rate a user anime based on a guess
```sh
$ ./malctl rate anime guess "Fate Extra" 5
Status: "completed", Score: 5, Episodes Watched: 10, Comments:
```


## Shortcomings
Of course more commands can be added such as login and update.
Also, it's only written to pull anime, but can easily add manga feature as well.

Also, didn't include steps to install the binary globally.

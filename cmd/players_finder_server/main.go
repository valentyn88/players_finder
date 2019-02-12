package main

import (
	"net/http"

	"encoding/json"

	"log"

	"github.com/valentyn88/players_finder/client"
	"github.com/valentyn88/players_finder/storage"
)

const countGoroutines = 10

func main() {

	tData := map[string]struct{}{
		"Germany":           struct{}{},
		"England":           struct{}{},
		"France":            struct{}{},
		"Spain":             struct{}{},
		"Manchester United": struct{}{},
		"Arsenal":           struct{}{},
		"Chelsea":           struct{}{},
		"Barcelona":         struct{}{},
		"Real Madrid":       struct{}{},
		"Bayern Munich":     struct{}{},
	}
	teamStorage := storage.InitTeam(tData)

	done := make(chan struct{})
	idsCh := make(chan int, 10)

	pp := make(map[string]storage.Player)
	playerStorage := storage.Init(pp)

	for i := 0; i < countGoroutines; i++ {
		go client.GetTeam(done, idsCh, "https://vintagemonster.onefootball.com/api/teams/en/%s.json", teamStorage, playerStorage)
	}

	go func() {
		var i = 1
		for {
			if teamStorage.Len() == 0 {
				close(idsCh)
				close(done)
				break
			}

			idsCh <- i
			i++
		}
	}()

	<-done

	sortedPlayers := playerStorage.SortedByNameList()

	h := handler{players: sortedPlayers}
	http.HandleFunc("/players", h.playersHandler)

	if err := http.ListenAndServe(":7008", nil); err != nil {
		log.Fatal(err)
	}
}

type handler struct {
	players []storage.Player
}

func (h handler) playersHandler(w http.ResponseWriter, _ *http.Request) {
	pp, err := json.Marshal(h.players)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(pp)
}

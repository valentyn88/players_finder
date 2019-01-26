package main

import (
	"fmt"

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
		"FC Bayern Munich":  struct{}{},
	}
	teamStorage := storage.InitTeam(tData)

	done := make(chan struct{})
	idsCh := make(chan int, 10)

	pp := make(map[string]storage.Player)
	playerStorage := storage.Init(pp)

	for i := 0; i < countGoroutines; i++ {
		go client.GetTeam(done, idsCh, teamStorage, playerStorage)
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
	for _, player := range sortedPlayers {
		fmt.Println(player)
	}
}

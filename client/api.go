package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/valentyn88/players_finder/storage"
)

// Response api response.
type Response struct {
	Status interface{} `json:"status"`
	Data   Data        `json:"data"`
}

// Data response data.
type Data struct {
	Team Team `json:"team"`
}

// Team describes team with players.
type Team struct {
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

// Player describes single player.
type Player struct {
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Age       string `json:"age"`
}

// GetTeam get team
func GetTeam(done chan struct{}, idsCh chan int, reqURL string, teamStorage storage.Team, playerStorage storage.Storage) {
	for {
		select {
		case id, ok := <-idsCh:
			if !ok {
				return
			}

			resp, err := http.Get(fmt.Sprintf(reqURL, strconv.Itoa(id)))
			if err != nil {
				log.Printf("couldn't get response error: %s\n", err.Error())
				continue
			}

			r, err := getResponse(resp)
			if err != nil {
				log.Println(err.Error())
			}

			if !teamStorage.Find(r.Data.Team.Name) {
				continue
			}

			for _, p := range r.Data.Team.Players {
				sp := storage.Player{FullName: p.Name, Age: p.Age, BirthDate: p.BirthDate, Teams: []string{r.Data.Team.Name}}
				playerStorage.Upsert(sp)
			}

			teamStorage.Remove(r.Data.Team.Name)
		case <-done:
			return
		}
	}
}

func getResponse(resp *http.Response) (Response, error) {
	var (
		err error
		r   Response
	)

	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return r, errors.New("Status is not 200")
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return r, errors.Errorf("couldn't read response body error: %s\n", err.Error())
	}

	if err := json.Unmarshal(content, &r); err != nil {
		return r, errors.Errorf("couldn't unmarshal response error: %s\n", err.Error())
	}

	return r, nil
}

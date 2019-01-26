package storage

import (
	"reflect"
	"testing"
)

func TestStorage_Upsert(t *testing.T) {
	players := make(map[string]Player)
	plSt := Init(players)

	plSt.Upsert(Player{FullName: "test", Age: "25", BirthDate: "1995-02-14", Teams: []string{"Spain"}})
	plSt.Upsert(Player{FullName: "test", Age: "25", BirthDate: "1995-02-14", Teams: []string{"England"}})

	sortedPlayers := plSt.SortedByNameList()
	expected := Player{FullName: "test",
		Age:       "25",
		BirthDate: "1995-02-14",
		Teams:     []string{"Spain", "England"},
	}

	if !reflect.DeepEqual(expected, sortedPlayers[0]) {
		t.Fatalf("expected %v and got %v are unequal\n", expected, sortedPlayers[0])
	}
}

func TestPlayer_String(t *testing.T) {
	p := Player{FullName: "test test",
		Age:       "25",
		BirthDate: "1995-02-14",
		Teams:     []string{"Spain", "England"},
	}
	got := p.String()
	expected := "test test; 25; Spain, England"
	if expected != got {
		t.Fatalf("expected %v and got %v are unequal\n", expected, got)
	}
}

func TestStorage_SortedByNameList(t *testing.T) {
	players := make(map[string]Player)
	plSt := Init(players)

	plSt.Upsert(Player{FullName: "Antonio Fricko", Age: "25", BirthDate: "1995-02-14", Teams: []string{"Spain"}})
	plSt.Upsert(Player{FullName: "Cristian Ronaldo", Age: "30", BirthDate: "1988-02-14", Teams: []string{"Portugal"}})
	plSt.Upsert(Player{FullName: "Bengamin Fricko", Age: "25", BirthDate: "1995-02-14", Teams: []string{"Spain"}})

	got := plSt.SortedByNameList()
	expected := []string{"Antonio Fricko; 25; Spain", "Bengamin Fricko; 25; Spain", "Cristian Ronaldo; 30; Portugal"}
	for i, p := range got {
		if p.String() != expected[i] {
			t.Fatalf("expected %v and got %v are unequal\n", expected, got)
		}
	}
}

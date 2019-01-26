package storage

import (
	"reflect"
	"sync"
	"testing"
)

func TestTeam_Remove(t *testing.T) {
	teams := map[string]struct{}{"England": struct{}{}, "Spain": struct{}{}, "Turkey": struct{}{}}
	teamStorage := InitTeam(teams)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		teamStorage.Remove("England")
		wg.Done()
	}()

	go func() {
		teamStorage.Remove("Spain")
		wg.Done()
	}()

	wg.Wait()

	expected := map[string]struct{}{"Turkey": struct{}{}}
	got := teamStorage.data
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v and got %v are unequal\n", expected, got)
	}
}

func TestTeam_Find(t *testing.T) {
	teams := map[string]struct{}{"England": struct{}{}, "Spain": struct{}{}, "Turkey": struct{}{}}
	teamStorage := InitTeam(teams)

	testCases := []struct {
		name     string
		team     string
		expected bool
	}{
		{name: "team exists", team: "England", expected: true},
		{name: "team doesn't exists", team: "Belgium", expected: false},
	}

	for _, tc := range testCases {
		t.Log("Test case: ", tc.name)

		got := teamStorage.Find(tc.team)
		if tc.expected != got {
			t.Fatalf("expected %v and got %v are unequal\n", tc.expected, got)
		}
	}
}

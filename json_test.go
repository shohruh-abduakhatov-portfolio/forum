package main

import (
	"log"
	"net/http"

	//"net/http/httptest"
	"encoding/json"
	"io/ioutil"
	"testing"

	internal "./internal"
)

func TestIndexPage(t *testing.T) {
	resp, errResp := http.Get("http://localhost:8181/artists")
	if errResp != nil {
		log.Fatalln(errResp)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var actualArtistLocalStruct []internal.Artist
	errJSON := json.Unmarshal(body, &actualArtistLocalStruct)
	if errJSON != nil {
		log.Fatalln(errJSON)
	}

	expected, errExpected := internal.GetAllArtists()
	if errExpected != nil {
		log.Fatalln(errExpected)
	}
	for i := range expected {
		if expected[i].ID != actualArtistLocalStruct[i].ID {
			t.Logf("Expected %g, got %g", expected[i].ID, actualArtistLocalStruct[i].ID)
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Result struct {
	AgentName string
	Data      interface{}
	Error     error
}

type Agent interface {
	Name() string
	Execute() Result
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
}

type LocationAgent struct {
	ArtistID int
}

func (l LocationAgent) Name() string {
	return "LocationAgent"
}

func (l LocationAgent) Execute() Result {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", l.ArtistID)

	resp, err := http.Get(url)

	if err != nil {
		return Result{
			AgentName: l.Name(),
			Data:      nil,
			Error:     err,
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Result{
			AgentName: l.Name(),
			Data:      nil,
			Error:     err,
		}
	}

	var location Location
	err = json.Unmarshal(body, &location)

	if err != nil {
		return Result{
			AgentName: l.Name(),
			Data:      nil,
			Error:     err,
		}
	}

	return Result{
		AgentName: l.Name(),
		Data:      location,
		Error:     nil,
	}
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsURL string   `json:"locations"`
	DatesURL     string   `json:"concertDates"`
	RelationsURL string   `json:"relations"`
}

type ArtistsAgent struct {
	ArtistID int
}

func (a ArtistsAgent) Name() string {
	return "ArtistAgent"
}

func (a ArtistsAgent) Execute() Result {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", a.ArtistID)

	resp, err := http.Get(url)

	if err != nil {
		return Result{
			AgentName: a.Name(),
			Data:      nil,
			Error:     err,
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Result{
			AgentName: a.Name(),
			Data:      nil,
			Error:     err,
		}
	}

	var artist Artist

	err = json.Unmarshal(body, &artist)

	if err != nil {
		return Result{
			AgentName: a.Name(),
			Data:      nil,
			Error:     err,
		}
	}

	return Result{
		AgentName: a.Name(),
		Data:      artist,
		Error:     nil,
	}
}

func main() {
	var wg sync.WaitGroup

	results := make(chan Result, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		agent := ArtistsAgent{ArtistID: 2}
		results <- agent.Execute()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		agent := LocationAgent{ArtistID: 2}
		results <- agent.Execute()
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Error != nil {
			fmt.Printf("[%s] Error: %v\n", result.AgentName, result.Error)
			continue
		}

		switch result.AgentName {
		case "ArtistAgent":
			artist := result.Data.(Artist)
			fmt.Printf("[%s] Artist: %s\n", result.AgentName, artist.Name)
			fmt.Printf("[%s] Members: %v\n", result.AgentName, artist.Members)
		case "LocationAgent":
			location := result.Data.(Location)
			fmt.Printf("[%s] Locations:\n", result.AgentName)
			for _, loc := range location.Locations {
				fmt.Printf("  - %s\n", loc)
			}
		}
	}
}

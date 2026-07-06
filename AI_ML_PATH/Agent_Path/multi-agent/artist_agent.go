package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

// Date represents the data from the Groupie-Tracker dates API
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// DatesAgent fetches concert dates for an artist
type DatesAgent struct {
	ArtistID int
}

func (d DatesAgent) Name() string {
	return "DatesAgent"
}

func (d DatesAgent) Execute() Result {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", d.ArtistID)

	resp, err := http.Get(url)
	if err != nil {
		return Result{
			AgentName: d.Name(),
			Data:      nil,
			Error:     err,
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{
			AgentName: d.Name(),
			Data:      nil,
			Error:     err,
		}
	}

	var date Date
	err = json.Unmarshal(body, &date)
	if err != nil {
		return Result{
			AgentName: d.Name(),
			Data:      nil,
			Error:     err,
		}
	}

	return Result{
		AgentName: d.Name(),
		Data:      date,
		Error:     nil,
	}
}

// Relation represents the data from the Groupie-Tracker relation API
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// RelationAgent fetches relation data for an artist
type RelationAgent struct {
	ArtistID int
}

func (r RelationAgent) Name() string {
	return "RelationAgent"
}

func (r RelationAgent) Execute() Result {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", r.ArtistID)

	resp, err := http.Get(url)
	if err != nil {
		return Result{
			AgentName: r.Name(),
			Data:      nil,
			Error:     err,
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{
			AgentName: r.Name(),
			Data:      nil,
			Error:     err,
		}
	}

	var relation Relation
	err = json.Unmarshal(body, &relation)
	if err != nil {
		return Result{
			AgentName: r.Name(),
			Data:      nil,
			Error:     err,
		}
	}

	return Result{
		AgentName: r.Name(),
		Data:      relation,
		Error:     nil,
	}
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
	reader := bufio.NewReader(os.Stdin)

	var artistID int
	fmt.Print("Input the Artist ID: ")
	fmt.Scan(&artistID)
	reader.ReadString('\n')

	fmt.Print("Enter your query: ")
	query, _ := reader.ReadString('\n')
	query = strings.TrimSpace(query)

	orchestrator := NewOrchestrator()
	agents := orchestrator.SelectAgents(query, artistID)

	coordinator := Newcoordinator(agents)
	results := coordinator.Run()

	// Create synthesizer
	synthesizer := NewSynthesizer()

	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("[%s] Error: %v\n", result.AgentName, result.Error)
			continue
		}

		// Use synthesizer to convert raw data to natural language
		response := synthesizer.Synthesize(result)
		fmt.Printf("[%s] %s\n", result.AgentName, response)
	}
}

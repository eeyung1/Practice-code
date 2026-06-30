package main

import "strings"

type Orchestrator struct {}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{}
}

func (o *Orchestrator) SelectAgents(query string, artistID int) []Agent {
	var agents []Agent

	lower := strings.ToLower(query)

	wantsArtist := strings.Contains(lower, "artist") ||
	strings.Contains(lower, "band") ||
	strings.Contains(lower, "tell me about") ||
	strings.Contains(lower, "who is")

	wantsLocation := strings.Contains(lower, "location") ||
	strings.Contains(lower, "where") ||
	strings.Contains(lower, "concert") ||
	strings.Contains(lower, "perform") ||
	strings.Contains(lower, "show")

	if wantsArtist {
		agents = append(agents, ArtistsAgent{ArtistID: artistID})
	}

	if wantsLocation {
		agents = append(agents, LocationAgent{ArtistID: artistID})
	}

	if len(agents) == 0 {
		agents = []Agent{
			ArtistsAgent{ArtistID: artistID},
			LocationAgent{ArtistID: artistID},
		}
	}

	return agents
}
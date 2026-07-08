package main

import "strings"

type Orchestrator struct{}

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

	// More specific checks
	wantsDates := strings.Contains(lower, "when") ||
		strings.Contains(lower, "date") ||
		strings.Contains(lower, "tour date") ||
		strings.Contains(lower, "concert date")

	wantsLocation := strings.Contains(lower, "where") ||
		strings.Contains(lower, "location") ||
		strings.Contains(lower, "concert location") ||
		strings.Contains(lower, "performed") // past tense, not "perform"

	// Add this to the existing wants conditions
	wantsRelation := strings.Contains(lower, "relation") ||
		strings.Contains(lower, "link") ||
		strings.Contains(lower, "connect") ||
		strings.Contains(lower, "all data") ||
		strings.Contains(lower, "full details") ||
		strings.Contains(lower, "tour schedule")

	if wantsRelation {
		agents = append(agents, RelationAgent{ArtistID: artistID})
	}

	if wantsArtist {
		agents = append(agents, ArtistsAgent{ArtistID: artistID})
	}

	if wantsLocation {
		agents = append(agents, LocationAgent{ArtistID: artistID})
	}

	if wantsDates {
		agents = append(agents, DatesAgent{ArtistID: artistID})
	}

	if len(agents) == 0 {
		agents = []Agent{
			ArtistsAgent{ArtistID: artistID},
			LocationAgent{ArtistID: artistID},
			DatesAgent{ArtistID: artistID},
		}
	}

	return agents
}

package main

import (
	"fmt"
	"sync"
)

type Result struct {
	AgentName string
	Data interface{}
	Error error
}

type Agent interface {
	Name() string
	Execute() Result
}

type ArtistAgent struct {
	Artist string
}

func (a ArtistAgent) Name() string {
	return "ArtistAgent"
}

func (a ArtistAgent) Execute() Result {
	return Result{
		AgentName: a.Name(),
		Data: fmt.Sprint("Artist: %s", a.Artist),
		Error: nil,
	}
}

type WeatherAgent struct {
	Location string
}

func (w WeatherAgent) Name() string {
	return "WeatherAgent"
}

func (w WeatherAgent) Execute() Result {
	return Result{
		AgentName: w.Name(),
		Data: fmt.Sprintf("Weather in %s: 22°C", w.Location),
		Error: nil,
	}
}

func main(){
	agents := []Agent{
		ArtistAgent{Artist: "Burna Boy"},
		WeatherAgent{Location: "Lagos"},
		ArtistAgent{Artist: "Wizkid"},
	}

	var wg sync.WaitGroup

	results := make(chan Result, len(agents))

	for _, agent := range agents {
		wg.Add(1)

		go func(a Agent) {
			defer wg.Done()
			results <- a.Execute()
		}(agent)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Printf("[%s] %v\n", result.AgentName, result.Data)
	}
}
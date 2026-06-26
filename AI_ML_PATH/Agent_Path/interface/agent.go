package main

import (
    "fmt"
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

// ArtistAgent
type ArtistAgent struct {
    Artist string
}

func (a ArtistAgent) Name() string {
    return "ArtistAgent"
}

func (a ArtistAgent) Execute() Result {
    return Result{
        AgentName: a.Name(),
        Data:      fmt.Sprintf("Artist: %s", a.Artist),
        Error:     nil,
    }
}

// WeatherAgent
type WeatherAgent struct {
    Location string
}

func (w WeatherAgent) Name() string {
    return "WeatherAgent"
}

func (w WeatherAgent) Execute() Result {
    return Result{
        AgentName: w.Name(),
        Data:      fmt.Sprintf("Weather in %s: 22°C", w.Location),
        Error:     nil,
    }
}

func main() {
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

/*package main

import (
	"fmt"
	"sync"
)

type Agent interface {
	Name() string
	Execute() string
}

type ArtistAgent struct {
	Artist string
}

func (a ArtistAgent) Name() string {
	return "ArtistAgent"
}

func (a ArtistAgent) Execute() string {
	return "Fetching artist: " + a.Artist
}

type WeatherAgent struct {
	Location string
}

func (w WeatherAgent) Name() string {
	return "WeatherAgent"
}

func (w WeatherAgent) Execute() string {
	return "Weather in " + w.Location + ": 22°C"
}

func RunAgent(a Agent) string {
	return a.Execute()
}

func main(){
	agents := []Agent{
		ArtistAgent{Artist: "Burna Boy"},
		WeatherAgent{Location: "Lagos"},
		ArtistAgent{Artist: "Wizkid"},
	}

	var wg sync.WaitGroup

	results := make(chan string, len(agents))

	for _, agent := range agents {
		wg.Add(1)
		go func(a Agent) {
			defer wg.Done()
			result := RunAgent(a)
			results <- result
		}(agent)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}
}*/
package main

import "sync"

type Coordinator struct {
	agents []Agent
}

func Newcoordinator(agents []Agent) *Coordinator {
	return &Coordinator{agents: agents}
}

func (c *Coordinator) Run() []Result {
	var wg sync.WaitGroup

	results := make(chan Result, len(c.agents))

	for _, agent := range c.agents {
		wg.Add(1)
		go func(a Agent) {
			defer wg.Done()
			results <- a.Execute()
		}(agent)
	}

	go func(){
		wg.Wait()
		close(results)
	}()

	var allResults []Result
	for result := range results {
		allResults = append(allResults, result)
	}

	return allResults
}
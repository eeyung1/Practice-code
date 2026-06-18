package main

import "fmt"

type Artist struct {
	ID   int
	Name string
}

type Location struct {
	ID        int
	Locations []string
}

type ArtistInfo struct {
	Artist
	Location
}

func main() {
	var infos []ArtistInfo

	infos = append(infos,
		ArtistInfo{
			Artist: Artist{
				ID:   1,
				Name: "Queen",
			},
			Location: Location{
				ID:        1,
				Locations: []string{"England", "Germany"},
			},
		},
	)

	infos = append(infos,
		ArtistInfo{
			Artist: Artist{
				ID:   2,
				Name: "ABBA",
			},
			Location: Location{
				ID:        2,
				Locations: []string{"Sweden"},
			},
		},
	)

	// fmt.Println(infos[0])
	// fmt.Println(infos)

	for _, info := range infos {
		fmt.Println(info.Name)

		for _, location := range info.Locations {
			fmt.Println("-", location)
		}

		fmt.Println()
	}

}

package main

import (
	"fmt"
	"log"
	"time"

	"blaus/pkg/http"
	"blaus/pkg/lib/csv"
	"blaus/pkg/lib/txt"
)

const (
	QuantityToFetchRepositories = 1020
)

func main() {
	requester, err := http.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	repositories, err := requester.QueryRepos(QuantityToFetchRepositories)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Took %s to query %d repositories\n", time.Since(now), len(repositories))
	err = txt.Save(repositories)
	if err != nil {
		fmt.Println(err)
	}

	now = time.Now()
	err = csv.Save(repositories)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Took %s to csv %d repositories\n", time.Since(now), len(repositories))
}

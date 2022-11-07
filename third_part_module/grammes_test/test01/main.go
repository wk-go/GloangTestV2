package main

import (
	"log"

	"github.com/northwesternmutual/grammes"
)

func main() {
	// Creates a new client with the localhost IP.
	client, err := grammes.DialWithWebSocket("ws://192.168.3.150:8345")
	if err != nil {
		log.Fatalf("Error while creating client: %s\n", err.Error())
	}

	// Executing a basic query to assure that the client is working.
	res, err := client.ExecuteStringQuery("1+3")
	if err != nil {
		log.Fatalf("Querying error: %s\n", err.Error())
	}

	// Print out the result as a string
	for _, r := range res {
		log.Println(string(r))
	}

	// Executing a query to add a vertex to the graph.
	client.AddVertex("testingvertex")

	// Create a new traversal string to build your traverser.
	g := grammes.Traversal()

	// Executing a query to fetch all of the labels from the vertices.
	res, err = client.ExecuteQuery(g.V().Label())
	if err != nil {
		log.Fatalf("Querying error: %s\n", err.Error())
	}

	// Log out the response.
	for _, r := range res {
		log.Println(string(r))
	}
}

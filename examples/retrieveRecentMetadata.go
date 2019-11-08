package main

import (
	"github.com/j4ng5y/onetimesecret-go"
	"log"
)

func main() {
	client := onetimesecret.New(&onetimesecret.Credentials{
		Username: "jordan@example.com", // Required
		APIToken: "abcdefg1234567", // Required
	})

	retrieveRecentMetataRequest := &onetimesecret.RetrieveRecentMetadataRequest{}

	retrieveRecentMetataResponse, err := client.RetrieveRecentMetadata(retrieveRecentMetataRequest)
	if err != nil {
		log.Print(err)
	}
	log.Print(retrieveRecentMetataResponse)
}
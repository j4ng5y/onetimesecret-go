package main

import (
	"github.com/j4ng5y/onetimesecret-go"
	"log"
)

func main() {
	client := onetimesecret.New(&onetimesecret.Credentials{
		Username: "jordan@example.com", // Required
		APIToken: "abcdefg1234567",     // Required
	})

	retrieveMetadataRequest := &onetimesecret.RetrieveMetadataRequest{
		MetadataKey: "abcdefg12345", // Required: the key of the metadata to retrieve
	}

	retrieveMetadataResponse, err := client.RetrieveMetadata(retrieveMetadataRequest)
	if err != nil {
		log.Print(err)
	}
	log.Print(retrieveMetadataResponse)
}

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

	burnRequest := &onetimesecret.BurnSecretRequest{
		MetadataKey: "abcdefg12345", // Required: The metadata key of the secret to burn
	}
	burnResponse, err := client.BurnSecret(burnRequest)
	if err != nil {
		log.Print(err)
	}
	log.Print(burnResponse)
}

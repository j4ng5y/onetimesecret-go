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

	retrieveSecretRequest := &onetimesecret.RetrieveSecretRequest{
		SecretKey:  "abcdefg12345", // Required: the key of the secret to retrieve
		Passphrase: "",             // Optional: Only required if a passphrase was set when creating the secret
	}

	retrieveSecretResponse, err := client.RetrieveSecret(retrieveSecretRequest)
	if err != nil {
		log.Print(err)
	}
	log.Print(retrieveSecretResponse)
}

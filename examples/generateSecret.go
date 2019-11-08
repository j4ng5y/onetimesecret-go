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

	generateRequest := &onetimesecret.GenerateSecretRequest{
		Passphrase: "", // Optional: Set a passphrase to decrypt the secret
		TTL:        0,  // Optional: Set the secret lifetime
		Recipient:  "", // Optional: Set the recipients email
	}

	generateResponse, err := client.GenerateSecret(generateRequest)
	if err != nil {
		log.Print(err)
	}
	log.Print(generateResponse)
}

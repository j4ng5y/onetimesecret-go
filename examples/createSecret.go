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

	createRequest := &onetimesecret.CreateSecretRequest{
		Secret:     "abcdefg12345", // This is the only required field
		Passphrase: "",             // Optionally: Set this to require a passphrase to decrypt
		TTL:        0,              // Optionally: Set the time the secret lives
		Recipient:  "",             // Optionally: Set the recipients email
	}

	createResponse, err := client.CreateSecret(createRequest)
	if err != nil {
		log.Print(err)
	}
	log.Print(createResponse)
}

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func decodeJWT(tokenString string) (map[string]interface{}, error) {
	// Split the token into 3 parts: Header, Payload, and Signature
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	// Decode the Payload part (Base64URL decoding)
	payload, err := decodeBase64URL(parts[1])
	if err != nil {
		return nil, fmt.Errorf("error decoding payload: %v", err)
	}

	// Parse the Payload into a map
	var jsonPayload map[string]interface{}
	err = json.Unmarshal(payload, &jsonPayload)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling payload: %v", err)
	}

	return jsonPayload, nil
}

// Function to decode Base64URL into bytes
func decodeBase64URL(base64URL string) ([]byte, error) {
	// Convert Base64URL to standard Base64
	base64String := base64URL
	base64String = strings.ReplaceAll(base64String, "-", "+")
	base64String = strings.ReplaceAll(base64String, "_", "/")

	// Ensure the Base64 string length is a multiple of 4
	padding := len(base64String) % 4
	if padding > 0 {
		base64String += strings.Repeat("=", 4-padding)
	}

	// Decode the standard Base64 string
	return base64.StdEncoding.DecodeString(base64String)
}

func main() {
	// Check if the JWT token is provided via command-line arguments
	if len(os.Args) < 2 {
		log.Fatal("JWT token must be provided as an argument.")
	}

	// Retrieve the JWT token from the command-line arguments
	tokenString := os.Args[1]

	// Decode the token
	payload, err := decodeJWT(tokenString)
	if err != nil {
		log.Fatalf("Error decoding JWT: %v", err)
	}

	// Output the decoding result in JSON format
	output, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling payload to JSON: %v", err)
	}

	// Display the result in JSON format
	fmt.Println(string(output))
}

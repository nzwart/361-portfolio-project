package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"zwartn-microservice-A/types"
)

// Constant that defines parameters and char sets for pwd generation
const (
	minLength = 12
	maxLength = 64
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	numbers   = "0123456789"
	special   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

func generatePassword(length int) (string, error) {
	if length < minLength {
		return "", fmt.Errorf("password length must be at least %d characters", minLength)
	}
	if length > maxLength {
		return "", fmt.Errorf("password length must not exceed %d characters", maxLength)
	}

	allChars := uppercase + lowercase + numbers + special
	var password strings.Builder

	password.WriteByte(getRandomChar(uppercase))
	password.WriteByte(getRandomChar(lowercase))
	password.WriteByte(getRandomChar(numbers))
	password.WriteByte(getRandomChar(special))

	for i := 4; i < length; i++ {
		password.WriteByte(getRandomChar(allChars))
	}

	passwordRunes := []rune(password.String())
	for i := len(passwordRunes) - 1; i > 0; i-- {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		passwordRunes[i], passwordRunes[j.Int64()] = passwordRunes[j.Int64()], passwordRunes[i]
	}

	return string(passwordRunes), nil
}

func getRandomChar(charset string) byte {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		panic(err)
	}
	return charset[n.Int64()]
}

func analyzePassword(password string) types.CompositionStats {
	stats := types.CompositionStats{
		TotalLength: len(password),
	}

	for _, char := range password {
		switch {
		case strings.ContainsRune(uppercase, char):
			stats.UppercaseCount++
		case strings.ContainsRune(lowercase, char):
			stats.LowercaseCount++
		case strings.ContainsRune(numbers, char):
			stats.NumberCount++
		case strings.ContainsRune(special, char):
			stats.SpecialCharCount++
		}
	}

	return stats
}

func calculateSecurityScore(stats types.CompositionStats) types.SecurityScore {
	score := 0

	if stats.TotalLength >= minLength {
		score++
	}
	if stats.UppercaseCount > 0 && stats.LowercaseCount > 0 {
		score++
	}
	if stats.NumberCount > 0 {
		score++
	}
	if stats.SpecialCharCount > 0 {
		score++
	}
	if stats.TotalLength >= 16 {
		score++
	}

	evaluation := "Weak"
	if score >= 4 {
		evaluation = "Very Strong"
	} else if score >= 3 {
		evaluation = "Strong"
	} else if score >= 2 {
		evaluation = "Moderate"
	}

	return types.SecurityScore{
		Score:      score,
		Evaluation: evaluation,
	}
}

func handlePasswordGeneration(w http.ResponseWriter, r *http.Request) {
	// Print statements for video feedback
	fmt.Printf("\n[Server] Received request for password generation")
	fmt.Printf("[Server] Requested length parameter: %s\n", r.URL.Query().Get("length"))

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	length := minLength
	if lengthStr := r.URL.Query().Get("length"); lengthStr != "" {
		requestedLength, err := strconv.Atoi(lengthStr)
		if err != nil {
			http.Error(w, "Invalid length parameter", http.StatusBadRequest)
			return
		}
		length = requestedLength
	}

	password, err := generatePassword(length)
	if err != nil {
		fmt.Printf("[Server] Error generating password: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("[Server] Generated password of length %d\n", len(password))

	analysis := analyzePassword(password)
	security := calculateSecurityScore(analysis)

	fmt.Printf("[Server] Password analysis complete - Security score: %d (%s)\n",
		security.Score, security.Evaluation)

	response := types.PasswordPackage{
		Password: password,
		Analysis: analysis,
		Security: security,
		Metadata: types.Metadata{
			Standard:   "NIST SP 800-63B",
			Compliance: "Compliant",
			Version:    "1.0.0",
		},
	}

	fmt.Printf("[Server] Sending response to client\n")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("[Server] Password generation service initializing...")
	http.HandleFunc("/generate", handlePasswordGeneration)

	fmt.Println("[Server] Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

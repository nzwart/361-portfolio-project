package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"zwartn-microservice-A/types"
)

func main() {
	// Print statement for video feedback
	fmt.Println("[Client] Password generation test client starting...")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	lengths := []int{12, 16, 20}

	for _, length := range lengths {
		fmt.Printf("\n[Client] Requesting password with length %d\n", length)
		fmt.Printf("[Client] Sending request to http://localhost:8080/generate?length=%d\n", length)

		resp, err := client.Get(fmt.Sprintf("http://localhost:8080/generate?length=%d", length))
		if err != nil {
			fmt.Printf("[Client] Error making request: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("[Client] Received response with status: %s\n", resp.Status)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("[Client] Error reading response: %v\n", err)
			continue
		}

		var passwordPackage types.PasswordPackage
		if err := json.Unmarshal(body, &passwordPackage); err != nil {
			fmt.Printf("[Client] Error parsing JSON: %v\n", err)
			continue
		}

		fmt.Printf("[Client] Successfully received and parsed password package:\n")
		fmt.Printf("  - Password: %s\n", passwordPackage.Password)
		fmt.Printf("  - Length: %d\n", passwordPackage.Analysis.TotalLength)
		fmt.Printf("  - Security Score: %d (%s)\n",
			passwordPackage.Security.Score,
			passwordPackage.Security.Evaluation)
	}
}

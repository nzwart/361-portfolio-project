package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/pebbe/zmq4"
)

//
// Environmental effect microservice
//

func environmentHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")

	log.Printf("Request to environment API - Method: %s, URL: %s, Location: %s\n",
		r.Method, r.URL, location)

	url := fmt.Sprintf("http://localhost:8082/generate?location=%s", location)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to connect to environment service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from environment service", http.StatusInternalServerError)
		return
	}

	log.Printf("Environment API response: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

//
// Combat tactics microservice
//

func tacticsHandler(w http.ResponseWriter, r *http.Request) {
	aggression := r.URL.Query().Get("aggression")
	isGroup := r.URL.Query().Get("group")

	log.Printf("Request to tactics API - Method: %s, URL: %s, Aggression: %s, Group: %s\n",
		r.Method, r.URL, aggression, isGroup)

	url := fmt.Sprintf("http://localhost:8083/generate?aggression=%s&group=%s", aggression, isGroup)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to connect to tactics service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from tactics service", http.StatusInternalServerError)
		return
	}

	log.Printf("Tactics API response: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

//
// Plot twist microservice
//

type PlotTwist struct {
	Category string `json:"category"`
	Text     string `json:"text"`
}

type PlotTwistResponse struct {
	Success bool      `json:"success"`
	Error   string    `json:"error,omitempty"`
	Data    PlotTwist `json:"data,omitempty"`
}

func plotTwistHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	if category == "" {
		category = "combat"
	}

	log.Printf("Request to plot twist API - Method: %s, URL: %s, Category: %s\n", r.Method, r.URL, category)

	url := fmt.Sprintf("http://localhost:8081/generate?category=%s", category)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to connect to plot twist service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response from plot twist service", http.StatusInternalServerError)
		return
	}

	log.Printf("Plot twist API response: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
		return
	}

	var plotResp PlotTwistResponse
	if err := json.Unmarshal(body, &plotResp); err != nil {
		http.Error(w, "Failed to parse plot twist service response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plotResp)
}

//
// Name gen microservice
//

// NameRequest represents the request structure for name generation
type NameRequest struct {
	NameNum  *int   `json:"name_num,omitempty"`
	Race     string `json:"race,omitempty"`
	AddTitle bool   `json:"add_title,omitempty"`
}

func nameGeneratorHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request to name generator API - Method: %s, URL: %s\n", r.Method, r.URL)

	socket, err := zmq4.NewSocket(zmq4.REQ)
	if err != nil {
		http.Error(w, "Failed to create socket", http.StatusInternalServerError)
		return
	}
	defer socket.Close()

	err = socket.Connect("tcp://localhost:5324")
	if err != nil {
		http.Error(w, "Failed to connect to name service", http.StatusInternalServerError)
		return
	}

	request := map[string]interface{}{}
	requestJSON, _ := json.Marshal(request)
	_, err = socket.SendBytes(requestJSON, 0)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}

	response, err := socket.Recv(0)
	if err != nil {
		http.Error(w, "Failed to receive response", http.StatusInternalServerError)
		return
	}

	log.Printf("Name generator API response received: %s\n", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"name": response})
}

// Middleware to add CORS headers to all responses
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type Message struct {
	Page    string `json:"page"`
	Content string `json:"content"`
}

func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Page: "Frontpage", Content: "Welcome to the Frontpage!"}
	json.NewEncoder(w).Encode(response)
}

func encBuilderHandler(w http.ResponseWriter, r *http.Request) {
	response := Message{Page: "Encbuilder", Content: "This is the Encbuilder page."}
	json.NewEncoder(w).Encode(response)
}

func monstersHandler(w http.ResponseWriter, r *http.Request) {
	// Open the monsters.json file
	file, err := os.Open("5e-SRD-Monsters.json")
	if err != nil {
		http.Error(w, "Unable to read monster data", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Decode the JSON into a slice of maps
	var rawMonsters []map[string]interface{}
	if err := json.NewDecoder(file).Decode(&rawMonsters); err != nil {
		panic(err)
	}

	// Convert each monster to the flat Monster struct
	monsters := make([]Monster, len(rawMonsters))
	for i, data := range rawMonsters {
		monsters[i] = parseMonster(data)
	}

	// Respond with the JSON data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monsters)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/frontpage", frontPageHandler)
	mux.HandleFunc("/api/encbuilder", encBuilderHandler)
	mux.HandleFunc("/api/monsters", monstersHandler)
	mux.HandleFunc("/api/generate-name", nameGeneratorHandler)
	mux.HandleFunc("/api/plot-twist", plotTwistHandler)
	mux.HandleFunc("/api/tactics", tacticsHandler)
	mux.HandleFunc("/api/environment", environmentHandler)

	// Wrap all handlers with the CORS middleware
	http.ListenAndServe(":8080", enableCORS(mux))
}

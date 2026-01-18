package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strings"
)

type categories struct {
	Category []string `json:"category"`
}

type messages struct {
	Category []string `json:"category"`
	Message  []string `json:"messages"`
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	mux := http.NewServeMux()

	mux.HandleFunc("/no/", handlerFunc)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(server.ListenAndServe())

}

func handlerFunc(w http.ResponseWriter, r *http.Request) {

	var jsonResponse map[string]string

	category := strings.TrimPrefix(r.URL.Path, "/no/")

	if category == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonResponse := map[string]any{
			"categories": []string{
				"random",
				"corporate_speak",
				"friendly_firm",
				"funny_light",
				"over_dramatic",
				"polite_professional",
				"passive_aggressive",
				"sarcastic",
				"tech_nerd",
			},
		}
		json.NewEncoder(w).Encode(jsonResponse)
	} else {
		switch r.Method {
		case http.MethodGet:
			response, err := responseReader(category)
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				jsonResponse = map[string]string{
					"error":   "invalid_category",
					"message": "Category is not supported",
				}
			} else {
				w.WriteHeader(http.StatusOK)
				jsonResponse = map[string]string{
					"message": response,
				}
			}

			json.NewEncoder(w).Encode(jsonResponse)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}

}

func responseReader(reqCategory string) (string, error) {

	var categories categories
	var messages messages
	var category string
	var response string

	allowed_categories := map[string]string{
		"corporate_speak":     "corporate_speak.json",
		"friendly_firm":       "friendly_firm.json",
		"funny_light":         "funny_light.json",
		"over_dramatic":       "over_dramatic.json",
		"polite_professional": "polite_professional.json",
		"passive_aggressive":  "passive_aggressive.json",
		"sarcastic":           "sarcastic.json",
		"tech_nerd":           "tech_nerd.json",
	}

	if reqCategory == "random" {
		data, err := os.ReadFile("./data/category.json")
		if err != nil {
			fmt.Println("Unable to read file")
		}

		json.Unmarshal(data, &categories)
		number := rand.IntN(len(categories.Category))
		category = categories.Category[number]

	} else {
		category = reqCategory
	}

	data, err := os.ReadFile("./data/" + allowed_categories[category])
	if err != nil {
		fmt.Println("Catergory Unavailable")
		return "", fmt.Errorf("Invalid Category")
	}

	json.Unmarshal(data, &messages)
	number := rand.IntN(len(messages.Message))
	response = messages.Message[number]

	return response, nil
}

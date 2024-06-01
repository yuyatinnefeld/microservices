package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	AppName    string `json:"appName"`
	Language   string `json:"language"`
	Version    string `json:"version"`
	Message    string `json:"message"`
	Enviroment string `json:"env"`
	PodID      string `json:"podID"`
}

func setEnvOrDefault(envName string, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}
	return defaultValue
}

func fetchAPIResource(w http.ResponseWriter, r *http.Request) {
	appName := "payment"
	language := "golang"

	version := setEnvOrDefault("VERSION", "0.0.0")
	message := setEnvOrDefault("MESSAGE", "MESSAGE_NOT_DEFINED")
	env := setEnvOrDefault("ENV", "ENV_NOT_DEFINED")
	podID := setEnvOrDefault("MY_POD_NAME", "PODID_NOT_DEFINED")

	response := Response{
		AppName:  appName,
		Language: language,
		Version:  version,
		Message:  message,
		Enviroment: env,
		PodID:    podID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func createAPIResource(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Here you can process the received data as needed
		fmt.Printf("Received data: %+v\n", data)

		response := struct {
			Message string `json:"message"`
		}{
			Message: "Data received successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
    
func main() {
	http.HandleFunc("/", fetchAPIResource)
	http.HandleFunc("/post", createAPIResource)
	port := 8888
	fmt.Printf("Serving on port %d..\n", port)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	AppName   string `json:"appName"`
	Language  string `json:"language"`
	Version   string `json:"version"`
	Message   string `json:"message"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func setEnvOrDefault(envName string, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}
	return defaultValue
}

func fetchAPIResource(w http.ResponseWriter, r *http.Request) {
	log.Println("###### 🚀 START APPLICATION 🚀 ###### ")

	appName := "vault"
	language := "golang"
	version := setEnvOrDefault("VERSION", "0.0.0")
	message := setEnvOrDefault("MESSAGE", "MESSAGE_NOT_DEFINED")
	vURL := setEnvOrDefault("VAULT_ADDR", "ENV_NOT_DEFINED")
	vToken := setEnvOrDefault("VAULT_TOKEN", "PODID_NOT_DEFINED")

	log.Println("###### VALIDATE ENV VARIABLES ######")
	log.Println("appName: ", appName)
	log.Println("language: ", language)
	log.Println("version: ", version)
	log.Println("message: ", message)
	log.Println("vURL: ", vURL)
	log.Println("vToken: ", vToken)

	log.Println("###### CREATE VAULT REQUEST ###### ")
	secretsPath := "/v1/secret/data/yuya_password/config"
	log.Println("vToken: ", secretsPath)

	url := fmt.Sprintf("%s%s", vURL, secretsPath)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Vault-Token", vToken)

	log.Println("###### FETCH VAULT RESPONSE ###### ")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	log.Println("###### VALIDATE RESPONSE STATUS ###### ")

	if resp.StatusCode != http.StatusOK {
		log.Println("ERROR: StatusCode = ", resp.StatusCode)
	}else{
		log.Println("SUCCESS: StatusCode = ", resp.StatusCode)
	}

	log.Println("###### EXTRACT SECRETS FROM RESPONSE ###### ")
	var content map[string]interface{}

	log.Println(resp.Body)

	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	raw_data := content["data"].(map[string]interface{})
	log.Println("RAW DATA: ", raw_data)
	vSecret, ok := raw_data["data"].(map[string]interface{})

	username, ok := vSecret["username"].(string)
	if !ok {
		log.Println("Username not found or not a string")
		return
	}

	password, ok := vSecret["password"].(string)
	if !ok {
		log.Println("Password not found or not a string")
		return
	}

	log.Println("Username: ", username)
	log.Println("Password: ", password)

	response := Response{
		AppName:  appName,
		Language: language,
		Version:  version,
		Message:  message,
		Username: username,
		Password: password,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Println("###### 🍻 FINISH APPLICATION 🍻 ###### ")

}

func main() {
	http.HandleFunc("/", fetchAPIResource)
	port := 8899
	fmt.Printf("Serving on port %d..\n", port)
	log.Fatal(http.ListenAndServe(":8899", nil))
}
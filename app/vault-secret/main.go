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

func retrieveVaultSecrets(vURL, vToken, secretsPath string) (string, string, error) {

	url := fmt.Sprintf("%s%s", vURL, secretsPath)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Vault-Token", vToken)
	log.Println("✅ CREATE VAULT REQUEST ✅")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()
	log.Println("✅ FETCH VAULT RESPONSE ✅")

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to fetch secrets: status code %d", resp.StatusCode)
	}

	log.Println("✅ VALIDATE RESPONSE STATUS ✅")

	var content map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		return "", "", err
	}

	rawData, ok := content["data"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("invalid response structure")
	}

	vSecret, ok := rawData["data"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("invalid secret structure")
	}

	username, ok := vSecret["username"].(string)
	if !ok {
		return "", "", fmt.Errorf("username not found or not a string")
	}

	password, ok := vSecret["password"].(string)
	if !ok {
		return "", "", fmt.Errorf("password not found or not a string")
	}

	log.Println("✅ EXTRACT SECRETS FROM RESPONSE ✅")

	return username, password, nil
}

func fetchAPIResource(w http.ResponseWriter, r *http.Request) {
	log.Println("###### 🚀 START APPLICATION 🚀 ###### ")

	appName := "vault"
	language := "golang"
	version := setEnvOrDefault("VERSION", "0.0.0")
	message := setEnvOrDefault("MESSAGE", "MESSAGE_NOT_DEFINED")
	vURL := setEnvOrDefault("VAULT_ADDR", "http://192.168.64.1:8200")
	vToken := setEnvOrDefault("VAULT_TOKEN", "root")
	secretsPath := setEnvOrDefault("SECRETS_PATH", "/v1/secret/data/yuya_password/config")

	log.Println("✅ VALIDATE ENV VARIABLES ✅")
	log.Println("appName: ", appName)
	log.Println("language: ", language)
	log.Println("version: ", version)
	log.Println("message: ", message)
	log.Println("vURL: ", vURL)
	log.Println("vToken: ", vToken)
	log.Println("secretsPath: ", secretsPath)

	username, password, err := retrieveVaultSecrets(vURL, vToken, secretsPath)

	if err != nil {
		log.Println("❌ Error ❌")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

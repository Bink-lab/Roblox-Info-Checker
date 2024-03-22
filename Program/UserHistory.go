package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// UserHistoryResponse represents the structure of the JSON response from the API
type UserHistoryResponse struct {
	Data []struct {
		Name string `json:"name"`
	} `json:"data"`
}

func main() {
	// Ask the user for a user ID
	fmt.Print("Enter the user ID: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userID := scanner.Text()

	// Make a GET request to the Roblox API endpoint
	url := fmt.Sprintf("https://users.roblox.com/v1/users/%s/username-history?limit=10&sortOrder=Asc", userID)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %s\n", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	// Parse JSON response
	var userHistory UserHistoryResponse
	if err := json.Unmarshal(body, &userHistory); err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return
	}

	// Print the usernames if available
	if len(userHistory.Data) > 0 {
		fmt.Println("Name/s:")
		for _, userData := range userHistory.Data {
			fmt.Println(userData.Name)
		}
	} else {
		fmt.Println("No old usernames found.")
	}

	// Prompt the user to press Enter to exit
	fmt.Println("\nPress Enter to exit...")
	bufio.NewScanner(os.Stdin).Scan()
}

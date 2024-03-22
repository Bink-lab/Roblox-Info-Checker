package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Hello Kitty watermark ;)
func watermark() {
	fmt.Println(`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢠⡾⠲⠶⣤⣀⣠⣤⣤⣤⡿⠛⠿⡴⠾⠛⢻⡆⠀⠀⠀
⠀⠀⠀⣼⠁⠀⠀⠀⠉⠁⠀⢀⣿⠐⡿⣿⠿⣶⣤⣤⣷⡀⠀⠀
⠀⠀⠀⢹⡶⠀⠀⠀⠀⠀⠀⠈⢯⣡⣿⣿⣀⣸⣿⣦⢓⡟⠀⠀
⠀⠀⢀⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠹⣍⣭⣾⠁⠀⠀
⠀⣀⣸⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣸⣷⣤⡀
⠈⠉⠹⣏⡁⠀⢸⣿⠀⠀⠀⢀⡀⠀⠀⠀⣿⠆⠀⢀⣸⣇⣀⠀
⠀⠐⠋⢻⣅⣄⢀⣀⣀⡀⠀⠯⠽⠂⢀⣀⣀⡀⠀⣤⣿⠀⠉⠀
⠀⠀⠴⠛⠙⣳⠋⠉⠉⠙⣆⠀⠀⢰⡟⠉⠈⠙⢷⠟⠉⠙⠂⠀
⠀⠀⠀⠀⠀⢻⣄⣠⣤⣴⠟⠛⠛⠛⢧⣤⣤⣀⡾⠁⠀⠀⠀⠀`)
}

// UserInfo represents the structure of the JSON response from the API
type UserInfo struct {
	Name            string `json:"name"`
	DisplayName     string `json:"displayName"`
	Description     string `json:"description"`
	Created         string `json:"created"`
	IsBanned        bool   `json:"isBanned"`
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"`
	ID              int    `json:"id"`
}

// UserHistoryResponse represents the structure of the JSON response from the old usernames API
type UserHistoryResponse struct {
	Data []struct {
		Name string `json:"name"`
	} `json:"data"`
}

func main() {
	watermark()

	// Ask the user for a user ID
	fmt.Print("\n  Enter the user ID: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userID := scanner.Text()

	// Make a GET request to the Roblox API endpoint to fetch user info
	url := fmt.Sprintf("https://users.roblox.com/v1/users/%s", userID)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %s\n", err)
		return
	}
	defer response.Body.Close()

	// Check if the response is successful
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", response.Status)
		return
	}

	// Decode the response body to get user info
	var userInfo UserInfo
	if err := json.NewDecoder(response.Body).Decode(&userInfo); err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
		return
	}

	// Print the user info
	fmt.Println("  User Info:")
	fmt.Printf("  Name: %s\n", userInfo.Name)
	fmt.Printf("  Display Name: %s\n", userInfo.DisplayName)
	fmt.Printf("  Description: %s\n", userInfo.Description)
	fmt.Printf("  Created: %s\n", userInfo.Created)
	fmt.Printf("  Is Banned: %t\n", userInfo.IsBanned)
	fmt.Printf("  Has Verified Badge: %t\n", userInfo.HasVerifiedBadge)
	fmt.Printf("  ID: %d\n", userInfo.ID)

	// Make a GET request to fetch old usernames
	historyURL := fmt.Sprintf("https://users.roblox.com/v1/users/%s/username-history?limit=10&sortOrder=Asc", userID)
	historyResponse, err := http.Get(historyURL)
	if err != nil {
		fmt.Printf("Error making GET request for username history: %s\n", err)
		return
	}
	defer historyResponse.Body.Close()

	// Decode the response body to get old usernames
	var userHistory UserHistoryResponse
	if err := json.NewDecoder(historyResponse.Body).Decode(&userHistory); err != nil {
		fmt.Printf("Error decoding JSON for username history: %s\n", err)
		return
	}

	// Print the old usernames
	fmt.Println("\n  Old Usernames:")
	for _, data := range userHistory.Data {
		fmt.Printf("  - %s\n", data.Name)
	}

	// Prompt to exit
	fmt.Println("\n  Press Enter to exit...")
	bufio.NewScanner(os.Stdin).Scan()
}

package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    // Ask the user to enter the URL and file name
    fmt.Print("Enter the URL: ")
    var url string
    fmt.Scanln(&url)
    fmt.Print("Enter the file name: ")
    var extensie string
    fmt.Scanln(&extensie)
    
    // Check if the URL is valid
    if _, err := http.ParseURL(url); err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Send HTTP GET request to the URL
    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err) 
        return
    }
    defer response.Body.Close()

    // Create a new file to save the downloaded content
    outputFile, err := os.Create(extensie + ".txt") // Append ".txt" extension to the file name
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer outputFile.Close()

    // Copy the response body to the file
    _, err = io.Copy(outputFile, response.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("File downloaded successfully!")
}

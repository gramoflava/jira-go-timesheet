package main

import (
    "encoding/json"
    "fmt"
    "os"
//    "github.com/andygrunwald/go-jira"
//    "log"
//    "os"
//    "strings"
)

type Configuration struct {
    Users    []string
    Groups   []string
}

func main() {
    fmt.Println("hello world")

    file, _ := os.Open("~/conf.json")
    defer file.Close()

    decoder := json.NewDecoder(file)

    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
        configuration.Users = ["test1"]
        configuration.Groups = ["test2"]
    }
    fmt.Println(configuration.Users) // output: [UserA, UserB]

    // Read credentials from input
    // Login to JIRA
    // Get worklogs
    // Format to CSV
}


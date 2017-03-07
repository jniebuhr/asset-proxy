package main

import (
    "os"
    "encoding/json"
    "fmt"
)

type ServerConfiguration struct {
    port int
}

type FilterConfiguration struct {
    name string
    options map[string]string
}

type Configuration struct {
    server ServerConfiguration
    filters map[string]FilterConfiguration
    filterMappings map[string][]string
}

func LoadConfiguration() Configuration {
    file, _ := os.Open("config/application.json")
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error: ", err)
    }

    return configuration
}
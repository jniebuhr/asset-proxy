package main

import (
    "encoding/json"
    "os"
    "fmt"
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/gorillamux"
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

func loadConfiguration() Configuration {
    file, _ := os.Open("config/application.json")
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error: ", err)
    }

    return configuration
}

func main() {
    config := loadConfiguration()
    server := iris.New()
    server.Adapt(
        iris.DevLogger(),
        gorillamux.New(),
    )
    server.Get("{path:.*}", func(ctx *iris.Context) {
        ctx.Writef(ctx.GetString("path"))
    })
    server.Listen(fmt.Sprintf(":%d", config.server.port))
}
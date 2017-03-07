package main

import (
    "fmt"
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/gorillamux"
    "github.com/spf13/viper"
    "github.com/jniebuhr/asset-proxy/filters"
)


func main() {
    SetupConfig()
    filters.Init()
    server := iris.New()
    server.Adapt(
        iris.DevLogger(),
        gorillamux.New(),
    )
    server.Get("{path:.*}", HandleRequest)
    server.Listen(fmt.Sprintf(":%d", viper.GetInt("server.port")))
}
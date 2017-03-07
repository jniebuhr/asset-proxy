package main

import (
    "fmt"
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/gorillamux"
)


func main() {
    config := LoadConfiguration()
    server := iris.New()
    server.Adapt(
        iris.DevLogger(),
        gorillamux.New(),
    )
    server.Get("{path:.*}", HandleRequest)
    server.Listen(fmt.Sprintf(":%d", config.server.port))
}
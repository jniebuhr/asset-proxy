package main

import (
    "github.com/spf13/viper"
    "fmt"
)


func SetupConfig() {
    viper.SetConfigName("asset-proxy")
    viper.AddConfigPath("/etc/asset-proxy")
    viper.AddConfigPath("$HOME/.asset-proxy")
    viper.AddConfigPath("config")

    viper.SetDefault("verbose", false)
    viper.SetDefault("server.port", 8080)
    viper.SetDefault("directories.cache", "data/cache")
    viper.SetDefault("directories.source", "data/source")
    viper.SetDefault("directories.meta", "data/meta")
    viper.SetDefault("alwaysFilter", false)

    err := viper.ReadInConfig() // Find and read the config file
    if err != nil { // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
}

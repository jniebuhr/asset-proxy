package filters

import "github.com/spf13/viper"

type Filter interface {
    Initialize()
    Process(content []byte) []byte
}

var filters map[string]Filter = map[string]Filter{}
var filterMappings map[string][]string

func Init() {
    filterDefinitions := viper.GetStringMap("filters")
    for key, definition := range(filterDefinitions) {
        def := definition.(map[string]interface{})
        filter := instantiateFilter(def["name"].(string))
        filter.Initialize()
        filters[key] = filter
    }
    filterMappings = viper.GetStringMapStringSlice("filterMappings")
}

func instantiateFilter(name string) Filter {
    if name == "JsMinify" {
        return JsMinify{}
    } else if name == "CssMinify" {
        return CssMinify{}
    } else if name == "PngCompress" {
        return PngCompress{}
    }
    return nil
}

func Process(content []byte, mime string) []byte {
    filterKeys := filterMappings[mime]
    for _, key := range(filterKeys) {
        filter := filters[key]
        content = filter.Process(content)
    }
    return content
}

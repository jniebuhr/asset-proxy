package filters

import (
    "github.com/tdewolff/minify"
    "github.com/tdewolff/minify/js"
    "fmt"
)

type JsMinify struct {
}

var jsMinifier = minify.New()

func (f JsMinify) Initialize() {
    jsMinifier.AddFunc("application/javascript", js.Minify)
}

func (f JsMinify) Process(content []byte) []byte {
    retval, err := jsMinifier.Bytes("application/javascript", content)
    if err != nil {
        fmt.Println("Error while minifying css ", err);
        return content
    }
    return retval
}

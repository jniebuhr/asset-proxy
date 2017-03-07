package filters

import (
    "github.com/tdewolff/minify"
    "github.com/tdewolff/minify/css"
    "fmt"
)

type CssMinify struct {
}

var cssMinifier = minify.New()

func (f CssMinify) Initialize() {
    cssMinifier.AddFunc("text/css", css.Minify)
}

func (f CssMinify) Process(content []byte) []byte {
    retval, err := cssMinifier.Bytes("text/css", content)
    if err != nil {
        fmt.Println("Error while minifying css ", err);
        return content
    }
    return retval
}

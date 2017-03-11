package filters

import (
    "github.com/foobaz/lossypng/lossypng"
    "image"
    "github.com/tdewolff/buffer"
    "fmt"
    "image/png"
    "bytes"
)

type PngCompress struct {
}

func (f PngCompress) Initialize() {
}

func (f PngCompress) Process(content []byte) []byte {
    in, _, err := image.Decode(buffer.NewReader(content))
    if err != nil {
        fmt.Println("couldn't decode image")
        return content
    }
    out := lossypng.Compress(in, lossypng.NoConversion, 10)
    outWriter := new(bytes.Buffer)
    png.Encode(outWriter, out)
    return outWriter.Bytes()
}

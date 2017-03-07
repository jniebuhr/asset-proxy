package main

import (
    "gopkg.in/kataras/iris.v6"
    "fmt"
    "os"
    "io"
    "net/http"
    "github.com/spf13/viper"
    "crypto/sha1"
    "io/ioutil"
    "path/filepath"
    "strings"
    "encoding/json"
    "github.com/jniebuhr/asset-proxy/filters"
)

type MetaData struct {
    Mime string `json:"mime"`
}

func HandleRequest(ctx *iris.Context) {
    url := absoluteUrl(ctx.Request.RequestURI)
    hash := calculateHash(url)
    if !sourceExists(hash) {
        downloadFile(url, hash)
    }
    meta := readMeta(hash)
    if !cacheExists(hash) || viper.GetBool("alwaysFilter") {
        writeCache(hash, filters.Process(sourceContent(hash), meta.Mime))
    }
    ctx.SetHeader("Content-Type", meta.Mime)
    ctx.Write(cacheContent(hash))
}

func sourceContent(hash string) []byte {
    content, _ := ioutil.ReadFile(sourcePath(hash))
    return content
}

func sourceExists(hash string) bool {
    _, err := os.Stat(sourcePath(hash))
    return err == nil
}

func cacheContent(hash string) []byte {
    content, _ := ioutil.ReadFile(cachePath(hash))
    return content
}

func cacheExists(hash string) bool {
    _, err := os.Stat(cachePath(hash))
    return err == nil
}

func writeCache(hash string, content []byte) {
    os.MkdirAll(cacheDir(hash), os.ModePerm)
    ioutil.WriteFile(cachePath(hash), content, os.ModePerm)
}

func hashDir(hash string) string {
    return filepath.Join(hash[0:1], hash[1:3])
}

func sourceDir(hash string) string {
    return filepath.Join(viper.GetString("directories.source"), hashDir(hash))
}

func sourcePath(hash string) string {
    return filepath.Join(sourceDir(hash), hash)
}

func cacheDir(hash string) string {
    return filepath.Join(viper.GetString("directories.cache"), hashDir(hash))
}

func cachePath(hash string) string {
    return filepath.Join(cacheDir(hash), hash)
}

func metaDir(hash string) string {
    return filepath.Join(viper.GetString("directories.meta"), hashDir(hash))
}

func metaPath(hash string) string {
    return filepath.Join(metaDir(hash), hash)
}

func absoluteUrl(url string) string {
	return viper.GetString("baseUrl") + url
}

func calculateHash(url string) string {
    digest := sha1.New()
    digest.Write([]byte(url))
    return fmt.Sprintf("%x", digest.Sum(nil))
}

func readMeta(hash string) MetaData {
    var m MetaData
    content, _ := ioutil.ReadFile(metaPath(hash))
    json.Unmarshal(content, &m)
    return m
}

func writeMeta(hash string, contentType string) {
    m := MetaData{contentType}
    os.MkdirAll(metaDir(hash), os.ModePerm)
    content, err := json.Marshal(m)
    if err != nil {
        fmt.Println("Error while serializing", err)
        return
    }
    ioutil.WriteFile(metaPath(hash), content, os.ModePerm)
}

func downloadFile(url string, hash string) {
    os.MkdirAll(sourceDir(hash), os.ModePerm)
    output, err := os.Create(sourcePath(hash))
    if err != nil {
        fmt.Println("Error while creating", sourcePath(hash), "-", err)
        return
    }
    defer output.Close()

    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Error while downloading", url, "-", err)
        return
    }
    defer response.Body.Close()

    contentType := response.Header.Get("content-type")
    if (strings.Contains(contentType, ";")) {
        contentType = strings.Split(contentType, ";")[0]
        contentType = strings.TrimSpace(contentType)
    }

    writeMeta(hash, contentType)

    n, err := io.Copy(output, response.Body)
    if err != nil {
        fmt.Println("Error while downloading", url, "-", err)
        return
    }

    fmt.Println(n, "bytes downloaded.")
}
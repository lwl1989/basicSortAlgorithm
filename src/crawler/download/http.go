package download

import (
    "log"
    "net/http"
    "io/ioutil"
    "github.com/djimenez/iconv-go"
)

func Get(url string) (string){
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()
    input, err := ioutil.ReadAll(resp.Body)
    out := make([]byte, len(input))
    out = out[:]
    iconv.Convert(input, out, "gb2312", "utf-8")
    return string(out)
}
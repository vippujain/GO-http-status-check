package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"

)

func main() {
    resp, err := http.Get("https://jenkins-nc.wnv2b.cci.att.com/login")
    if err !=nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer resp.Body.Close()
        content, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Printf("%s", resp.StatusCode)
            os.Exit(1)
        }
        fmt.Printf("%s", resp.StatusCode)
        fmt.Printf("%s\n", string(content))
    }
}

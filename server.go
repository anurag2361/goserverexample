package main

import (
        "net/http"
        "fmt"
        "log"
        "bytes"
)

func main() {
        var(
                buf bytes.Buffer
                logger=log.New(&buf,"Hello: ",log.Lshortfile)
        )
        logger.Print("The server runs on localhost:3000")
        fmt.Print(&buf)
        http.Handle("/", http.FileServer(http.Dir("./static")))
        http.ListenAndServe(":3000", nil)
}
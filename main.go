package main

import "net/http"
import "log"
import "fmt"
import "os"

func main() {
    http.HandleFunc("POST /wake/{addr}", func(w http.ResponseWriter, r *http.Request) {
        addr := r.PathValue("addr")
        parsedAddr, err := ParseMac(addr)

        if err != nil {
            log.Print(err)
            w.WriteHeader(400)
            fmt.Fprintf(w, "%s", err)
            return
        }

        log.Printf("Waking %s", addr)
        err = Wake(parsedAddr)
        if err != nil {
            log.Print(err)
            w.WriteHeader(500)
            fmt.Fprintf(w, "%s", err)
            return
        }

        fmt.Fprintf(w, "Ok")
    })

    listenPort := os.Getenv("LISTEN_PORT")
    if listenPort == "" {
        listenPort = "8080"
    }

    var listenOn = fmt.Sprintf(":%s", listenPort)
    log.Printf("Listening on %s...", listenOn)
    log.Fatal(http.ListenAndServe(listenOn, nil))
}

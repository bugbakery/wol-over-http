package main

import "net/http"
import "log"
import "fmt"

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

    var listenOn = ":8080";

    log.Printf("Listening on %s...", listenOn)
    log.Fatal(http.ListenAndServe(listenOn, nil))
}

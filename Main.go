package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    fs := http.FileServer(http.Dir("./files"))
    http.Handle("/", fs)

    // Use Railway-assigned PORT
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Printf("Server starting on port %s, serving files from ./files...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

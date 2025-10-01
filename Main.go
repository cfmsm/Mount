package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "https://sites.google.com/view/aerospry/home", http.StatusFound)
    })
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Printf("Server starting on port %s, redirecting all traffic...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

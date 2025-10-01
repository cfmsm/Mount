package main
import (
    "fmt"
    "log"
    "net/http"
    "os"
)
func main() {
    cwd, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    fs := http.FileServer(http.Dir(cwd))
    http.Handle("/", fs)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    fmt.Printf("Server starting on port %s, serving files from %s...\n", port, cwd)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}

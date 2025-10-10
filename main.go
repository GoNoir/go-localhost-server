package main 

import (
    "fmt"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    srv := http.Server{Addr: ":8080", Handler: mux}

    if err := srv.ListenAndServe(); err != nil {
        fmt.Println("server error:", err)
    }
}

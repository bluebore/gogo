package main

import (
    "fmt"
    "log"
    "net/http"
)


func abs(v int) int {
    if v > 0 {
        return v
    } else {
        return -v
    }
}

func homepage(writer http.ResponseWriter, request *http.Request) {
    request.ParseForm()
    fmt.Fprint(writer, "Hello golang web service")
}

func main() {
    fmt.Println("Utils run");

    http.HandleFunc("/", homepage)
    if err := http.ListenAndServe(":8203", nil); err != nil {
        log.Fatal("failed to start server", err)
    }
}

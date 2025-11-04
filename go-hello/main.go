package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    jokes := []string{
        "Why do Gophers love Go? Because it makes concurrency easy!",
        "I told my code a joke... it didn’t catch it!",
        "Go programmers never panic… oh wait.",
    }
    rand.Seed(time.Now().UnixNano())
    fmt.Fprintf(w, "%s", jokes[rand.Intn(len(jokes))])
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Joke API running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
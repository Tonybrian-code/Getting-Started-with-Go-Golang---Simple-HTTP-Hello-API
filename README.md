# Getting-Started-with-Go-Golang---Simple-HTTP-Joke-API

## 1. Title & Objective

Title: Getting Started with Go (Golang) – Simple-HTTP-Joke-API
Technology Chosen: Go (Golang)
**Reason for Choice:**
Go is a modern, open-source programming language designed for simplicity, concurrency, and performance. It’s used by major companies such as Google, Uber, and Dropbox for building fast, scalable applications.

**Goal:**
Create and run a simple Joke API HTTP server using Go to understand its syntax, setup, and workflow.

## 2. Quick Summary of the Technology

**What is Go (Golang)?**
Go is a statically typed, compiled programming language created at Google. It combines the performance of C with the readability of modern languages.

**Where is it used?**
Go is used in backend web development, cloud tools, microservices, and DevOps platforms.

**Real-world example:**
Docker, Kubernetes, and Terraform — all major DevOps tools — are written in Go.

## 3. System Requirements

| Requirement | Description |
|--------------|-------------|
| **OS** | Windows, macOS, or Linux |
| **Text Editor / IDE** | VS Code (recommended) or GoLand |
| **Go Version** | Go 1.22 or higher |
| **Internet** | Required for setup and dependency installation |

---

##  4. Installation & Setup

### 1. Install Go

Download and install Go from the official website:  
👉 [https://go.dev/dl/](https://go.dev/dl/)

### 2. Verify Installation

Open a terminal or command prompt and run:
```bash
go version
#Expected output: go version go1.22.0 windows/amd64
```

### 3. Create Workspace
```bash
mkdir go-hello
cd go-hello

# Initialize your module:
go mod init go-hello
#creates: go.mod file
```
### 4. Create your main file
```
touch main.go
```

## 5. Minimal Working Example
### Themed Hello World: Joke API
```
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
    fmt.Fprintf(w, jokes[rand.Intn(len(jokes))])
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Joke API running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

### Run it
```
go run main.go
```

### Expected output
```
#Every page refresh shows a new joke!
```

### Visit: http://localhost:8080
### You should see:

a new joke each time the page refreshes

## 6. AI Prompt Journal

### Setup guide and download links
**prompt:** How do I install Go on Windows?

**AI’s response summary:**
Provided step-by-step setup instructions including downloading Go, running the installer, and verifying the installation.
**Evaluation:** Helpful

### working code example
**prompt:** Show me how to create a simple HTTP server in Go.

**AI’s response summary:**
Provided a working code example using net/http, including a simple handler function that responds to requests.
**Evaluation:** Very useful

### import, func main(), and http usage
**prompt:** Explain each line like I’m a beginner.

**AI’s response summary:**
Explained import, func main(), and the usage of http.HandleFunc and http.ListenAndServe in simple terms.
**Evaluation:** Deep understanding

## 7. Common Issues & Fixes

### Issue 1: go: command not found

**Cause:** Go not added to PATH

**Solution:** Add Go’s bin folder to system environment variables

### Issue 2: Blank page

**Cause:** Port conflict

**Solution:** Add response writer line

### Issue 3: Module error

**Cause:** go mod not initialized

**Solution:** Run go mod init go-hello

## 8. References

- [Official Go Docs](https://go.dev/doc/)
- [Tour of Go](https://tour.golang.org/)
- [Go by Example](https://gobyexample.com/)

## Comparison – Go vs Node.js

| Feature     | Go (Golang)         | Node.js                  |
| ----------- | ------------------- | ------------------------ |
| Type        | Compiled            | Interpreted              |
| Speed       | Very fast           | Moderate                 |
| Concurrency | Built-in goroutines | Async model              |
| Ideal Use   | APIs, DevOps tools  | Web apps, real-time apps |
| Syntax      | Simple and clean    | JavaScript-based         |

**Insight:** Go feels cleaner and more efficient for backend services, while Node.js is easier for JavaScript developers.





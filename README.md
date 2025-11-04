# Getting-Started-with-Go-Golang---Simple-HTTP-Hello-API

1. Title & Objective

Title: Getting Started with Go (Golang) – Simple-HTTP-Hello-API
Technology Chosen: Go (Golang)
Reason for Choice:
Go is a modern, open-source programming language designed for simplicity, concurrency, and performance. It’s used by major companies such as Google, Uber, and Dropbox for building fast, scalable applications.

Goal:
Create and run a simple “Hello, World!” HTTP server using Go to understand its syntax, setup, and workflow.

2. Quick Summary of the Technology

What is Go (Golang)?
Go is a statically typed, compiled programming language created at Google. It combines the performance of C with the readability of modern languages.

Where is it used?
Go is used in backend web development, cloud tools, microservices, and DevOps platforms.

Real-world example:
Docker, Kubernetes, and Terraform — all major DevOps tools — are written in Go.

3. System Requirements
| Requirement           | Details                                      |
| --------------------- | -------------------------------------------- |
| **OS**                | Windows, macOS, or Linux                     |
| **Text Editor / IDE** | VS Code (recommended) or GoLand              |
| **Go SDK**            | Version 1.22+                                |
| **Internet Access**   | Required for downloading Go and dependencies |

4.Installation & Setup Instructions
Step 1: Install Go

Visit https://go.dev/dl/

Download the installer for your OS.

Follow the installation steps.

Step 2: Verify Installation
go version


Expected output:

go version go1.22.0 windows/amd64

Step 3: Create Workspace
mkdir go-hello
cd go-hello


Initialize your module:

go mod init go-hello

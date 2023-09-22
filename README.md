# How to start writting Golang API program/project

# Golang First API

A brief description of your project.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

Create a folder for your Golang project e.g "Go_API_Project"
Open the folder in VsCode
Open VsCode terminal
Run

1.  go mod init
2.  go mod init Go_API_Project
    ## This is specify the module path name of your Go project, the name can be any suitable name
3.  go get github.com/gin-gonic/gin
    ## This will install gonic-gin. "Gin" is a popular web framework for the Go programming language, and "Gonic" is one of its implementations.
4.  Create main.go file in your "Go_API_Project" folder, copy and paste below starter code into it

    ```
    package main

    import (
    "github.com/gin-gonic/gin"
    "net/http"
    )

    func main() {
    r := gin.Default()

        r.GET("/", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
        })

        r.Run(":8080")

    }
    ```

5.  Open VsCode terminal and run
    go run main.go

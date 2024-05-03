package main

import (
	// "fmt"
	// "log"
	// "net/http"
	// "os"

	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	// "github.com/joho/godotenv"
	"context"
	"fmt"

	"gihub.com/Vibgitcode27/rssBack/applications"
)

func main() {
	app := applications.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("failed to start app : %v", err)
	}
}

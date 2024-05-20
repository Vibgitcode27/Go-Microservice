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
	"os"
	"os/signal"

	"gihub.com/Vibgitcode27/rssBack/applications"
)

func main() {
	app := applications.New()

	// We actually don't have any other context therefore we are creating a root level context
	// We are doing this all bs for graceful shutdown

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt /* Which is equal to SIGINT or ctrl + c*/)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Printf("failed to start app : %v", err)
	}
}

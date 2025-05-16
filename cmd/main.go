package main

import (
	"devops_project/api"
	db_utils "devops_project/db/utils"
	"fmt"
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("APP_ENV") // e.g., "production" or "development"

	var envFile string
	switch env {
	case "production":
		envFile = ".env.production"
	default:
		envFile = ".env"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("No %s file found or error loading it\n", envFile)
	}
	err = sentry.Init(sentry.ClientOptions{
		Dsn:              "https://003d0b4a32dd21b03d532fc6c74bcbba@o4509322526916608.ingest.us.sentry.io/4509322528292864",
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	db, err := db_utils.InitDB()
	if err != nil {
		panic("Unable to connecto to db")
	}
	r := gin.Default()
	r.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
	api.RegisterRoutes(r, db)
	r.Run(":8080")
}

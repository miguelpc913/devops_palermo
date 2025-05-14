package main

import (
	"devops_project/api"
	"fmt"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

var err = sentry.Init(sentry.ClientOptions{
	Dsn: "https://003d0b4a32dd21b03d532fc6c74bcbba@o4509322526916608.ingest.us.sentry.io/4509322528292864",
})

func main() {
	if err == nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	r := gin.Default()
	r.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
	api.RegisterRoutes(r)
	r.Run(":8080")
}

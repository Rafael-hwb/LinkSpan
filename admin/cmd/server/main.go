package main

import (
	"context"
	"log"
	"time"

	"linkspan/admin/bootstrap"
	"linkspan/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := bootstrap.OpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	redisClient, err := bootstrap.OpenRedis()
	if err != nil {
		return err
	}
	defer redisClient.Close()

	r := gin.New()
	r.Use(middleware.RequestLogger())
	r.Use(middleware.ErrorHandler())

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Minute,
	)

	err = bootstrap.RegisterRoutes(
		ctx,
		r,
		db,
		redisClient,
	)

	cancel()

	if err != nil {
		return err
	}

	return r.Run(":8001")
}

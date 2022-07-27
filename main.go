package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/daniel1sender/alura-flix/pkg/domain/videos/usecases"
	"github.com/daniel1sender/alura-flix/pkg/gateways"
	handler "github.com/daniel1sender/alura-flix/pkg/gateways/http/videos"
	"github.com/daniel1sender/alura-flix/pkg/gateways/postgres/videos"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	ctx := context.Background()
	if err := gateways.RunMigrations(); err != nil {
		log.Printf("error to run migrations: %v", err)
	}
	conn, _ := pgxpool.Connect(ctx, "postgres://postgres:1234@localhost:5432/challenge?sslmode=disable")

	flixStorage := videos.NewStorage(conn)
	flixUseCase := usecases.NewUseCase(flixStorage)
	handler := handler.NewHandler(flixUseCase)

	r := gin.Default()
	r.GET("/videos", handler.GetAll)
	r.GET("/videos/:id", handler.GetByID)
	r.POST("/videos", handler.Create)
	r.DELETE("/videos/:id", handler.DeleteByID)
	r.PATCH("/videos/:id", handler.UpdateByID)

	server := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to listen and serve due to: %s", err)
	}
}

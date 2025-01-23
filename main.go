package main

import (
	"bucketX/config"
	_ "bucketX/docs"
	"bucketX/middlewares"
	"bucketX/routes"
	metadataObject "bucketX/services/file_metadataObject"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"go.uber.org/zap"
)

// @title           BucketX API
// @version         1.0
// @description     This is the API documentation for BucketX API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   X7 team

// @BasePath  /api/v1/
func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load config", zap.Error(err))
	}

	log.Println(cfg)

	if err := metadataObject.Initialize(cfg.Metadata); err != nil {
		logger.Fatal("Failed to initialize metadata service", 
			zap.String("file_path", cfg.Metadata.FilePath),
			zap.Error(err),
		)
	}

	router := gin.Default()
	router.Use(middlewares.LoggerMiddleware(logger))
	router.StaticFile("/docs/swagger.json", "./docs/swagger.json")
	router.GET("/swagger", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "http://localhost:8080/docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "BucketX API Reference",
			},
			DarkMode: true,
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate API reference",
			})
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})


	routes.RegisterRoutes(router)

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownGraceTime)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exited properly")
}
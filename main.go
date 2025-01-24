package main

import (
	"bucketX/config"
	"bucketX/controllers"
	_ "bucketX/docs"
	"bucketX/middlewares"
	"bucketX/routes"
	"bucketX/services"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

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

	if err := services.Initialize(cfg.Metadata); err != nil {
		logger.Fatal("Failed to initialize metadata service",
			zap.String("file_path", cfg.Metadata.FilePath),
			zap.Error(err),
		)
	}

	router := gin.Default()
	router.Use(middlewares.LoggerMiddleware(logger))
	router.StaticFile("/docs/swagger.json", "./docs/swagger.json")
	router.GET("/swagger", controllers.ServeDocsController)

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

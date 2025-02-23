package main

import (
	"bucketX/config"
	"bucketX/controllers"
	_ "bucketX/docs"
	"bucketX/middlewares"
	"bucketX/routes"
	"bucketX/services/metadataObject"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"time"

	"github.com/gin-contrib/cors"
	"go.uber.org/zap"
)

// @title           BucketX API
// @version         0.1
// @description     This is the API documentation for BucketX API.

// @contact.name   X7 team

// @BasePath  /api/v1/
func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	metadataObject.InitializeAoF()
	defer metadataObject.AOF.Close()

	cfg, configErr := config.LoadConfig()
	if configErr != nil {
		logger.Fatal("Failed to load config", zap.Error(configErr))
	}

	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(middlewares.LoggerMiddleware(logger))
	// Turned off for now
	// router.Use(middlewares.AuthMiddleware(cfg.Server.AccessKey))
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
		if er := srv.ListenAndServe(); er != nil && er != http.ErrServerClosed {
			logger.Fatal("Server failed to start", zap.Error(er))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownGraceTime)
	defer cancel()

	if shutdownErr := srv.Shutdown(ctx); shutdownErr != nil {
		logger.Error("Server forced to shutdown", zap.Error(shutdownErr))
	}

	logger.Info("Server exited properly")
}

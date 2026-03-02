package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"event-backend/infrastructure/config"
	"event-backend/infrastructure/database"
	"event-backend/infrastructure/middlewares"
	redisFactory "event-backend/infrastructure/redis/factories"
	redisInterfaces "event-backend/infrastructure/redis/interfaces"
	redisServices "event-backend/infrastructure/redis/services"
	"event-backend/migration"
	"event-backend/seeder"
)

var (
	router            *gin.Engine
	db                *gorm.DB
	redisCache        redisInterfaces.RedisCacheInterface
	redisLock         redisInterfaces.RedisDistributedLockInterface
	execMigration     *string
	runMigration      *string
	migrationFileName *string
	runSeeder         *string
	seederClass       *string
)

func main() {
	extractArgs()
	if *runMigration == "true" && *execMigration == "create" {
		migration.Create(nil, *migrationFileName)
		os.Exit(0)
	}
	initializeDatabase()
	handleMigrationAndSeeding()
	initializeRedis()
	initializeRouter()
	initializeRepositories()
	initializeServices()
	startHttpServer()
}

func extractArgs() {
	execMigration = flag.String("exec", "up", "--exec [up/down/fresh/create]")
	runMigration = flag.String("migration", "false", "--migration [true/false]")
	migrationFileName = flag.String("fileName", "", "--fileName <name>")
	runSeeder = flag.String("dbseed", "false", "--dbseed [true/false]")
	seederClass = flag.String("class", "", "--class [SeederName,...] (optional)")
	flag.Parse()
}

func handleMigrationAndSeeding() {
	if *runMigration == "true" {
		migration.Run(db, *execMigration)
		os.Exit(0)
	}
	if *runSeeder == "true" {
		var classes []string
		if *seederClass != "" {
			classes = strings.Split(*seederClass, ",")
		}
		if err := seeder.Run(db, classes); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}

func initializeDatabase() {
	var err error
	db, err = database.NewDBConnection()
	if err != nil {
		panic(err)
	}
	log.Println("database initialized")
}

func initializeRedis() {
	client, err := redisFactory.NewRedisClient()
	if err != nil {
		panic(err)
	}
	redisCache = redisServices.NewRedisCacheService(client)
	redisLock = redisServices.NewRedisDistributedLockService(client)
	log.Println("redis initialized")
}

func initializeRouter() {
	router = gin.New()
	router.ContextWithFallback = true

	gin.SetMode(config.AppEnv)

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))
	router.Use(middlewares.ExceptionMiddleware())
}

func initializeRepositories() {
	// akan diisi nanti per domain
}

func initializeServices() {
	// akan diisi nanti per domain
}

func startHttpServer() {
	srv := &http.Server{
		Addr:    ":" + config.AppPort,
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("starting server on port %s", config.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server stopped")
}

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

	categoryInterfaces "event-backend/app/category/interfaces"
	categoryRepositories "event-backend/app/category/repositories"
	categoryServices "event-backend/app/category/services"
	categoryControllers "event-backend/presentation/http/category/controllers"

	authInterfaces "event-backend/app/auth/interfaces"
	authRepositories "event-backend/app/auth/repositories"
	authServices "event-backend/app/auth/services"
	authControllers "event-backend/presentation/http/auth/controllers"

	roleInterfaces "event-backend/app/role/interfaces"
	roleRepositories "event-backend/app/role/repositories"
	roleServices "event-backend/app/role/services"
	roleControllers "event-backend/presentation/http/role/controllers"

	eventInterfaces "event-backend/app/event/interfaces"
	eventRepositories "event-backend/app/event/repositories"
	eventServices "event-backend/app/event/services"
	eventControllers "event-backend/presentation/http/event/controllers"

	userInterfaces "event-backend/app/user/interfaces"
	userRepositories "event-backend/app/user/repositories"
	userServices "event-backend/app/user/services"
	userControllers "event-backend/presentation/http/user/controllers"

	eventregistrationInterfaces "event-backend/app/event_registration/interfaces"
	eventregistrationRepoInterfaces "event-backend/app/event_registration/interfaces/repositories"
	eventregistrationRepositories "event-backend/app/event_registration/repositories"
	eventregistrationServices "event-backend/app/event_registration/services"
	eventregistrationControllers "event-backend/presentation/http/event_registration/controllers"
)

var (
	router                           *gin.Engine
	db                               *gorm.DB
	redisCache                       redisInterfaces.RedisCacheInterface
	redisLock                        redisInterfaces.RedisDistributedLockInterface
	execMigration                    *string
	flagMigration                    *string
	migrationFileName                *string
	runSeeder                        *string
	flagSeeder                       *string
	seederClass                      *string
	categoryQueryRepository          categoryInterfaces.CategoryQueryRepositoryInterface
	categoryStoreRepository          categoryInterfaces.CategoryStoreRepositoryInterface
	categoryService                  categoryInterfaces.CategoryServiceInterface
	authQueryRepository              authInterfaces.AuthQueryRepositoryInterface
	authStoreRepository              authInterfaces.AuthStoreRepositoryInterface
	authService                      authInterfaces.AuthServiceInterface
	roleQueryRepository              roleInterfaces.RoleQueryRepositoryInterface
	roleStoreRepository              roleInterfaces.RoleStoreRepositoryInterface
	roleService                      roleInterfaces.RoleServiceInterface
	eventQueryRepository             eventInterfaces.EventQueryRepositoryInterface
	eventStoreRepository             eventInterfaces.EventStoreRepositoryInterface
	eventTicketQueryRepository       eventInterfaces.EventTicketQueryRepositoryInterface
	eventTicketStoreRepository       eventInterfaces.EventTicketStoreRepositoryInterface
	eventService                     eventInterfaces.EventServiceInterface
	userQueryRepository              userInterfaces.UserQueryRepositoryInterface
	userStoreRepository              userInterfaces.UserStoreRepositoryInterface
	userService                      userInterfaces.UserServiceInterface
	eventRegistrationQueryRepository eventregistrationRepoInterfaces.EventRegistrationQueryRepositoryInterface
	eventRegistrationStoreRepository eventregistrationRepoInterfaces.EventRegistrationStoreRepositoryInterface
	eventRegistrationService         eventregistrationInterfaces.EventRegistrationServiceInterface
)

func main() {
	extractArgs()
	initializeDatabase()
	runnerMigration()
	runnerSeeder()
	initializeRedis()
	initializeRouter()
	initializeRepositories()
	initializeServices()
	initializeControllers()
	initializeHttpServer()
}

func extractArgs() {
	execMigration = flag.String("exec", "up", "--exec [up/down/fresh/create]")
	flagMigration = flag.String("migration", "false", "--migration [true/false]")
	migrationFileName = flag.String("fileName", "", "--fileName <name>")
	flagSeeder = flag.String("dbseed", "false", "--dbseed [true/false]")
	seederClass = flag.String("class", "", "--class [SeederName,...] (optional)")
	flag.Parse()
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

func runnerMigration() {
	if *flagMigration != "true" {
		return
	}
	if *execMigration == "create" {
		migration.Create(nil, *migrationFileName)
		os.Exit(0)
	}
	migration.Run(db, *execMigration)
	os.Exit(0)
}

func runnerSeeder() {
	if *flagSeeder != "true" {
		return
	}
	var classes []string
	if *seederClass != "" {
		classes = strings.Split(*seederClass, ",")
	}
	if err := seeder.Run(db, classes); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func initializeRouter() {
	router = gin.New()
	router.ContextWithFallback = true

	gin.SetMode(config.AppGinMode)

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(gin.Logger())
	router.Use(cors.New(corsConfig))
	router.Use(middlewares.ExceptionMiddleware())
}

func initializeRepositories() {
	categoryQueryRepository = categoryRepositories.NewCategoryQueryRepository(db)
	categoryStoreRepository = categoryRepositories.NewCategoryStoreRepository(db)
	authQueryRepository = authRepositories.NewAuthQueryRepository(db)
	authStoreRepository = authRepositories.NewAuthStoreRepository(db)
	roleQueryRepository = roleRepositories.NewRoleQueryRepository(db)
	roleStoreRepository = roleRepositories.NewRoleStoreRepository(db)
	eventQueryRepository = eventRepositories.NewEventQueryRepository(db)
	eventStoreRepository = eventRepositories.NewEventStoreRepository(db)
	eventTicketQueryRepository = eventRepositories.NewEventTicketQueryRepository(db)
	eventTicketStoreRepository = eventRepositories.NewEventTicketStoreRepository(db)
	userQueryRepository = userRepositories.NewUserQueryRepository(db)
	userStoreRepository = userRepositories.NewUserStoreRepository(db)
	eventRegistrationQueryRepository = eventregistrationRepositories.NewEventRegistrationQueryRepository(db)
	eventRegistrationStoreRepository = eventregistrationRepositories.NewEventRegistrationStoreRepository(db)
}

func initializeServices() {
	categoryService = categoryServices.NewCategoryService(categoryQueryRepository, categoryStoreRepository)
	authService = authServices.NewAuthService(authQueryRepository, authStoreRepository, roleQueryRepository)
	roleService = roleServices.NewRoleService(roleQueryRepository, roleStoreRepository)
	eventService = eventServices.NewEventService(db, eventQueryRepository, eventStoreRepository, eventTicketQueryRepository, eventTicketStoreRepository, categoryQueryRepository)
	userService = userServices.NewUserService(userQueryRepository, userStoreRepository, roleQueryRepository)
	eventRegistrationService = eventregistrationServices.NewEventRegistrationService(db, eventRegistrationQueryRepository, eventRegistrationStoreRepository, eventTicketQueryRepository, eventTicketStoreRepository)
}

func initializeControllers() {
	categoryControllers.NewCategoryController(router, categoryService)
	authControllers.NewAuthController(router, authService)
	roleControllers.NewRoleController(router, roleService)
	eventControllers.NewEventController(router, eventService)
	userControllers.NewUserController(router, userService)
	eventregistrationControllers.NewEventRegistrationController(router, eventRegistrationService)
}

func initializeHttpServer() {
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

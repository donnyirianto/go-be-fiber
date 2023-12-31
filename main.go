package main

import (
	"time"

	"github.com/donnyirianto/go-be-fiber/configuration"
	"github.com/donnyirianto/go-be-fiber/controller"
	_ "github.com/donnyirianto/go-be-fiber/docs"
	"github.com/donnyirianto/go-be-fiber/exception"
	repository "github.com/donnyirianto/go-be-fiber/repository/impl"
	service "github.com/donnyirianto/go-be-fiber/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title Go Fiber Clean Architecture
// @version 1.0.0
// @description Baseline project using Go Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7777
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description Authorization For JWT
func main() {
	// setup configuration
	config := configuration.New()
	database, err := configuration.NewDatabase(config)
	exception.PanicLogging(err)
	redis, err := configuration.NewRedis(config)
	exception.PanicLogging(err)

	// repository
	productRepository := repository.NewProductRepositoryImpl(database)
	transactionRepository := repository.NewTransactionRepositoryImpl(database)
	transactionDetailRepository := repository.NewTransactionDetailRepositoryImpl(database)
	userRepository := repository.NewUserRepositoryImpl(database)

	// service
	productService := service.NewProductServiceImpl(&productRepository, redis)
	transactionService := service.NewTransactionServiceImpl(&transactionRepository)
	transactionDetailService := service.NewTransactionDetailServiceImpl(&transactionDetailRepository)
	userService := service.NewUserServiceImpl(&userRepository)

	// controller
	productController := controller.NewProductController(&productService, config)
	transactionController := controller.NewTransactionController(&transactionService, config)
	transactionDetailController := controller.NewTransactionDetailController(&transactionDetailService, config)
	userController := controller.NewUserController(&userService, config)

	// setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())

	// use Gzip compression middleware
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(cors.New())

	// Use the cache middleware
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			// Check if Gzip is accepted by the client
			return c.Get("Accept-Encoding") == "gzip" || c.Get("TE") == "gzip;q=0.8, deflate;q=0.6, br;q=0.4"
		},
		Expiration:   1 * time.Minute,
		CacheControl: true,
	}))

	// routing
	productController.Route(app)
	transactionController.Route(app)
	transactionDetailController.Route(app)
	userController.Route(app)

	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		data := map[string]string{"message": "hello!"}
		return c.JSON(data)
	})

	// Metrics Page`
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Be Matrics Page"}))

	// start app
	err = app.Listen(config.GetString("SERVER_PORT"))
	exception.PanicLogging(err)
}

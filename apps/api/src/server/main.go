package server

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func middlewares(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(etag.New())
	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)
}

func New() *fiber.App {
	// middlerware
	// middle := middleware.NewJWTMiddleware(cfg.AuthConfig.JWTSecret)

	// register usecase
	// authHandler := handler.NewAuthHanlder(cfg.AuthConfig)
	// user
	// userRepo := gorm_repository.NewUserGormRepository(db)
	// userUC := usecase.NewUserUsecase(userRepo)
	// userHandler := handler.NewUserHandler(userUC, *cfg)

	// Media
	// mediaHandler := handler.NewMediaHandler(*cfg, minioClient)

	// slug
	// slugRepo := gorm_repository.NewSlugGormRepository(db)
	// slugUC := usecase.NewSlugUseCase(slugRepo)

	// post
	// postRepo := gorm_repository.NewPostGormRepository(db)
	// postUC := usecase.NewPostUseCase(postRepo)
	// postHandler := handler.NewPostHandler(postUC, slugUC, userUC, cfg, minioClient, rdb)

	// comment
	// commentRepo := gorm_repository.NewCommentGormRepository(db)
	// commentUC := usecase.NewCommentUseCase(commentRepo)
	// commentHandler := handler.NewCommentHandler(commentUC, rdb)

	// notifyRepo := notification.NewNotifyRepository(db, rdb)

	// app
	app := fiber.New()
	// app.Use(func(c *fiber.Ctx) error {
	// 	c.Locals("db", db)
	// 	c.Locals("userService", userUC)
	// 	c.Locals("notifyRepository", *notifyRepo)
	// 	return c.Next()
	// })
	app.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/")
	})

	// api := app.Group("/api")
	// v1 := api.Group("/v1")
	// routes.UserRouter(v1, *userHandler, *middle)
	// routes.AuthRouter(v1, *authHandler, *userHandler, *middle)
	// routes.MediaRouter(v1, *mediaHandler, *middle)
	// routes.PostRouter(v1, *postHandler, *middle)
	// routes.CommentRouter(v1, *commentHandler, *middle)
	// reaction.RegisterReactionApi(v1, *middle)

	return app
}

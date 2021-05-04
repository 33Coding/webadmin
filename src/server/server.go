package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"strings"
	"time"
)

func (srv *Webadmin) Run() {

	//RECOVER

	srv.APP.Use(recover.New())
	srv.APP.Use(helmet.New())
	srv.APP.Use(cors.New(cors.Config{
		Next:         nil,
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodOptions,
		}, ","),
		AllowHeaders:     "",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))
	// srv.APP.Use(csrf.New()) //TODO: reinlesen
	srv.APP.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	srv.APP.Static("/", "./web/public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        false,
		CacheDuration: 1 * time.Hour,
		MaxAge:        3600,
	})

	// web
	srv.APP.Get("/", srv.index)
	srv.APP.Get("/user/list", srv.userList)

	// JWT Middleware
	srv.APP.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    srv.PrivateKey.Public(),
	}))

	defer srv.ORM.Close()
	srv.APP.Listen(":3333")
}

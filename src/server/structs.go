package server

import (
	"crypto/rsa"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
)

type Webadmin struct {
	ORM        *pg.DB
	APP        *fiber.App
	PrivateKey *rsa.PrivateKey
}

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"log"
	"webadmin/src/db"
	"webadmin/src/server"
)

var SERVER = server.Webadmin{}

func main() {
	engine := html.New("./web/templates", ".html")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	ormDb, err := db.ConnectORM()

	if err != nil {
		log.Println(err)
		log.Fatal("error connecting to database")
	} else {
		log.Println("db connected")
	}

	rng := rand.Reader
	privateKey, err := rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	SERVER.APP = app

	SERVER.ORM = ormDb

	SERVER.PrivateKey = privateKey

	SERVER.Run()
}

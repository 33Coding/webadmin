package server

import (
	"github.com/gofiber/fiber/v2"
)

func (srv *Webadmin) index(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"Title": "Hello, Twitch 👋!",
	}, "layouts/main")
}
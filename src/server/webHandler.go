package server

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"webadmin/src/db"
)

func (srv *Webadmin) index(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"Title": "Hello, Twitch 👋!",
	}, "layouts/main")
}

func (srv *Webadmin) userList(c *fiber.Ctx) error {

	var users []db.User
	err := srv.ORM.Model(&users).Select()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	header := make([]TableHeader, 2)
	header[0] = TableHeader{Title: "Id"}
	header[1] = TableHeader{Title: "Email"}

	data := make([]TableData, len(users))
	for i, user := range users {
		data[i] = TableData{
			Values: []string{
				strconv.FormatInt(user.Id, 10),
				user.Email,
			},
		}
	}

	return c.Render("list", fiber.Map{
		"Title":  "Users",
		"Data":   data,
		"Header": header,
	}, "layouts/main")
}

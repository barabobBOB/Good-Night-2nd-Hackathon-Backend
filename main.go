package Good_Night_2nd_Hackathon_Backend

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pingpong by fiber\n")
	})

	log.Fatal(app.Listen(":3000"))
}

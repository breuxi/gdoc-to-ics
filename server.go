package main

import (
	"fmt"
	"os"

	"github.com/breuxi/gdoc-to-ics/converter"
	"github.com/breuxi/gdoc-to-ics/loader"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/ical/:id/:filename?", func(c *fiber.Ctx) error {
		gdoc, err := loader.GetGDocCSV(c.Params("id"), c.Query("sheet_id", "0"))

		if err != nil {
			return c.SendStatus(500)
		}

		ics, err := converter.ConvertGDocCSVToIcs(gdoc.Content, gdoc.Filename)

		if err != nil {
			return c.SendStatus(500)
		}

		c.Set("Content-Type", "text/calendar")
		c.Set("Charset", "utf-8")

		var filename string

		if len(c.Params("filename", "")) > 0 {
			filename = c.Params("filename")
		} else {
			filename = gdoc.Filename
		}

		c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", c.Params("filename", filename+".ics")))

		return c.SendString(ics)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://github.com/breuxi/gdoc-to-ics/")
	})

	var addr string

	val, ok := os.LookupEnv("GDTICS_LISTEN")
	if !ok {
		addr = "127.0.0.1:8080"
	} else {
		addr = val
	}

	app.Listen(addr)
}

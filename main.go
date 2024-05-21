package main

import (
	"context"

	containers "github.com/docker/docker/api/types/container"

	"github.com/gofiber/fiber/v2"
	"github.com/nimnest/nimnest/docker"
)

func main() {
	app := fiber.New()

	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK",
		})
	})

	app.Get("/containers", func(c *fiber.Ctx) error {
		containers, err := client.ContainerList(ctx, containers.ListOptions{})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(containers)
	})

	app.Listen(":3000")
}

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Post("/create-blog-post", creatBlogPost)

	app.Post("/retrive-blog-post", retriveBlogPost)

	app.Post("/retrive-all-blog-post", retriveAllBlogs)

	app.Post("/remove-blog-post", deleteBlog)

	log.Fatal(app.Listen(":3000"))
}
